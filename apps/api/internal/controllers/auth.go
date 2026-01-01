package controllers

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"

	"omnicampus/api/internal/db"
	"omnicampus/api/internal/db/sqlc"
	"omnicampus/api/pkg/redis"
	"omnicampus/api/pkg/utils"
)

func checkOnboarded(email string) error {
	useOnboarded, err := db.Queries.GetStudentOnboardingStatusByEmail(context.Background(), email)

	if err != nil {
		return err
	}

	if useOnboarded {
		return errors.New("user already onboarded")
	}

	return nil
}

func SendOTP(email string) error {
	err := checkOnboarded(email)
	if err != nil {
		return err
	}
	
	expireTime, _ := strconv.ParseInt(os.Getenv("OTP_EXPIRY_MINUTES"), 10, 64)

	userExists, err := db.Queries.UserExistsByEmail(context.Background(), email)
	if err != nil {
		return err
	}

	if !userExists {
		return errors.New("user not found")
	}

	key := "otp:" + email

	exists, err := redis.Client.Exists(redis.Ctx, key).Result()
	if err != nil {
		return err
	}

	if exists == 1 {
		return errors.New("otp already sent")
	}

	otp, err := utils.GenerateOTP()
	if err != nil {
		return errors.New("failed to generate otp")
	}

	hashed := utils.HashOTP(otp)

	err = redis.Client.Set(
		redis.Ctx,
		key,
		hashed,
		time.Duration(expireTime) * time.Minute,
	).Err()

	if err != nil {
		return err
	}

	err = utils.SendMail(email, otp)
	if err != nil {
		redis.Client.Del(redis.Ctx, key)
		return err
	}

	return nil
}

func VerifyOTP(email, otp string) error {
	err := checkOnboarded(email)
	if err != nil {
		return err
	}

	key := "otp:" + email

	storedHash, err := redis.Client.Get(redis.Ctx, key).Result()
	if err != nil {
		if redis.IsKeyMissing(err) {
			return errors.New("otp expired or invalid")
		}
		return err
	}

	if utils.HashOTP(otp) != storedHash {
		return errors.New("invalid otp")
	}

	redis.Client.Del(redis.Ctx, key)

	return nil
}

func SetPassword(email, password string) error {
	err := checkOnboarded(email)
	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	userID, err := db.Queries.GetUserIDByEmail(context.Background(), email)
	if err != nil {
		return err
	}

	err = db.Queries.UpdateUserPasswordByID(
		context.Background(),
		sqlc.UpdateUserPasswordByIDParams{
			ID:       userID,
			Password: hashedPassword,
		},
	)

	err = db.Queries.SetStudentOnboardedByEmail(context.Background(), email)
	if err != nil {
		return err
	}

	return err
}

func LoginUser(email, password string) (string, error) {
	user, err := db.Queries.GetUserByEmail(context.Background(), email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = utils.CheckPassword(password, user.Password)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	return utils.CreateJWT(user.ID.String(), email)
}

func Me(tokenStr string) (interface{}, error) {
	parsed, err := jwt.Parse(
		[]byte(tokenStr),
		jwt.WithKey(jwa.HS256, utils.JwtKey()),
		jwt.WithValidate(true),
	)
	if err != nil {
		return nil, errors.New("invalid or expired token")
	}

	userID := parsed.Subject()
	if userID == "" {
		return nil, errors.New("invalid token payload")
	}

	var uid pgtype.UUID
	if err := uid.Scan(userID); err != nil {
		return nil, errors.New("invalid user id in token")
	}

	user, err := db.Queries.GetStudentByID(context.Background(), uid)
	if err != nil {
		return nil, errors.New("user not found")
	}

	user.Password = ""

	return user, nil
}



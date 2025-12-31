package controllers

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

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
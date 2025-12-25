package controllers

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	"omnicampus/api/internal/db"
	"omnicampus/api/pkg/redis"
	"omnicampus/api/pkg/utils"
)

func SendOTP(email string) error {

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

	println("OTP:", otp)

	return nil
}

func VerifyOTP(email, otp string) (string, error) {
	key := "otp:" + email

	storedHash, err := redis.Client.Get(redis.Ctx, key).Result()
	if err != nil {
		if redis.IsKeyMissing(err) {
			return "", errors.New("otp expired or invalid")
		}
		return "", err
	}

	if utils.HashOTP(otp) != storedHash {
		return "", errors.New("invalid otp")
	}

	redis.Client.Del(redis.Ctx, key)

	userID, err := db.Queries.GetUserIDByEmail(context.Background(), email)
	if err != nil {
		return "", err
	}

	return utils.CreateJWT(userID.String(), email)
}
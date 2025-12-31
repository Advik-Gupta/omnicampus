package utils

import (
	"errors"
	"hash/fnv"
	"net/smtp"
	"os"
	"strconv"
)

func SendMail(email, otp string) error {
	cfg := PickSMTP(email)

	auth := smtp.PlainAuth("", cfg.User, cfg.Pass, cfg.Host)

	if cfg.Host == "" || cfg.User == "" || cfg.Pass == "" {
		return errors.New("smtp config missing")
	}

	msg := []byte(
		"Subject: Omnicampus Login OTP\n" +
			"\r\n" +
			"Your OTP for Omnicampus is: " + otp + "\n" +
			"Valid for 5 minutes.",
	)

	return smtp.SendMail(
		cfg.Host+":"+cfg.Port,
		auth,
		cfg.User,
		[]string{email},
		msg,
	)

}

type SMTPConfig struct {
	Host string
	Port string
	User string
	Pass string
}

func PickSMTP(email string) SMTPConfig {
	h := fnv.New32a()
	h.Write([]byte(email))

	count, _ := strconv.Atoi(os.Getenv("SMTP_ACCOUNTS"))
	if count == 0 {
		count = 1
	}

	idx := int(h.Sum32() % uint32(count))

	prefix := "SMTP_" + strconv.Itoa(idx) + "_"

	return SMTPConfig{
		Host: os.Getenv(prefix + "HOST"),
		Port: os.Getenv(prefix + "PORT"),
		User: os.Getenv(prefix + "USER"),
		Pass: os.Getenv(prefix + "PASS"),
	}
}
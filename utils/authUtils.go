package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(userID int64) (string, time.Time, error) {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return "", time.Time{}, errors.New("secret key not found in environment variables")
	}

	expiration := time.Now().Add(time.Hour * 24) // 1 Day

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expiration.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expiration, nil
}

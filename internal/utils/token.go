package utils

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateRefreshToken() string {
	return uuid.NewString()
}

func HashRefreshToken(currentvalue string) string {
	hashoftoken, err := bcrypt.GenerateFromPassword([]byte(currentvalue), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err.Error())
	}
	return string(hashoftoken)
}

func CompareRefreshToken(hash, currentvalue []byte) error {
	if err := bcrypt.CompareHashAndPassword(hash, currentvalue); err != nil {
		return err
	}
	return nil
}

func CreateAccessToken(secretKey string, userId uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Second * 60 * 15).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func ValidateToken(tokenString, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token is expired")
		}
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

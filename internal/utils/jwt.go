package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateJWT(userID uuid.UUID) (accessToken string, refreshToken string, err error) {
	accessClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
		"type":    "access",
	}
	access := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err = access.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", "", err
	}

	refreshClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
		"type":    "refresh",
	}
	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err = refresh.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"learn-gin-gorm/configs"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GetTokenLifespan() (time.Duration, error) {
	configs.LoadEnv()

	hoursStr := os.Getenv("JWT_TOKEN_HOUR_LIFESPAN")
	if hoursStr == "" {
		return 0, fmt.Errorf("JWT_TOKEN_HOUR_LIFESPAN is not set or empty")
	}

	hours, err := strconv.Atoi(hoursStr)
	if err != nil {
		return 0, fmt.Errorf("failed to convert JWT_TOKEN_HOUR_LIFESPAN to integer: %v", err)
	}
	return time.Duration(hours) * time.Hour, nil
}

func ExtractTokenFromHeader(header string) string {
	tokenString := ""
	if len(header) > 7 && header[:7] == "Bearer " {
		tokenString = header[7:]
	}
	return tokenString
}

func GenerateToken(email string) (string, error) {
	configs.LoadEnv()
	jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	duration, err := GetTokenLifespan()
	if err != nil {
		return "", err
	}

	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*Claims, error) {
	configs.LoadEnv()
	jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	return claims, nil
}

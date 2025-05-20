package helpers

import (
	"errors"
	"fmt"
	"go-fiber-vercel/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	jwtSecret = []byte("awkajhwdau8121k2312938ainwioa8121kna8wye1923")
	// jwtExpiry = 24 * time.Hour
	jwtExpiry = 5 * time.Minute
)

type JWTModels struct {
	UserID   uint   `json:"userId"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	jwt.RegisteredClaims
}

func GenerateToken(user models.User) (string, error) {
	claims := JWTModels{
		user.ID,
		user.Name,
		user.Email,
		user.Nickname,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*JWTModels, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTModels{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTModels); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func HashPassword(password, email, gameId, serverId string) (string, error) {
	combinedPassword := fmt.Sprintf("%s:%s:%s:%s", password, email, gameId, serverId)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(combinedPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword, password, email, gameId, serverId string) error {
	combinedPassword := fmt.Sprintf("%s:%s:%s:%s", password, email, gameId, serverId)

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(combinedPassword))
}

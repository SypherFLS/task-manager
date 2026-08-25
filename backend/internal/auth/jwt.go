package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTManager struct {
	secret []byte 
}

func NewJWTManager(secret string) *JWTManager {
    return &JWTManager{
        secret: []byte(secret),
    }
}

func (j *JWTManager) Generate(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id" : userID,
		"exp" : time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(j.secret)
}

func (j *JWTManager) Validate(token string) (int, error)
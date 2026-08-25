package auth

import (
	"time"
	"tm/internal/apperrors"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secret []byte
}

func NewJWTManager(secret string) *JWTManager {
	return &JWTManager{
		secret: []byte(secret),
	}
}

type Claims struct {
	UserID int `json:"user_id"`

	jwt.RegisteredClaims
}

func (j *JWTManager) Generate(userID int) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(24 * time.Hour),
			),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(j.secret)
}

func (j *JWTManager) Validate(tokenString string) (int, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (any, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, apperrors.WrongSignMethod
			}

			return j.secret, nil
		},
	)

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, apperrors.InvalidToken
	}

	return claims.UserID, nil
}

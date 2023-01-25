package token_util

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"

	token_const "parishioner_management/internal/constant/token"
)

type JwtCustomClaims struct {
	UserID        string `json:"user_id"`
	ImpersonateBy int    `json:"impersonate_by,omitempty"`
	jwt.StandardClaims
}

//Generate new jwt token
func NewToken(ctx context.Context, userID string, jwtSecret string) (string, error) {
	//Create token class with signing method and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(token_const.ExpiredAccessToken).Unix(),
		},
	})

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", err
	}

	return t, nil
}

func GetExpirationTime(ctx context.Context, tokenString string, secretKey string) (int64, error) {
	claims, err := Decode(ctx, tokenString, secretKey)

	if err != nil {
		return 0, err
	}

	value, found := claims[token_const.ExpirationTimeKey]

	if !found {
		return 0, errors.New("token invalid")
	}

	return int64(value.(float64)) * 1000, nil
}

func Decode(ctx context.Context, tokenString string, secretKey string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	if _, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	}); err != nil {
		return nil, err
	}

	return claims, nil
}

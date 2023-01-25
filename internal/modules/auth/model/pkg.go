package auth_model

import (
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrUsernameOrPasswordNotMatch = errors.New("username or password not match")
	ErrGenTokenFailed             = errors.New("generate token failed")
	ErrInvalidToken               = errors.New("invalid token")
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (data UserLogin) Validate() error {
	return validation.ValidateStruct(&data,
		validation.Field(&data.Username, validation.Required),
		validation.Field(&data.Password, validation.Required, validation.Length(8, 30)),
	)
}

type UserRefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

func (data UserRefreshToken) Validate() error {
	return validation.ValidateStruct(&data,
		validation.Field(&data.RefreshToken, validation.Required),
	)
}

type UserResponse struct {
	ID        primitive.ObjectID `json:"id,omitempty"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	UserName  string             `json:"username"`
	FullName  string             `json:"full_name"`
	Phone     string             `json:"phone"`
	Email     string             `json:"email"`
}

type LoginResponse struct {
	Token               string       `json:"token"`
	RefreshToken        string       `json:"refresh_token"`
	TokenExpirationTime int64        `json:"token_expiration_time"`
	AuthInfo            UserResponse `json:"auth_info"`
}

type RefreshTokenResponse struct {
	Token               string `json:"token"`
	TokenExpirationTime int64  `json:"token_expiration_time"`
}

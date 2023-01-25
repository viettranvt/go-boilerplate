package auth_business

import (
	"context"
	"os"
	"time"

	configs "parishioner_management/internal/configs"
	database_field_const "parishioner_management/internal/constant/database/field"
	account_database "parishioner_management/internal/databases/account"
	refresh_token_database "parishioner_management/internal/databases/refresh-token"
	auth_model "parishioner_management/internal/modules/auth/model"
	date_util "parishioner_management/internal/utils/date"
	hash_util "parishioner_management/internal/utils/hash"
	token_util "parishioner_management/internal/utils/token"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginStorage interface {
	GetAccount(ctx context.Context, filters map[string]interface{}, withDeleted bool, moreKeys ...interface{}) (*account_database.Model, error)
	InsertRefreshToken(ctx context.Context, userID primitive.ObjectID, userAgent string) (*refresh_token_database.Model, error)
}

type loginBusiness struct {
	store LoginStorage
}

func NewLoginBusiness(store LoginStorage) *loginBusiness {
	return &loginBusiness{store: store}
}

func (biz *loginBusiness) GetAccountByUserName(ctx context.Context, username string) (*account_database.Model, error) {
	filter := make(map[string]interface{})
	filter[database_field_const.Username] = username

	return biz.store.GetAccount(ctx, filter, false)
}

func (biz *loginBusiness) AuthenticatePassword(ctx context.Context, password string, passwordHash string) error {
	if !hash_util.IsValidPassword(password, passwordHash) {
		return auth_model.ErrUsernameOrPasswordNotMatch
	}

	return nil
}

func (biz *loginBusiness) GenAccessToken(ctx context.Context, userID string) (string, error) {
	secretKey := os.Getenv(configs.EnvJwtSecretKey)
	token, err := token_util.NewToken(ctx, userID, secretKey)

	if err != nil {
		return "", auth_model.ErrGenTokenFailed
	}

	return token, nil
}

func (biz *loginBusiness) GetTokenExpirationTime(ctx context.Context, token string) (int64, error) {
	secretKey := os.Getenv(configs.EnvJwtSecretKey)
	tokenExpirationAt, err := token_util.GetExpirationTime(ctx, token, secretKey)

	if err != nil {
		return 0, err
	}

	tokenExpirationDate := date_util.ConvertMillisecondToTime(tokenExpirationAt)
	tokenExpirationTime := date_util.CalculateDateDistanceByMillisecond(time.Now(), tokenExpirationDate)

	return tokenExpirationTime, nil
}

func (biz *loginBusiness) GenRefreshToken(ctx context.Context, userID primitive.ObjectID, userAgent string) (*refresh_token_database.Model, error) {
	return biz.store.InsertRefreshToken(ctx, userID, userAgent)
}

func (biz *loginBusiness) ToResponse(ctx context.Context, data account_database.Model, token string, tokenExpirationTime int64, refreshToken string) auth_model.LoginResponse {
	response := auth_model.LoginResponse{
		AuthInfo: auth_model.UserResponse{
			ID:        data.ID,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
			UserName:  data.UserName,
			FullName:  data.FullName,
			Phone:     data.Password,
			Email:     data.Email,
		},
		Token:               token,
		TokenExpirationTime: tokenExpirationTime,
		RefreshToken:        refreshToken,
	}

	return response
}

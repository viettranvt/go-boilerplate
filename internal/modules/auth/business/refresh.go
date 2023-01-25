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
	token_util "parishioner_management/internal/utils/token"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RefreshTokenStorage interface {
	GetRefreshToken(ctx context.Context, filters map[string]interface{}, withDeleted bool, moreKeys ...interface{}) (*refresh_token_database.Model, error)
	GetAccount(ctx context.Context, filters map[string]interface{}, withDeleted bool, moreKeys ...interface{}) (*account_database.Model, error)
	InsertRefreshToken(ctx context.Context, userID primitive.ObjectID, userAgent string) (*refresh_token_database.Model, error)
}

type refreshTokenBusiness struct {
	store RefreshTokenStorage
}

func NewRefreshTokenBusiness(store RefreshTokenStorage) *refreshTokenBusiness {
	return &refreshTokenBusiness{store: store}
}

func (biz *refreshTokenBusiness) FindRefreshTokenInfo(ctx context.Context, refreshToken string) (*refresh_token_database.Model, error) {
	filter := make(map[string]interface{})
	filter[database_field_const.RefreshToken] = refreshToken

	return biz.store.GetRefreshToken(ctx, filter, false)
}

func (biz *refreshTokenBusiness) FindAccountByID(ctx context.Context, accountID primitive.ObjectID) (*account_database.Model, error) {
	filter := make(map[string]interface{})
	filter[database_field_const.ID] = accountID

	return biz.store.GetAccount(ctx, filter, false)
}

func (biz *refreshTokenBusiness) GenAccessToken(ctx context.Context, userID string) (string, error) {
	secretKey := os.Getenv(configs.EnvJwtSecretKey)
	token, err := token_util.NewToken(ctx, userID, secretKey)

	if err != nil {
		return "", auth_model.ErrGenTokenFailed
	}

	return token, nil
}

func (biz *refreshTokenBusiness) GetTokenExpirationTime(ctx context.Context, token string) (int64, error) {
	secretKey := os.Getenv(configs.EnvJwtSecretKey)
	tokenExpirationAt, err := token_util.GetExpirationTime(ctx, token, secretKey)

	if err != nil {
		return 0, err
	}

	tokenExpirationDate := date_util.ConvertMillisecondToTime(tokenExpirationAt)
	tokenExpirationTime := date_util.CalculateDateDistanceByMillisecond(time.Now(), tokenExpirationDate)

	return tokenExpirationTime, nil
}

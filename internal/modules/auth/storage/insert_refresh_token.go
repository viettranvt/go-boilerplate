package auth_storage

import (
	"context"
	database_model_const "parishioner_management/internal/constant/database/model"
	token_const "parishioner_management/internal/constant/token"
	refresh_token_database "parishioner_management/internal/databases/refresh-token"
	database_util "parishioner_management/internal/utils/database"
	"time"

	"github.com/segmentio/ksuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (mongo *mongoStore) InsertRefreshToken(ctx context.Context, userID primitive.ObjectID, userAgent string) (*refresh_token_database.Model, error) {
	refreshTokenCollection := database_util.GetCollection(mongo.db, database_model_const.RefreshToken)
	model := &refresh_token_database.Model{
		AccountID:    userID,
		Device:       userAgent,
		Expired:      time.Now().Add(token_const.ExpiredTimeRefreshToken),
		RefreshToken: ksuid.New().String(),
	}

	if _, err := refreshTokenCollection.InsertOne(ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

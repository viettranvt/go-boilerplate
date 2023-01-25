package auth_storage

import (
	"context"
	database_field_const "parishioner_management/internal/constant/database/field"
	database_model_const "parishioner_management/internal/constant/database/model"
	refresh_token_database "parishioner_management/internal/databases/refresh-token"
	database_util "parishioner_management/internal/utils/database"

	"go.mongodb.org/mongo-driver/bson"
)

func (mongo *mongoStore) GetRefreshToken(ctx context.Context, filters map[string]interface{}, withDeleted bool, moreKeys ...interface{}) (*refresh_token_database.Model, error) {
	refreshTokenCollection := database_util.GetCollection(mongo.db, database_model_const.RefreshToken)
	filter := bson.M{}

	for key, value := range filters {
		filter[key] = value
	}

	if !withDeleted {
		filter[database_field_const.DeletedAt] = nil
	}

	var result refresh_token_database.Model

	if err := refreshTokenCollection.FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

package auth_storage

import (
	"context"
	database_field_const "parishioner_management/internal/constant/database/field"
	database_model_const "parishioner_management/internal/constant/database/model"
	account_database "parishioner_management/internal/databases/account"
	database_util "parishioner_management/internal/utils/database"

	"go.mongodb.org/mongo-driver/bson"
)

func (mongo *mongoStore) GetAccount(ctx context.Context, filters map[string]interface{}, withDeleted bool, moreKeys ...interface{}) (*account_database.Model, error) {
	accountCollection := database_util.GetCollection(mongo.db, database_model_const.Account)
	filter := bson.M{}

	for key, value := range filters {
		filter[key] = value
	}

	if !withDeleted {
		filter[database_field_const.DeletedAt] = nil
	}

	var result account_database.Model

	if err := accountCollection.FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

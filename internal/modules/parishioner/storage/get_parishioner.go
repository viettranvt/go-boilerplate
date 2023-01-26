package parishioner_storage

import (
	"context"
	database_field_const "parishioner_management/internal/constant/database/field"
	database_model_const "parishioner_management/internal/constant/database/model"
	parishioner_database "parishioner_management/internal/databases/parishioner"
	database_util "parishioner_management/internal/utils/database"

	"go.mongodb.org/mongo-driver/bson"
)

func (mongo *mongoStore) GetParishioner(
	ctx context.Context,
	filters map[string]interface{},
	withDeleted bool,
	moreKeys ...interface{},
) (*parishioner_database.Model, error) {
	parishionerCollection := database_util.GetCollection(mongo.db, database_model_const.Parishioner)
	filter := bson.M{}

	for key, value := range filters {
		filter[key] = value
	}

	if !withDeleted {
		filter[database_field_const.DeletedAt] = nil
	}

	var result parishioner_database.Model

	if err := parishionerCollection.FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

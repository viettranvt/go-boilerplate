package parishioner_storage

import (
	"context"
	"parishioner_management/internal/common"
	database_field_const "parishioner_management/internal/constant/database/field"
	database_model_const "parishioner_management/internal/constant/database/model"
	parishioner_database "parishioner_management/internal/databases/parishioner"
	database_util "parishioner_management/internal/utils/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (mongo *mongoStore) GetParishionerListHasPagination(
	ctx context.Context,
	filters map[string]interface{},
	paging *common.Paging,
	withDeleted bool,
	moreKeys ...interface{},
) ([]parishioner_database.Model, error) {
	parishionerCollection := database_util.GetCollection(mongo.db, database_model_const.Parishioner)
	filter := bson.M{}

	// setup pagination for
	findOptions := options.Find()
	findOptions.SetLimit(int64(paging.Limit))
	findOptions.SetSkip(database_util.CalculateSkipItem(paging.Page, paging.Limit))

	for key, value := range filters {
		filter[key] = value
	}

	if !withDeleted {
		filter[database_field_const.DeletedAt] = nil
	}

	// count parishioner
	totalCount, err := parishionerCollection.CountDocuments(ctx, filter)

	if err != nil {
		return nil, err
	}

	paging.Total = totalCount
	result := make([]parishioner_database.Model, 0)
	cursor, err := parishionerCollection.Find(ctx, filter, findOptions)

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		//Create a value into which the single document can be decoded
		var parishionerData parishioner_database.Model
		if err := cursor.Decode(&parishionerData); err != nil {
			return nil, err
		}

		result = append(result, parishionerData)
	}

	return result, nil
}

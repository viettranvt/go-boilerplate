package parishioner_storage

import (
	"context"
	database_field_const "parishioner_management/internal/constant/database/field"
	database_model_const "parishioner_management/internal/constant/database/model"
	database_util "parishioner_management/internal/utils/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (mongo *mongoStore) UpdateParishioner(
	ctx context.Context,
	id primitive.ObjectID,
	dataUpdate map[string]interface{},
) error {
	data := bson.M{}

	for key, value := range dataUpdate {
		data[key] = value
	}

	parishionerCollection := database_util.GetCollection(mongo.db, database_model_const.Parishioner)
	if _, err := parishionerCollection.UpdateOne(ctx, bson.M{database_field_const.ID: id}, data); err != nil {
		return err
	}

	return nil
}

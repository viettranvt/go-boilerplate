package parishioner_storage

import (
	"context"
	database_model_const "parishioner_management/internal/constant/database/model"
	parishioner_database "parishioner_management/internal/databases/parishioner"
	database_util "parishioner_management/internal/utils/database"
)

func (mongo *mongoStore) InsertParishioner(
	ctx context.Context,
	model *parishioner_database.Model,
) error {
	parishionerCollection := database_util.GetCollection(mongo.db, database_model_const.Parishioner)
	if _, err := parishionerCollection.InsertOne(ctx, model); err != nil {
		return err
	}

	return nil
}

package database_util

import (
	database_model_const "parishioner_management/internal/constant/database/model"

	"go.mongodb.org/mongo-driver/mongo"
)

// this function will get collection from client of mongo db
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	database := client.Database(database_model_const.DatabaseName)
	accountCollection := database.Collection(collectionName)

	if accountCollection == nil {
		panic("get collection failed")
	}

	return accountCollection
}

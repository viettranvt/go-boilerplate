package scripts

import (
	"context"
	"log"
	"os"
	configs "parishioner_management/internal/configs"
	database_field_const "parishioner_management/internal/constant/database/field"
	database_model_const "parishioner_management/internal/constant/database/model"
	mongo_field_const "parishioner_management/internal/constant/mongo/field"
	account_model "parishioner_management/internal/models/account"
	database_util "parishioner_management/internal/utils/database"
	hash_util "parishioner_management/internal/utils/hash"
	phone_util "parishioner_management/internal/utils/phone"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateAdminAccount() {
	log.Println("create account admin")

	dbUrl := os.Getenv(configs.EnvMongoDBUrl)

	if dbUrl == "" {
		log.Fatalln("Can't load mongo database url")
	}

	// connect mongo db
	ctx := context.Background()
	mongoDBClient, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))

	if err != nil {
		log.Fatalln(err)
	}

	// when func run finish this func will close the connection of mongo
	defer func() {
		if err := mongoDBClient.Disconnect(ctx); err != nil {
			log.Fatalln(err)
		}
	}()

	accountCollection := database_util.GetCollection(mongoDBClient, database_model_const.Account)
	filter := bson.M{database_field_const.Username: "admin"}
	var data account_model.Model

	if err := accountCollection.FindOne(ctx, filter).Decode(&data); err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("create new account")
			passwordHash, err := hash_util.Hash("12345678")

			if err != nil {
				log.Fatalln(err)
			}

			model := &account_model.Model{
				UserName: "admin",
				FullName: "admin",
				Phone:    phone_util.RemovePrefix("0328839669"),
				Email:    "admin.thanhtuan@gmail.com",
				Password: passwordHash,
			}

			if _, err = accountCollection.InsertOne(ctx, model); err != nil {
				log.Fatalln(err)
			}

			log.Println("create account successful")

			return
		}

		log.Fatalln(err)
	}

	// record has been deleted
	if !data.DeletedAt.IsZero() {
		// remove time of deletedAt field
		filter := bson.M{database_field_const.ID: data.ID}
		update := bson.M{mongo_field_const.Set: bson.M{database_field_const.DeletedAt: nil}}

		if _, err := accountCollection.UpdateOne(ctx, filter, update); err != nil {
			log.Fatalln(err)
		}
	}

	log.Println("create account successful")
}

package components

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainMySqlDBConnection() *gorm.DB
}

type appCtx struct {
	MySqlDB *gorm.DB
	MongoDB *mongo.Client
}

func NewAppContext(mySqlDB *gorm.DB, mongoDB *mongo.Client) *appCtx {
	return &appCtx{MySqlDB: mySqlDB, MongoDB: mongoDB}
}

func (ctx *appCtx) GetMainMySqlDBConnection() *gorm.DB {
	return ctx.MySqlDB
}

func (ctx *appCtx) GetMainMongoDBConnection() *mongo.Client {
	return ctx.MongoDB
}

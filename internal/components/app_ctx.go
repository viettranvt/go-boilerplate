package components

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type AppContext interface {
	GetMainMongoDBConnection() *mongo.Client
}

type appCtx struct {
	MongoDB *mongo.Client
}

func NewAppContext(mongoDB *mongo.Client) *appCtx {
	return &appCtx{MongoDB: mongoDB}
}

func (ctx *appCtx) GetMainMongoDBConnection() *mongo.Client {
	return ctx.MongoDB
}

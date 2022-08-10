package components

import (
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainMySqlDBConnection() *gorm.DB
}

type appCtx struct {
	MySqlDB *gorm.DB
}

func NewAppContext(mySqlDB *gorm.DB) *appCtx {
	return &appCtx{MySqlDB: mySqlDB}
}

func (ctx *appCtx) GetMainMySqlDBConnection() *gorm.DB {
	return ctx.MySqlDB
}

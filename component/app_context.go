package component

import (
	"elements-service/component/uploadprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetUploadProvider() uploadprovider.UploadProvider
	GetSecretKey() string
}

type appCtx struct {
	db         *gorm.DB
	upProvider uploadprovider.UploadProvider
	secretKey  string
}

func NewAppContext(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string) *appCtx {
	return &appCtx{db: db, upProvider: upProvider, secretKey: secretKey}
}

func (ctx appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx appCtx) GetUploadProvider() uploadprovider.UploadProvider {
	return ctx.upProvider
}

func (ctx appCtx) GetSecretKey() string {
	return ctx.secretKey
}

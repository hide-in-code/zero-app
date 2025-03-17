package svc

import (
	"zero-app/app/auth/api/internal/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DbEngin *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	// gorm支持
	db, err := gorm.Open(sqlite.Open(c.Sqlite.DataSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//todo: migrate

	return &ServiceContext{
		Config:  c,
		DbEngin: db,
	}
}

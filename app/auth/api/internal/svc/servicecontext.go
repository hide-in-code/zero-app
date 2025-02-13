package svc

import (
	"zero-app/app/auth/api/internal/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ServiceContext struct {
	Config  config.Config
	DbEngin *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	// gorm支持
	db, err := gorm.Open("mysql", c.Mysql.DataSource)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:  c,
		DbEngin: db,
	}
}

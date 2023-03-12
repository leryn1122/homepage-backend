package database

import (
	"fmt"
	"github.com/leryn1122/homepage/v2/pkg/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DbClient *gorm.DB

func init() {
	dbConfig := &config.Configuration.DbConfig
	db, err := Initialize(dbConfig)
	if err != nil {
		logrus.Fatalf("err: %v", err)
	}
	DbClient = db
}

func Initialize(config *config.DbConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.Scheme)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

package dao

import (
	"gin_ranking/config"
	logger "gin_ranking/pkg"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	db, err := gorm.Open(mysql.Open(config.MysqlDB), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 解决插入表的时候会自动添加复数的问题，如：user变成users
			SingularTable: true,
		},
	})

	if err != nil {
		logger.Error(map[string]interface{}{"mysql connection error": err.Error()})
	}

	if db.Error != nil {
		logger.Error(map[string]interface{}{"database error": db.Error})
	}

	DB, err := db.DB()
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)
	DB.SetConnMaxLifetime(time.Hour) // 60 mins
}

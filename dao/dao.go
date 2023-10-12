package dao

import (
	"fmt"
	"gin_ranking/config"
	logger "gin_ranking/pkg"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var (
	DBEngine  *gorm.DB
	err error
)

func init() {

	// attention: 初始化中使用 = 运算符，以确保它们成为全局变量，并且可以被其他包正常访问
	DBEngine, err = gorm.Open(mysql.Open(config.MysqlDB), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 解决插入表的时候会自动添加复数的问题，如：user变成users
			SingularTable: true,
		},
	})

	if err != nil {
		logger.Error(map[string]interface{}{"mysql connection error": err.Error()})
	}

	if DBEngine.Error != nil {
		logger.Error(map[string]interface{}{"database error": DBEngine.Error})
	}

	db, _ := DBEngine.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour) // 60 mins

	curDBName := DBEngine.Migrator().CurrentDatabase()
	fmt.Println("[D] current database name: ", curDBName) // crud_list
}

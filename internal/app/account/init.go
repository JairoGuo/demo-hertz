package account

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dsn = "root:root@tcp(localhost:3306)/hertz?charset=utf8&parseTime=True&loc=Local"

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
}

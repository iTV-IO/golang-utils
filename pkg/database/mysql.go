package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(conf Config) *gorm.Dialector {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.USER,
		conf.PASSWORD,
		conf.HOST,
		conf.PORT,
		conf.DATABASE,
	)

	msql := mysql.Open(dsn)
	return &msql
}

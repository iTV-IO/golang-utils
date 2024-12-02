package database

import (
	"errors"

	"gorm.io/gorm"
)

func NewDatabase(conf Config) (*gorm.DB, error) {
	switch conf.DRIVER {
	case "mysql":
		db := NewMysql(conf)
		return gorm.Open(*db, &gorm.Config{})
	case "postgres":
		db := NewPostgres(conf)
		return gorm.Open(*db, &gorm.Config{})
	default:
		return nil, errors.New("database driver not support")
	}
}

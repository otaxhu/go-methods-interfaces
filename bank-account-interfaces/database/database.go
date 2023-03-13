package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewTransaction() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(os.Getenv("dsn")))
	if err != nil {
		return nil, err
	}
	return db, nil
}

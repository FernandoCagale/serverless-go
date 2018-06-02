package datastore

import (
	"github.com/FernandoCagale/serverless-go/src/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func New(connection string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)
	db.SingularTable(true)
	db.AutoMigrate(&models.Task{})

	return db, nil
}

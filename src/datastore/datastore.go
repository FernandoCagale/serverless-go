package datastore

import (
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

	return db, nil
}

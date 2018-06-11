package datastore

import (
	"errors"
	"net/http"

	"github.com/gorilla/context"
	"github.com/jinzhu/gorm"
)

func GetConnection(r *http.Request) (*gorm.DB, error) {
	if rv := context.Get(r, "db"); rv != nil {
		conn := rv.(*gorm.DB)
		if err := conn.DB().Ping(); err != nil {
			return nil, errorConnection()
		}
		return conn, nil
	}
	return nil, errorConnection()
}

func errorConnection() error {
	return errors.New("Connection invalid")
}

package utils

import (
	"errors"
	"net/http"

	"github.com/gorilla/context"
	"github.com/jinzhu/gorm"
)

func GetConnection(r *http.Request) (*gorm.DB, error) {
	if rv := context.Get(r, "db"); rv != nil {
		return rv.(*gorm.DB), nil
	}
	return nil, errors.New("Connection invalid")
}

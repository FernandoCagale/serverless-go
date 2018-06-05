package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *Auth) BeforeSave(scope *gorm.Scope) (err error) {
	password := []byte(u.Password)
	if pw, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost); err == nil {
		scope.SetColumn("password", pw)
		return nil
	}
	return err
}

func (u *Auth) ValidatePassword(password string) (valid bool) {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return false
	}
	return true
}

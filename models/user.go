package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string
	Email      string
	Password   string
	Username   string
	FirebaseId string
	Blocked    bool
}

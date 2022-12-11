package models

import "gorm.io/gorm"

type Records struct {
	gorm.Model
	price       int
	description string `gorm:"type:text"`
	UserID      int
	User        User
	CategoryId  int
	Category    Category
	Image       string
}

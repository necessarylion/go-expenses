package main

import (
	"github.com/jinzhu/gorm"
)

// Up is executed when this migration is applied
func Up_20221211141146(txn *gorm.DB) {
	type Category struct {
		gorm.Model
		Name string `gorm:"type:varchar(225)"`
	}
	txn.AutoMigrate(&Category{})
}

// Down is executed when this migration is rolled back
func Down_20221211141146(txn *gorm.DB) {
	txn.DropTable("categories")
}

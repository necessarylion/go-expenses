package main

import (
	"github.com/jinzhu/gorm"
)

// Up is executed when this migration is applied
func Up_20221211140519(txn *gorm.DB) {
	type Record struct {
		gorm.Model
		Price       int
		Description string
		UserID      int
		CategoryId  int
		Image       string `gorm:"type:varchar(225)"`
	}
	txn.AutoMigrate(&Record{})
}

// Down is executed when this migration is rolled back
func Down_20221211140519(txn *gorm.DB) {
	txn.DropTable("records")
}

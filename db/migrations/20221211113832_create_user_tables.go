package main

import (
	"github.com/jinzhu/gorm"
)

// Up is executed when this migration is applied
func Up_20221211113832(txn *gorm.DB) {
	type User struct {
		gorm.Model
		Name       string `gorm:"type:varchar(225)"`
		Email      string `gorm:"type:varchar(225)"`
		Password   string `gorm:"type:varchar(225)"`
		Username   string `gorm:"type:varchar(225)"`
		FirebaseId string `gorm:"type:varchar(225)"`
		Blocked    bool
	}
	txn.AutoMigrate(&User{})
}

// Down is executed when this migration is rolled back
func Down_20221211113832(txn *gorm.DB) {
	txn.DropTable("users")
}

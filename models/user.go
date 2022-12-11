package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Password   string         `json:"password"`
	Username   string         `json:"username"`
	FirebaseId string         `json:"firebase_id"`
	Blocked    bool           `json:"blocked"`
}

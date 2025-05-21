package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Email   string `gorm:"unique;not null"`
	Phone   string
	Address string
	UserID  uint `json:"user_id"`
}

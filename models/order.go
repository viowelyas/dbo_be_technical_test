package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID uint      `gorm:"not null"`
	Customer   Customer  `gorm:"foreignKey:CustomerID"`
	OrderDate  time.Time `gorm:"not null"`
	Status     string    `gorm:"not null"`
	TotalPrice float64   `gorm:"not null"`
}

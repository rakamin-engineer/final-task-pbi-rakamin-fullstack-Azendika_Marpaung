package models

import (
	"time"
)

type User struct {
	ID        uint    `gorm:"primaryKey"`
	Username  string  `gorm:"unique;not null"`
	Email     string  `gorm:"unique;not null"`
	Password  string  `gorm:"not null"`
	Photos    []Photo `gorm:"constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Photo struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Caption   string
	PhotoUrl  string `gorm:"not null"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

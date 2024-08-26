package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	IsAdmin   int    `gorm:"default:2"` // 1: Admin, 2: User Biasa
	CreatedAt time.Time
	Posts     []Post `gorm:"foreignKey:UserID"`
}

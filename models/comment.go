package models

import (
	"time"
)

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"type:text;not null"`
	CreatedAt time.Time
	UserID    uint
	User      User `gorm:"foreignKey:UserID"`
	PostID    uint
	Post      Post `gorm:"foreignKey:PostID"`
}

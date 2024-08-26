package models

import (
	"time"
)

type Post struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Content     string `gorm:"type:text;not null"`
	PublishedAt time.Time
	UserID      uint
	User        User `gorm:"foreignKey:UserID"`
	CategoryID  uint
	Category    Category  `gorm:"foreignKey:CategoryID"`
	Comments    []Comment `gorm:"foreignKey:PostID"`
	Tags        []Tag     `gorm:"many2many:post_tags;"`
}

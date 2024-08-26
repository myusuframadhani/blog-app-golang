package models

type Category struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"unique;not null"`
	Description string `gorm:"type:text"`
	Posts       []Post `gorm:"foreignKey:CategoryID"`
}

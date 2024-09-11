package model

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string `gorm:"type:text"`
	UserID      uint
	User        User `gorm:"foreignKey:UserID"`
}

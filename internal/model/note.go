package model

type Note struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"type:text"`
}

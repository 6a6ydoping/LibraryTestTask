package models

type Book struct {
	ID    uint `gorm:"primaryKey"`
	Title string
	Genre string
	ISBN  string
}

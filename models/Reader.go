package models

type Reader struct {
	ID       uint `gorm:"primaryKey"`
	FullName string
}

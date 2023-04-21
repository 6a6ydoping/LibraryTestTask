package models

type Author struct {
	ID             uint `gorm:"primaryKey"`
	FullName       string
	Alias          string
	Specialization string
}

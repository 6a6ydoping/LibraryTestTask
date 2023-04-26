package models

type Author struct {
	ID             uint `gorm:"primaryKey"`
	FullName       string
	Alias          string
	Specialization string
	Books          []Book `gorm:"many2many:author_books;"`
}

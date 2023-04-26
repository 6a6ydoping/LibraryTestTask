package models

type Book struct {
	ID       uint `gorm:"primaryKey"`
	AuthorID uint `json:"author_id"`
	Title    string
	Genre    string
	ISBN     string
}

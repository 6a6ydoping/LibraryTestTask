package models

type BorrowedBooks struct {
	ID       uint `gorm:"primaryKey"`
	ReaderID uint
	BookID   uint
}

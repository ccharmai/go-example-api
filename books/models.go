package books

import "go-example-api/common"

// Database and response model
type BookModel struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// Request model with validation
type CreateBookModel struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// Request converter to database model
func (parent *CreateBookModel) getBookModel() *BookModel {
	return &BookModel{Title: parent.Title, Content: parent.Content}
}

// Get all books from database
func GetAllBooks(b *[]BookModel) error {
	return common.DB.Find(&b).Error
}

// Create new book in database
func BookCreate(b *BookModel) error {
	return common.DB.Create(b).Error
}

// Apply migrations from main.go while init server
func ApplyMigrations() {
	common.DB.AutoMigrate(&BookModel{})
}

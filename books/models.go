package books

import "go-example-api/common"

// Define models
type BookModel struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type PostBookModel struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// Define converters
func (parent *PostBookModel) getBookModel() *BookModel {
	return &BookModel{Title: parent.Title, Content: parent.Content}
}

// Define database methods
func getAllMethod(b *[]BookModel) error {
	return common.DB.Find(b).Error
}

func createMethod(b *BookModel) error {
	return common.DB.Create(b).Error
}

// Define module migrations
func ApplyMigrations() {
	common.DB.AutoMigrate(&BookModel{})
}

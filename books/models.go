package books

import "api/common"

type BookModel struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func GetAllBooks(b *[]BookModel) error {
	return common.DB.Find(&b).Error
}

func BookCreate(b *BookModel) error {
	return common.DB.Create(b).Error
}

// Apply migrations from main.go while init server
func ApplyMigrations() {
	common.DB.AutoMigrate(&BookModel{})
}

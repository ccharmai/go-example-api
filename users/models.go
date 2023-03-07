package users

import (
	"go-example-api/common"
	"go-example-api/helpers"
)

// Database and response model
type UserModel struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Email        string `json:"email" gorm:"unique"`
	PasswordHash string `json:"-"`
}

// Request model with validation
type CreateUserModel struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

// Request converter to database model
func (parent *CreateUserModel) getUserModel() *UserModel {
	passwordHash := helpers.CryptPassword(parent.Password)
	return &UserModel{Email: parent.Email, PasswordHash: passwordHash}
}

// Create new user in database
func UserCreate(u *UserModel) error {
	return common.DB.Create(u).Error
}

// Apply migrations from main.go while init server
func ApplyMigrations() {
	common.DB.AutoMigrate(&UserModel{})
}


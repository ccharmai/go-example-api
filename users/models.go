package users

import (
	"go-example-api/common"
	"go-example-api/helpers"
)

// Define models
type UserModel struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Email        string `json:"email" gorm:"unique"`
	PasswordHash string `json:"-"`
}

type UserInputModel struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

type UserLoginModel struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Define methods
func (u *UserInputModel) getUserModelWithPasswordHash() *UserModel {
	passwordHash := helpers.HashPassword(u.Password)
	return &UserModel{Email: u.Email, PasswordHash: passwordHash}
}

func (u *UserModel) checkPassword(password string) bool {
	return helpers.CompareHashAndPlaintextPassword(u.PasswordHash, password)
}

func (u *UserModel) createJWTTokens(ipAddr string) (string, string) {
	accessPayload := helpers.AccessTokenPayload{
		Email: u.Email,
	}
	refreshPayload := helpers.RefreshTokenPayload{
		Email:  u.Email,
		IpAddr: ipAddr,
	}

	accessToken := helpers.GenerateAccessToken(accessPayload)
	refreshToken := helpers.GenerateRefreshToken(refreshPayload)

	return accessToken, refreshToken
}

// Define database methods
func findUserByEmailMethod(u *UserModel, email string) error {
	return common.DB.First(u, "email = ?", email).Error
}

func createMethod(u *UserModel) error {
	return common.DB.Create(u).Error
}

// Apply migrations from main.go while init server
func ApplyMigrations() {
	common.DB.AutoMigrate(&UserModel{})
}

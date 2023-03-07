package users

import (
	"go-example-api/common"
	"go-example-api/helpers"
)

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

func (parent *UserInputModel) getUserModelForUserCreate() *UserModel {
	passwordHash := helpers.CryptPassword(parent.Password)
	return &UserModel{Email: parent.Email, PasswordHash: passwordHash}
}

func UserCreate(u *UserModel) error {
	return common.DB.Create(u).Error
}

func (u *UserModel) CreateAccessToken() string {
	return helpers.GenerateAccessToken(u.Email)
}

func (u *UserModel) CreateRefreshToken(ipAddr string) string {
	payloadRaw := map[string]interface{}{
		"email": u.Email,
		"ip":    ipAddr,
	}
	var payload interface{} = payloadRaw
	return helpers.GenerateRefreshToken(payload)
}

func (u *UserModel) CreateTokens(ipAddr string) (string, string) {
	accessToken := u.CreateAccessToken()
	refreshToken := u.CreateRefreshToken(ipAddr)
	return accessToken, refreshToken
}

func GetUserWithEmail(u *UserModel, email string) error {
	return common.DB.First(&u, "email = ?", email).Error
}

func (u *UserModel) CheckUserPassword(password string) bool {
	return helpers.ComparePasswords(u.PasswordHash, password)
}

// Apply migrations from main.go while init server
func ApplyMigrations() {
	common.DB.AutoMigrate(&UserModel{})
}

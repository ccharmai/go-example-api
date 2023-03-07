package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func postNewUserController(c *gin.Context) {
	var userRequest UserInputModel

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user = userRequest.getUserModelForUserCreate()

	if err := UserCreate(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func userLoginController(c *gin.Context) {
	var userRequest UserLoginModel

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user UserModel

	errorAuthUserMessage := "Email or password are invalid"

	if err := GetUserWithEmail(&user, userRequest.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errorAuthUserMessage})
		return
	}

	if status := user.CheckUserPassword(userRequest.Password); !status {
		c.JSON(http.StatusBadRequest, gin.H{"error": errorAuthUserMessage})
		return
	}

	accessToken, refreshToken := user.CreateTokens(c.ClientIP())

	c.JSON(http.StatusOK, gin.H{"accessToken": accessToken, "refreshToken": refreshToken})
}

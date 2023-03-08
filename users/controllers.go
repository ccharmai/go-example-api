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

	var user = userRequest.getUserModelWithPasswordHash()

	if err := createMethod(user); err != nil {
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

	if err := findUserByEmailMethod(&user, userRequest.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errorAuthUserMessage})
		return
	}

	if success := user.checkPassword(userRequest.Password); !success {
		c.JSON(http.StatusBadRequest, gin.H{"error": errorAuthUserMessage})
		return
	}

	accessToken, refreshToken := user.createJWTTokens(c.ClientIP())
	c.JSON(http.StatusOK, gin.H{"accessToken": accessToken, "refreshToken": refreshToken, "ip": c.ClientIP()})
}

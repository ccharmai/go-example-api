package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func postNewUserController(c *gin.Context) {
	var userRequest CreateUserModel

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user = userRequest.getUserModel()

	if err := UserCreate(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

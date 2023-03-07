package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllBooksController(c *gin.Context) {
	var books []BookModel

	if err := GetAllBooks(&books); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"books": books})
}

func postNewBookController(c *gin.Context) {
	var book BookModel

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := BookCreate(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"book": book})
}

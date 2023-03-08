package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllBooksController(c *gin.Context) {
	var books []BookModel

	if err := getAllMethod(&books); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"books": books})
}

func postNewBookController(c *gin.Context) {
	var bookRequest PostBookModel

	if err := c.ShouldBindJSON(&bookRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var book = bookRequest.getBookModel()

	if err := createMethod(book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"book": book})
}

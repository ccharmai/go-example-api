package books

import "github.com/gin-gonic/gin"

func AddRoutes(r *gin.Engine) {
	r.GET("books", getAllBooksController)
	r.POST("book", postNewBookController)
}

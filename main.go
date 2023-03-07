package main

import (
	"api/books"
	"api/common"

	"github.com/gin-gonic/gin"
)

func addRoutes(r *gin.Engine) {
	books.AddRoutes(r)
}

func applyMigrations() {
	books.ApplyMigrations()
}

func main() {
	common.InitDatabase()
	applyMigrations()

	r := gin.Default()
	addRoutes(r)

	r.Run(":8080")
}

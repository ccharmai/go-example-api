package main

import (
	"api/books"
	"api/common"
	"api/users"

	"github.com/gin-gonic/gin"
)

func addRoutes(r *gin.Engine) {
	books.AddRoutes(r)
	users.AddRoutes(r)
}

func applyMigrations() {
	books.ApplyMigrations()
	users.ApplyMigrations()
}

func main() {
	common.InitDatabase()
	applyMigrations()

	r := gin.Default()
	addRoutes(r)

	r.Run(":8080")
}

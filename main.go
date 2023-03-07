package main

import (
	"go-example-api/books"
	"go-example-api/common"
	"go-example-api/users"

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

package users

import (
	"go-example-api/middleware"

	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.Engine) {
	r.POST("user", postNewUserController)
	r.POST("login", userLoginController)
	r.GET("me", middleware.AuthRequired, meController)
}

package users

import "github.com/gin-gonic/gin"

func AddRoutes(r *gin.Engine) {
	r.POST("user", postNewUserController)
	r.POST("login", userLoginController)
}

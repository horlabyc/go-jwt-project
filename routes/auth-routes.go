package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/horlabyc/golang-jwt-project/controllers"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("auth/signup", controller.SignUp())
	router.POST("auth/login", controller.Login())
}

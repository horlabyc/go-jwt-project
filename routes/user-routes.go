package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/horlabyc/golang-jwt-project/controllers"
	middleware "github.com/horlabyc/golang-jwt-project/middlewares"
)

func UserRoutes(router *gin.Engine) {
	router.Use(middleware.Authenticate())
	router.POST("users", controller.GetUsers())
	router.GET("users/:userId", controller.GetUser())
}

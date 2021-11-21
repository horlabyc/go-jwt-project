package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/horlabyc/golang-jwt-project/controllers"
)

func UserRoutes(router *gin.Engine) {
	// router.POST("users", controller.GetUser())
	router.GET("users/:userId", controller.GetUser())
}

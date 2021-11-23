package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/horlabyc/golang-jwt-project/helpers"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
			c.Abort()
			return
		}
		claims, err := helpers.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("firstName", claims.Firstname)
		c.Set("lastName", claims.Lastname)
		c.Set("uid", claims.Uid)
		c.Set("userType", claims.Usertype)
		c.Next()
	}
}

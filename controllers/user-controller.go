package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/horlabyc/golang-jwt-project/database"
	"github.com/horlabyc/golang-jwt-project/helpers"
	"github.com/horlabyc/golang-jwt-project/models"
	"go.mongodb.org/mongo-driver/bson"
)

var userCollection = database.OpenCollection(database.Client, "users")

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")
		if err := helpers.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"userId": userId}).Decode(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, user)
	}
}

func GetUsers() {

}

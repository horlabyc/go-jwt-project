package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/horlabyc/golang-jwt-project/helpers"
	"github.com/horlabyc/golang-jwt-project/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func verifyPassword() {

}

func hashToken() {

}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		var count int64
		var validError error
		var user models.User
		defer cancel()
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validator.New().Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		count, validError = userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if validError != nil {
			log.Panic(validError)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking email"})
		}

		count, validError = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()
		if validError != nil {
			log.Panic(validError)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking phone"})
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email or phone already exists"})
		}
		time := time.Now().String()
		user.CreatedAt = &time
		user.UpdatedAt = &time
		user.Id = primitive.NewObjectID()
		userId := user.Id.Hex()
		user.UserId = &userId
		token, refreshToken, err := helpers.GenerateToken(*&user.Email, *&user.FirstName, *&user.LastName, *&user.Usertype, *&user.UserId)
		if err != nil {
			msg := fmt.Sprintf("Token could not be generated")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		}
		user.Token = &token
		user.RefreshToken = &refreshToken

		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("User could not be created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		}
		defer cancel()
		c.JSON(http.StatusOK, gin.H{"user": user, "resultInsertionNumber": resultInsertionNumber})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		// var user models.User

	}
}

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
	"golang.org/x/crypto/bcrypt"
)

func verifyPassword(hashedPassword string, password string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	valid := true
	errorMsg := ""
	if err != nil {
		errorMsg = fmt.Sprintf("email or password is incorrect")
		valid = false
	}
	return valid, errorMsg
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(hashedPassword)
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		// var count int64
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
		count, validError := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
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
		password := hashPassword(*user.Password)
		user.Password = &password
		time := time.Now().String()
		user.CreatedAt = &time
		user.UpdatedAt = &time
		user.Id = primitive.NewObjectID()
		userId := user.Id.Hex()
		user.UserId = &userId
		token, refreshToken, err := helpers.GenerateToken(user.Email, user.FirstName, user.LastName, user.Usertype, user.UserId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Token could not be generated"})
		}
		user.Token = &token
		user.RefreshToken = &refreshToken

		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User could not be created"})
		}
		defer cancel()
		c.JSON(http.StatusOK, gin.H{"user": user, "resultInsertionNumber": resultInsertionNumber})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		var user models.User
		var foundUser models.User
		defer cancel()
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email or password is incorrect"})
			return
		}
		passwordIsValid, msg := verifyPassword(*user.Password, *foundUser.Password)
		if passwordIsValid != true {
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}
		token, refreshToken, _ := helpers.GenerateToken(foundUser.Email, foundUser.FirstName, foundUser.LastName, foundUser.Usertype, foundUser.UserId)

		helpers.UpdateAllTokens(token, refreshToken, foundUser.UserId)
		e := userCollection.FindOne(ctx, bson.M{"userId": foundUser.UserId}).Decode(&foundUser)
		if e != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
			return
		}
		c.JSON(http.StatusOK, foundUser)
	}
}

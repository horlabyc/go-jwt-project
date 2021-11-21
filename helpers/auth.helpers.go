package helpers

import (
	"context"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/horlabyc/golang-jwt-project/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email     string
	Firstname string
	Lastname  string
	Usertype  string
	Uid       string
	jwt.StandardClaims
}

var secret = LoadEnv("SECRET_KEY")
var userCollection = database.OpenCollection(database.Client, "users")

func GenerateToken(email *string, firstName *string, lastName *string, usertype *string, userId *string) (token string, refreshToken string, err error) {
	tokenClaims := &SignedDetails{
		Email:     *email,
		Firstname: *firstName,
		Lastname:  *lastName,
		Usertype:  *usertype,
		Uid:       *userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims).SignedString([]byte(secret))
	if err != nil {
		log.Panic(err)
		return
	}
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secret))
	if err != nil {
		log.Panic(err)
		return
	}
	return token, refreshToken, err
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userId *string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	var updateData primitive.D
	updateData = append(updateData, bson.E{"token", signedToken})
	updateData = append(updateData, bson.E{"refreshToken", signedRefreshToken})
	time := time.Now().String()
	updateData = append(updateData, bson.E{"updatedAt", time})
	upsert := true
	options := options.UpdateOptions{
		Upsert: &upsert,
	}
	filter := bson.M{"userId": userId}
	_, err := userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"$set", updateData},
		},
		&options,
	)
	defer cancel()
	if err != nil {
		log.Panic(err)
		return
	}
	return

}

package helpers

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
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

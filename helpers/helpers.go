package helpers

import (
	"errors"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// use godot package to load/read the .env file and
// return the value of the key
func LoadEnv(envKey string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading env file")
	}
	return os.Getenv(envKey)
}

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("userType")
	err = nil
	if userType != role {
		err = errors.New("you are not authorized to access this resource")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("userType")
	uid := c.GetString("uid")
	err = nil
	if userType == "USER" && uid != userId {
		err = errors.New("you are not authorized to access this resource")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}

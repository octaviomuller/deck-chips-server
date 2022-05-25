package helpers

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/octaviomuller/deck-chips-server/database"
	"github.com/octaviomuller/deck-chips-server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EnvVarError(variableName string) {
	log.Fatal("You must set your '" + variableName + "' environmental variable.")
}

func GetSecret() string {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		log.Fatal("You must set your 'JWT_SECRET_KEY' environmental variable.")
	}

	return secret
}

func CreateAccessToken(user models.User) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"userId": user.Id})

	secret := GetSecret()

	tokenString, err := token.SignedString(secret)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &tokenString, nil
}

func ValidateAccessToken(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		secret := GetSecret()

		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIdString := claims["userId"]

		userId, err := primitive.ObjectIDFromHex(userIdString.(string))
		if err != nil {
			return nil, err
		}

		user, err := database.GetUserById(userId)
		if err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return nil, err
	}
}

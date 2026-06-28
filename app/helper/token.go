package helper

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GenerateToken(id string, role string) (string, error) {
	godotenv.Load()
	lifeSpan, err := strconv.Atoi(os.Getenv("LIFE_SPAN"))
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(lifeSpan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token_, _ := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	return token_, nil
}

func GetToken(c *fiber.Ctx) *string {
	bearer := c.Get("Authorization")
	if len(strings.Split(bearer, " ")) == 2 {
		token := strings.Split(bearer, " ")[1]
		return &token
	}
	return nil
}

func TokenValidator(c *fiber.Ctx) (*jwt.Token, error) {
	godotenv.Load()
	token := GetToken(c)
	if token == nil {
		return nil, fmt.Errorf("token not found")
	}
	return jwt.Parse(*token, func(token_ *jwt.Token) (interface{}, error) {
		if _, ok := token_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("token not valid")
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
}
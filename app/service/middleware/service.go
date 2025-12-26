package middleware

import (
	"fmt"
	"task-be/app/helper"
	"task-be/app/helper/blacklisttoken"
	"task-be/app/service/user"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type Service interface {
	Auth() fiber.Handler
	PermittedWithInstitution() fiber.Handler
}

type middleware struct {
	repository blacklisttoken.Repository
	userRepo   user.Repository
}

func NewMiddleware(repository blacklisttoken.Repository, userRepo user.Repository) *middleware {
	return &middleware{
		repository, userRepo,
	}
}

func (m *middleware) isTokenBlackList(c *fiber.Ctx) bool {

	token := helper.GetToken(c)

	return m.repository.GetByToken(*token) == nil

}

func (m *middleware) Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var message Message
		message.Message = "unauthorized"
		token, err := helper.TokenValidator(c)

		if err != nil {
			message.Message = err.Error()
			return c.Status(403).JSON(message)
		}
		if m.isTokenBlackList(c) {
			message.Message = "Token has blacklist"
			return c.Status(403).JSON(message)
		}

		claims, _ := token.Claims.(jwt.MapClaims)
		c.Set("id", claims["id"].(string))
		userId := fmt.Sprintf("%v", claims["id"])
		c.Locals("id",userId)
		c.Next()
		return nil
	}
}
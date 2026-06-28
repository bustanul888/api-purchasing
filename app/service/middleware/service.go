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
	IsAdmin() fiber.Handler
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

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(403).JSON(message)
		}

		var userId, role string
		if claims["id"] != nil {
			userId = fmt.Sprintf("%v", claims["id"])
		}
		if claims["role"] != nil {
			role = fmt.Sprintf("%v", claims["role"])
		}

		c.Set("id", userId)
		c.Set("role", role)
		c.Locals("id", userId)
		c.Locals("role", role)
		c.Next()
		return nil
	}
}

func (m *middleware) IsAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var message Message
		message.Message = "unauthorized"
		roleVal := c.Locals("role")
		if roleVal == nil {
			return c.Status(403).JSON(message)
		}
		role, ok := roleVal.(string)
		if !ok || role != "admin" {
			return c.Status(403).JSON(message)
		}
		return c.Next()
	}
}
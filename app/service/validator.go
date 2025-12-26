package service

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()
func BindAndValidate(c *fiber.Ctx, payload interface{}) bool {
	if err := c.BodyParser(payload); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return true
	}

	if err := validate.Struct(payload); err != nil {
		errors := []map[string]string{}

		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, map[string]string{
				"field": e.Field(),
				"rule":  e.Tag(),
			})
		}

		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "validation error",
			"errors":  errors,
		})
		return true
	}

	return false
}

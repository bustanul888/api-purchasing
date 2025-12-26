package service

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func JSON(c *fiber.Ctx, err error, data interface{}, responStatus ...int) error {
	if len(responStatus) > 0 && responStatus[0] != 500 {
		if err != nil {
			return c.Status(responStatus[0]).JSON(fiber.Map{"message": err.Error()})
			
		}
		return c.Status(responStatus[0]).JSON(data)
	} else if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "uuups..."})
		
	}
	return c.Status(http.StatusOK).JSON(data)
}

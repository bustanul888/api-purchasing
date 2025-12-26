package helper

import (
	"github.com/gofiber/fiber/v2"
)

type TypeOfMessage struct {
	Create fiber.Map
	Update fiber.Map
	Delete fiber.Map
}

var GetMessage = TypeOfMessage{
	Create: fiber.Map{"message": "saved"},
	Update: fiber.Map{"message": "updated"},
	Delete: fiber.Map{"message": "deleted"},
}
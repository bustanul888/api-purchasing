package auth

import (
	"task-be/app/helper/blacklisttoken"
	"task-be/app/service"
	"task-be/app/service/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(app *fiber.App, db *gorm.DB,middleware fiber.Handler) {
	blackListRepo := blacklisttoken.NewRepository(db)	
	userRepo := user.NewRepository(db)
	service_ := NewService(blackListRepo,userRepo)
	route := app.Group("/auth")

	route.Post("", func(c *fiber.Ctx) error {
		var request authRequest
		if service.BindAndValidate(c, &request) {
			return nil
		}
		response, status, err := service_.login(request)
		return service.JSON(c, err, response, status)
	})
	route.Get("/logout", middleware, func(c *fiber.Ctx) error {
		err := service_.logout(c)
		return service.JSON(c, err, fiber.Map{"message": "logout"})
	})

}
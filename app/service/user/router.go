package user

import (
	"task-be/app/helper"
	"task-be/app/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(app *fiber.App, db *gorm.DB, auth fiber.Handler, admin fiber.Handler) {
	Repository := NewRepository(db)
	service_ := NewService(Repository)
	route := app.Group("/user")

	route.Post("", auth, admin, func(c *fiber.Ctx) error{
		var req userRequest
		if service.BindAndValidate(c,&req){
			return nil
		}
		err := service_.create(req)
		return service.JSON(c,err,helper.GetMessage.Create)
	})

	route.Put("/my-profile", auth, func(c *fiber.Ctx) error{
		id := c.Locals("id").(string)
		var req myProfileUpdateRequest
		if service.BindAndValidate(c,&req){
			return nil
		}
		status,err := service_.myProfileUpdate(id,req)
		return service.JSON(c,err,helper.GetMessage.Update,status)
	})

	route.Get("/my-profile", auth, func(c *fiber.Ctx) error{
		id := c.Locals("id").(string)
		data := service_.getById(id)
		return c.Status(200).JSON(data)
	})

	route.Put("/:id", auth, admin, func(c *fiber.Ctx) error{
		id := c.Params("id")
		var req userUpdateRequest
		if service.BindAndValidate(c,&req){
			return nil
		}
		err := service_.update(id,req)
		return service.JSON(c,err,helper.GetMessage.Update)
	})

	route.Delete("/:id", auth, admin, func(c *fiber.Ctx) error{
		id := c.Params("id")
		err := service_.delete(id)
		return service.JSON(c,err,helper.GetMessage.Delete)
	})

	route.Get("", auth, admin, func(c *fiber.Ctx) error{
		data := service_.getAll()
		return c.Status(200).JSON(data)
	})

	route.Get("/:id", auth, admin, func(c *fiber.Ctx) error{
		id := c.Params("id")
		data := service_.getById(id)
		return c.Status(200).JSON(data)
	})
}
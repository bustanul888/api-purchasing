package supplier

import (
	"fmt"
	"task-be/app/helper"
	"task-be/app/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(app *fiber.App, db *gorm.DB) {
	Repository := NewRepository(db)
	service_ := NewService(Repository)
	route := app.Group("/supplier")

	route.Post("",func(c *fiber.Ctx) error{
		var req supplierRequest
		if service.BindAndValidate(c,&req){
			return nil
		}
		fmt.Println(req)
		err := service_.create(req)
		return service.JSON(c,err,helper.GetMessage.Create)
	})

	route.Put("/:id",func(c *fiber.Ctx) error{
		id := c.Params("id")
		var req supplierRequest
		if service.BindAndValidate(c,&req){
			return nil
		}
		err := service_.update(id,req)
		return service.JSON(c,err,helper.GetMessage.Update)
	})

	route.Delete("/:id",func(c *fiber.Ctx) error{
		id := c.Params("id")
		err := service_.delete(id)
		return service.JSON(c,err,helper.GetMessage.Delete)
	})

	route.Get("",func(c *fiber.Ctx) error{
		data := service_.getAll()
		return c.Status(200).JSON(data)
	})

	route.Get("/:id",func(c *fiber.Ctx) error{
		id := c.Params("id")
		data := service_.getById(id)
		return c.Status(200).JSON(data)
	})
}
package purchasing

import (
	"fmt"
	"task-be/app/helper"
	"task-be/app/service"
	"task-be/app/service/item"
	"task-be/app/service/purchasingdetail"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(app *fiber.App, db *gorm.DB, middleware fiber.Handler) {
	Repository := NewRepository(db)
	PurchasingDetailRepository := purchasingdetail.NewRepository(db)
	ItemRepository := item.NewRepository(db)
	service_ := NewService(Repository,PurchasingDetailRepository,ItemRepository)
	route := app.Group("/purchasing",middleware)

	route.Post("",func(c *fiber.Ctx) error{
		var req purchasingRequest
		userId := c.Locals("id").(string)
		fmt.Println(userId)
		if service.BindAndValidate(c,&req){
			return nil
		}
		err := service_.create(userId,req)
		return service.JSON(c,err,helper.GetMessage.Create)
	})

	route.Put("/:id",func(c *fiber.Ctx) error{
		id := c.Params("id")
		var req updatePurchasingRequest
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

	route.Get("/dashboard",func(c *fiber.Ctx) error{
		data := service_.dashboard()
		return c.Status(200).JSON(data)
	})
}
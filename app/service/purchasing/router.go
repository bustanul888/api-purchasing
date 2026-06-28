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

func Router(app *fiber.App, db *gorm.DB, auth fiber.Handler, admin fiber.Handler) {
	Repository := NewRepository(db)
	PurchasingDetailRepository := purchasingdetail.NewRepository(db)
	ItemRepository := item.NewRepository(db)
	service_ := NewService(Repository,PurchasingDetailRepository,ItemRepository)
	route := app.Group("/purchasing", auth)

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

	route.Put("/:id", admin, func(c *fiber.Ctx) error{
		id := c.Params("id")
		var req updatePurchasingRequest
		if service.BindAndValidate(c,&req){
			return nil
		}
		err := service_.update(id,req)
		return service.JSON(c,err,helper.GetMessage.Update)
	})

	route.Delete("/:id", admin, func(c *fiber.Ctx) error{
		id := c.Params("id")
		err := service_.delete(id)
		return service.JSON(c,err,helper.GetMessage.Delete)
	})

	route.Get("",func(c *fiber.Ctx) error{
		data := service_.getAll()
		return service.JSON(c, nil, data)
	})

	route.Get("/dashboard",func(c *fiber.Ctx) error{
		var startPtr, endPtr *string
		if start := c.Query("start_date"); start != "" {
			startPtr = &start
		}
		if end := c.Query("end_date"); end != "" {
			endPtr = &end
		}
		data, err := service_.dashboard(startPtr, endPtr)
		return service.JSON(c, err, data)
	})

	route.Get("/:id",func(c *fiber.Ctx) error{
		id := c.Params("id")
		data := service_.getById(id)
		return service.JSON(c, nil, data)
	})

	route.Get("/detail/:id",func(c *fiber.Ctx) error{
		id := c.Params("id")
		data := service_.getDetailById(id)
		return service.JSON(c, nil, data)
	})
}
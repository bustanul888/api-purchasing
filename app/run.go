package app

import (
	"fmt"
	"os"
	"task-be/app/db"
	"task-be/app/helper/blacklisttoken"
	"task-be/app/service/auth"
	"task-be/app/service/item"
	"task-be/app/service/middleware"
	"task-be/app/service/purchasing"
	"task-be/app/service/supplier"
	"task-be/app/service/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)	
func Run(){
	db := db.Connection()
	middleware := middleware.NewMiddleware(blacklisttoken.NewRepository(db),user.NewRepository(db))
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	auth.Router(app,db,middleware.Auth())
	supplier.Router(app,db,middleware.Auth(),middleware.IsAdmin())
	purchasing.Router(app,db,middleware.Auth(),middleware.IsAdmin())
	item.Router(app,db,middleware.Auth(),middleware.IsAdmin())
	user.Router(app,db,middleware.Auth(),middleware.IsAdmin())
	
	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	app.Listen(fmt.Sprintf(":%v",port))
}

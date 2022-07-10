package routes

import (
	"github.com/aryaputra777/rest/controller"
	"github.com/aryaputra777/rest/middleware"
	"github.com/gofiber/fiber/v2"
)

func Listinghandler(app *fiber.App) {

	app.Post("/login", controller.Login)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Rest API with Golang")
	})

	//user
	app.Delete("/user/:id", middleware.JWTProtected(), controller.Deleteuser)
	app.Get("/user/:id", middleware.JWTProtected(), controller.Detailuser)
	app.Get("/users", middleware.JWTProtected(), controller.Getdatauser)
	app.Post("/user", middleware.JWTProtected(), controller.Savedatauser)
	app.Put("/user/:id", middleware.JWTProtected(), controller.Updateuser)
}

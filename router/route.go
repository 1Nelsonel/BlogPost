package router

import (
	"github.com/1Nelsonel/BlogPost/controler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", controler.ListPost)
	app.Post("/create/post", controler.CreatePost)
	app.Put("/update/post/:id", controler.UpdatePost)
	app.Delete("/delete/post/:id", controler.DeletePost)
}
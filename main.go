package main

import (
	"github.com/1Nelsonel/BlogPost/database"
	"github.com/1Nelsonel/BlogPost/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init()  {
	database.ConnectDB()
}

func main()  {

	sqlDb, err := database.DBConn.DB()

	if err != nil {
		panic("Error in connectio")
	}

	defer sqlDb.Close()

	app := fiber.New()

	// Initialize default config
	app.Use(logger.New())

	app.Use(cors.New())
	
	router.SetupRoutes(app)

	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
	 return c.SendStatus(404) // => 404 "Not Found"
	})

	app.Listen(":3000")
	
}
package controler

import "github.com/gofiber/fiber/v2"

func ListPost(c *fiber.Ctx) error {
	return c.SendString("No Post found for now")	
}

func CreatePost(c *fiber.Ctx) error {
	return c.SendString("Cannot Create Post For Now")
}

func UpdatePost(c *fiber.Ctx) error {
	return c.SendString("Cannot Update Post For Now")
}

func DeletePost(c *fiber.Ctx) error {
	return c.SendString("Cannot Delete Post For Now")
}
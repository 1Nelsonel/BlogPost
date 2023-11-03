package controler

import (
	"log"

	"github.com/1Nelsonel/BlogPost/database"
	"github.com/1Nelsonel/BlogPost/model"
	"github.com/gofiber/fiber/v2"
    "github.com/go-playground/validator/v10"
)

// ==============================AUTHOR========================
// list author
func ListAuthors(c *fiber.Ctx) error {
	context := fiber.Map{
		"statustText": "0k",
		"msg":         "Blog List",
	}
	db := database.DBConn
	var Author []model.Author
	db.Find(&Author)

	context["author_records"] = Author

	c.SendStatus(200)

	return c.Render("authors", context)

}

// create author
func CreateAuthor(c *fiber.Ctx) error {
    context := fiber.Map{
        "statusText": "0k",
        "msg":        "Create Author",
    }

    if c.Method() == "POST" {
        // Parse and validate the request body
        authorRequest := new(model.Author)
        if err := c.BodyParser(authorRequest); err != nil {
            log.Println("Error in parsing request")
            context["statusText"] = ""
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        }

        // Validate the request data
        validate := validator.New()
        if err := validate.Struct(authorRequest); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        }

        db := database.DBConn
        record := model.Author{
            Name:  authorRequest.Name,
            Email: authorRequest.Email,
        }

        result := db.Create(&record)

        if result.Error != nil {
            log.Println("Error saving the data")
            context["statusText"] = ""
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error saving the author"})
        }

        context["msg"] = "Author is created"
        context["data"] = record
        c.SendStatus(fiber.StatusCreated)
        return c.Redirect("/authors")
    }

    // Render the HTML form when the request is not a POST request
    return c.Redirect("/authors")
}
// func CreateAuthor(c *fiber.Ctx) error {
// 	context := fiber.Map{
// 		"statusText": "0k",
// 		"msg":        "Create Author",
// 	}

// 	if c.Method() == "POST" {
// 		db := database.DBConn
// 		record := new(model.Author)

// 		if err := c.BodyParser(&record); err != nil {
// 			log.Println("Error in parsing request")
// 			context["statusText"] = ""
// 		}

// 		result := db.Create(record)

// 		if result.Error != nil {
// 			log.Println("Error saving the data")
// 			context["statusText"] = ""
// 		} else {
// 			// Author successfully created, redirect to "/authors"
// 			return c.Redirect("/authors")
// 		}

// 		context["msg"] = "Author is created"
// 		context["data"] = record
// 		c.SendStatus(201)
// 		return c.JSON(context)
// 	}
//     return c.Redirect("/authors")
// }


// update author
func UpdateAuthor(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"msg":        "Update Author",
	}
	db := database.DBConn

	// Get the author ID from the URL params
	authorID := c.Params("id")

	var author model.Author

	// Find the author by ID
	result := db.First(&author, authorID)
	if result.Error != nil {
		context["statusText"] = "Not Found"
		context["msg"] = "Author not found"
		c.SendStatus(404)
		return c.JSON(context)
	}

	// Parse the request body into a new author object
	updatedAuthor := new(model.Author)
	if err := c.BodyParser(&updatedAuthor); err != nil {
		log.Println("Error in parsing request")
		context["statusText"] = "Bad Request"
		context["msg"] = "Invalid request data"
		c.SendStatus(400)
		return c.JSON(context)
	}

	// Update the author's fields
	author.Name = updatedAuthor.Name
	author.Email = updatedAuthor.Email

	// Save the updated author
	db.Save(&author)

	context["msg"] = "Author is updated"
	context["data"] = author
	c.SendStatus(200)
	return c.JSON(context)
}

// delete author
func DeleteAuthor(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"msg":        "Delete Author",
	}
	db := database.DBConn

	// Get the author ID from the URL params
	authorID := c.Params("id")

	var author model.Author

	// Find the author by ID
	result := db.First(&author, authorID)
	if result.Error != nil {
		context["statusText"] = "Not Found"
		context["msg"] = "Author not found"
		c.SendStatus(404)
		return c.JSON(context)
	}

	// Delete the author
	db.Delete(&author)

	context["msg"] = "Author is deleted"
	c.SendStatus(204)
	return c.JSON(context)
}

// =============================================BLOG POST
func ListPost(c *fiber.Ctx) error {
	context := fiber.Map{
		"statustText": "0k",
		"msg":         "Blog List",
	}

	db := database.DBConn

	// var Author []model.Author
	var authors []model.Author
	var categories []model.Category
	var blogs []model.Blog

	db.Find(&authors)
	db.Find(&categories)
	db.Find(&blogs)

	context["author_records"] = authors
	context["category_records"] = categories
	context["blog_records"] = blogs

	c.SendStatus(200)
	return c.Render("index", context)
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

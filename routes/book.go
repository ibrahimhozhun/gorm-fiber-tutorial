package routes

import (
	"github.com/gofiber/fiber/v2"
	. "github.com/ibrahimhozhun/gorm-fiber-tutorial/database"
	"gorm.io/gorm"
)

// Struct for converting request body
type request struct {
	Name   string  `json:"name"`
	Author string  `json:"author"`
	Price  float32 `json:"price"`
}

func GetBooks(c *fiber.Ctx, db *gorm.DB) error {
	var books []Book

	// Get all books
	db.Find(&books)

	return c.JSON(fiber.Map{"books": books})
}

func GetSingleBook(c *fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id")

	var book Book

	// Find the book with the given id
	db.Find(&book, id)

	// If the book is found, return it
	if book.ID != 0 {
		return c.JSON(fiber.Map{"book": book})
	}

	// Otherwise, return an error
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book can't be found"})
}

func CreateBook(c *fiber.Ctx, db *gorm.DB) error {
	var body request

	// Convert request body
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot create book"})
	}

	// Create a new book
	db.Create(&Book{Name: body.Name, Author: body.Author, Price: body.Price})

	// Return succes message
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": true})
}

func UpdateBook(c *fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id")

	var body request

	// Convert request body
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse values"})
	}

	var book Book

	// Find the book with the given id
	tx := db.First(&book, id)

	// If the book is found update the book
	if book.ID != 0 {
		tx.Updates(Book{
			Model:  gorm.Model{},
			Name:   body.Name,
			Author: body.Author,
			Price:  body.Price,
		})

		return c.JSON(fiber.Map{"success": true})
	}

	// Otherwise return an error message
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book can't be found"})
}

func DeleteBook(c *fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id")

	var book Book

	// Find the book with the given id
	tx := db.First(&book, id)

	// If the book is found delete it
	if book.ID != 0 {
		tx.Delete(&Book{}, id)

		return c.JSON(fiber.Map{"success": true})
	}

	// Otherwise return an error message
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book can't be found"})
}

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ibrahimhozhun/gorm-fiber-tutorial/database"
	"github.com/ibrahimhozhun/gorm-fiber-tutorial/routes"
	"gorm.io/gorm"
)

// Setup routes and pass `db` to handlers so they can access the database
func setupRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/books", func(c *fiber.Ctx) error {
		return routes.GetBooks(c, db)
	})
	app.Get("/books/:id", func(c *fiber.Ctx) error {
		return routes.GetSingleBook(c, db)
	})
	app.Post("/books/create", func(c *fiber.Ctx) error {
		return routes.CreateBook(c, db)
	})
	app.Put("/books/update/:id", func(c *fiber.Ctx) error {
		return routes.UpdateBook(c, db)
	})
	app.Delete("/books/delete/:id", func(c *fiber.Ctx) error {
		return routes.DeleteBook(c, db)
	})
}

func main() {
	// Set default flags to log on output in date-time-filename format
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Initialize Fiber app
	app := fiber.New()

	// Initialize logger
	app.Use(logger.New())

	// Open database connection
	db, err := database.Open()

	// Close database connection when app exits
	defer database.Close(db)

	if err != nil {
		log.Printf("Can not connect to database\nError=> %v", err.Error())
	} else {
		// Setup routes
		setupRoutes(app, db)

		// Start listening on http://localhost:5000
		app.Listen("localhost:5000")
	}

}

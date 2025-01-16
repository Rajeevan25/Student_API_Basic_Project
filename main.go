package main

import (
	database "fiber_api/Database"
	"fiber_api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)
import "github.com/gofiber/fiber/v2/middleware/cors"
func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API")
}
func setupRoutes(app *fiber.App)  {
	app.Use(cors.New())
	app.Get("/api",welcome)

	app.Post("/api/students",routes.CreateStudent)
	app.Get("/api/students", routes.GetStudents)
	app.Get("/api/students/:id", routes.GetStudent)
	app.Delete("/api/students/:id", routes.DeleteStudent)
	app.Put("api/students/:id",routes.UpdateStudent)

}
func main() {
	database.ConnectDB()
	app := fiber.New()

	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
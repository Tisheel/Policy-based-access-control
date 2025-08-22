package main

import (
	"pbac/db"
	"pbac/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {

	db.InitDB()
	app := fiber.New()

	app.Post("/users", handlers.CreateUser)
	app.Post("/resources", handlers.CreateResource)
	app.Post("/actions", handlers.CreateAction)
	app.Post("/policies", handlers.CreatePolicy)
	app.Post("/check", handlers.CheckAccess)

	app.Get("/users", handlers.ListUsers)
	app.Get("/resources", handlers.ListResources)
	app.Get("/actions", handlers.ListActions)
	app.Get("/policies", handlers.ListPolicies)

	app.Listen(":8080")

}

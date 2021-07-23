package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/renatospaka/go-fiber/user"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Welcome to Fiber Http")
}

func Routers(app *fiber.App) {
	app.Get("/users", user.GetUsers)
	app.Get("/users/:id", user.GetUser)
	app.Post("/user", user.SaveUser)
	app.Put("/user/:id", user.UpdateUser)
	app.Delete("/users/:id", user.DeleteUser)
}

func main() {
	user.InitialMigration()

	app := fiber.New()
	app.Get("/", hello)
	Routers(app)

	app.Listen(":3000")
}

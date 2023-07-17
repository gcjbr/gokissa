package initializers

import (
	"github.com/gofiber/fiber/v2"
)

func SetRoutes() {
	App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("root")
	})

}

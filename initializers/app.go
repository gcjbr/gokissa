package initializers

import (
	"github.com/gofiber/fiber/v2"
)

var App *fiber.App

// Initializes the app server
func SetApp() {
	App = fiber.New()

}

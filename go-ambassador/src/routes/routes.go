package routes

import (
	"ambassador/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	user := api.Group("user")
	user.Post("register", controllers.Register)
}

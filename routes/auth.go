package routes

import (
	controller "github.com/Olapat/go-architecture/controllers/auth"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app fiber.Router) {
	api := app.Group("/auth")
	api.Post("/sign_in", controller.SignIn)
	api.Post("/reset_password/request", controller.RequestResetPassword)
	api.Get("/reset_password/verify/:token", controller.VerifyResetPassword)
	api.Patch("/reset_password/save", controller.ResetPassword)
}

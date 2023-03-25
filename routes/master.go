package routes

import (
	controller "github.com/Olapat/go-architecture/controllers/master"

	"github.com/gofiber/fiber/v2"
)

func MasterRoute(app fiber.Router) {
	api := app.Group("/master")
	api.Get("/province", controller.Province)
	api.Get("/district/:province_id", controller.District)
	api.Get("/sub_district/:province_id/:district_id", controller.SubDistrict)
}

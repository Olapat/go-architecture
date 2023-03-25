package main

import (
	"log"
	"os"

	"github.com/Olapat/go-architecture/db"
	_ "github.com/Olapat/go-architecture/docs" // you need to update github.com/rizalgowandy/go-swag-sample with your own project path
	"github.com/Olapat/go-architecture/routes"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/go-openapi/runtime/middleware"
	_ "github.com/go-openapi/spec"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

// @title Go-filber API
// @version 1.0
// @description This is a spec api Go-filber project
// @host localhost:9000
// @BasePath /api/v1
// @securityDefinitions.apiKey JWT
// @in                         header
// @name                       Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New(fiber.Config{
		ReadBufferSize: 4096 * 40,
	})

	app.Use(recover.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PATCH,DELETE,PUT",
	}))

	port := os.Getenv("PORT")

	db.Connect()
	app.Use(logger.New())
	app.Get("/", HealthCheck)

	app.Get("/swagger-doc.yaml", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/swagger.yaml")
	})

	opts1 := middleware.RedocOpts{SpecURL: "/swagger-doc.yaml", Path: "docs"}
	shrd := middleware.Redoc(opts1, nil)
	app.Get("/docs", adaptor.HTTPHandler(shrd))

	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger-doc.yaml", Path: "docs-swagger"}
	shsw := middleware.SwaggerUI(opts, nil)
	app.Get("/docs-swagger", adaptor.HTTPHandler(shsw))

	// PATH=$(go env GOPATH)/bin:$PATH
	// swag init --pd true
	app.Get("/swagger/*", swagger.HandlerDefault)
	api := app.Group("/api/v1")
	routes.RootRoute(api)

	app.Listen(port)
}

func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}

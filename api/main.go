package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Vijay-K-2003/go_urlshortner/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err.Error())
	}

	var ConfigDefault = cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "*",
		AllowCredentials: false,
		ExposeHeaders:    "*",
		MaxAge:           0,
	}

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(ConfigDefault))
	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}

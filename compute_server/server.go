package main

import (
	"compute_server/cache"
	"compute_server/database"
	"compute_server/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	// app.Use(cors.New(
	// 	cors.Config{
	// 		AllowOrigins: "http://localhost:3000",
	// 		AllowMethods: "GET,POST,PUT,DELETE",
	// 	}))
	database.ConnectDB()
	cache.Init()
	router.SetupRoutes(app)
	log.Fatal(app.Listen(":6944"))
}
package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/slaveofcode/okegani/repository/db"
	v1 "github.com/slaveofcode/okegani/v1"
)

func buildRoutes(app *fiber.App) {
	v1.BuildRoutes(app.Group("/v1"))
}

func setMiddlewares(app *fiber.App) {
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(compress.New())
}

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader: "OkeGaNi",
		Prefork:      true,
	})

	setMiddlewares(app)
	buildRoutes(app)

	db.Init(db.WithEnvDBConfig())
	defer db.CloseDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3131"
	}

	host, _ := os.Hostname()
	log.Println("App started at " + host + ":" + port)
	err := app.Listen(":" + port)

	if err != nil {
		log.Fatalln("Unable to Start the App:", err.Error())
	}
}

package main

import (
	"log"

	"github.com/FarrellOkello/Go_course/FarrellOkello/Go_course/database"
	"github.com/FarrellOkello/Go_course/FarrellOkello/Go_course/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New())

	routes.Setup(app)
	log.Fatal(app.Listen(":4000"))
}

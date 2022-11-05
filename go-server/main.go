package main

import (
	route "go-server/src/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)


func main () {
    godotenv.Load()
    PORT := os.Getenv("PORT")

    app := fiber.New(); route.Setup(app)
    app.Listen(":"+PORT)
}

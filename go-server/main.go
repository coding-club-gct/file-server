package main

import (
	"go-server/database"
	route "go-server/src/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)


func main () {
    godotenv.Load()
    PORT := os.Getenv("PORT")
    database.Conn()
    app := fiber.New(); route.Setup(app)
    app.Listen(":"+PORT)
}

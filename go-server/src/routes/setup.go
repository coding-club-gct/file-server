package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup(app *fiber.App) {
    app.Get("/", func (c *fiber.Ctx) error {
	return c.SendString("Hello, world!")
    })
    api := app.Group("/api", logger.New())
    store := api.Group("/store")
    
    HandleStore(store)
}

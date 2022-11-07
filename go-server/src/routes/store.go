package route

import (
	controller "go-server/src/controllers"

	"github.com/gofiber/fiber/v2"
)



func HandleStore (store fiber.Router) {
    store.Post("/addFiles", controller.AddFiles)
}

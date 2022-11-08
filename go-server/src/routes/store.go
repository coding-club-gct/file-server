package route

import (
	controller "go-server/src/controllers"

	"github.com/gofiber/fiber/v2"
)



func HandleStore (store fiber.Router) {

    // categories
    store.Get("/all-categories-full", func (c *fiber.Ctx) error {
	return controller.GetAllCategories(c, true)
    })
    store.Get("/all-categories", func (c *fiber.Ctx) error {
	return controller.GetAllCategories(c, false)
    })
    store.Post("/create-categories", controller.CreateCategories)

    //files
    store.Get("/all-files", controller.GetFiles)
    store.Post("/add-files", controller.AddFiles)
}

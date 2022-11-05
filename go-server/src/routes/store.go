package route

import (
	model "go-server/models"

	"github.com/gofiber/fiber/v2"
)


func HandleStore (store fiber.Router) {
    store.Post("/", func (c *fiber.Ctx) error {
	cat := new(model.Category)
	if err := c.BodyParser(cat); err != nil {
	    return err
	}
	return c.JSON(cat)
    })        
}

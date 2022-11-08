package controller

import (
	"go-server/database"
	model "go-server/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllCategories (c *fiber.Ctx, full bool) error {
    var categories []model.Category
    var result *gorm.DB
    if full {
	result = database.Database.Db.Preload("Files").Find(&categories)
    } else {
	result = database.Database.Db.Find(&categories)
    }
    if result.Error != nil {
	return result.Error
    }
    return c.JSON(categories)
}

func CreateCategories (c *fiber.Ctx) error {  
    var datas []map[string]string
    if err := c.BodyParser(&datas); err != nil {
        return err
    } 
    categories := make([]model.Category, 0)
    for _, data := range datas {
	categories = append(categories, model.Category{
	    Name: data["name"],
	    Banner: data["banner"],
	})
    }
    result := database.Database.Db.Create(&categories)
    if result.Error != nil {
        return result.Error
    }
   return c.JSON(categories)
}

func GetFiles (c *fiber.Ctx) error {
    var files []model.File
    result := database.Database.Db.Find(&files)
    if result.Error != nil {
	return result.Error
    }
    return c.JSON(files)
}

func AddFiles (c *fiber.Ctx) error {
    type payload struct {
	CatID uint `json:"cat_id"`
	Name string `json:"name"`
	Banner string `json:"banner"`
    }
    var datas []payload
    if err := c.BodyParser(&datas); err != nil {
	return err
    }
    files := make([]model.File, 0)
    for _, data := range datas {
	files = append(files, model.File{
	    CatID: data.CatID,
	    Name: data.Name,
	    Banner: data.Banner,
	})	
    }
    result := database.Database.Db.Create(&files)
    if result.Error != nil {
	return result.Error
    }
    return c.JSON(files)
}

func GetFilesUnderCat (c *fiber.Ctx) error {
    var categories []model.Category
    res := database.Database.Db.Preload("Files").Find(&categories)
    if res.Error != nil {
	return res.Error
    }
    return c.JSON(categories)
}















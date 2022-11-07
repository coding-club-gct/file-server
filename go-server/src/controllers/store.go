package controller

import (
	"go-server/database"

	"github.com/gofiber/fiber/v2"
)

type Category struct {
    CatName string `json:"catName"`
    CatBanner string `json:"catBanner"`
}

type File struct {
    FileName string `json:"fileName"`
    FilePath string `json:"filePath"`
    FileBanner string `json:"fileBanner"`
}

type AddFilesPayload struct {
    Category
    Files []File
}

func AddFiles (c *fiber.Ctx) error {  
    reqBody := new(AddFilesPayload) 
    if err := c.BodyParser(reqBody); err != nil {
        return err
    } 
    result := database.Database.Db.Create(&reqBody)
    if result.Error != nil {
        return result.Error
    }
    result = database.Database.Db.Create(&reqBody.Files)
    if result.Error != nil {
        return result.Error
    }
    return c.JSON(reqBody)
}

















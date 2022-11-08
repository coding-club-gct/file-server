package model

import "gorm.io/gorm"

type File struct {
    gorm.Model
    CatID uint `json:"cat_id"`
    Name string `json:"fileName"`
    Banner string `json:"fileBanner"`
}

package model

import "gorm.io/gorm"

type Category struct {
    gorm.Model
    CatName string `json:"catName" gorm:"unique"`
    CatBanner string `json:"catBanner"`
    Files []File `gorm:"foreignKey:FileName"`
}

package model

import "gorm.io/gorm"

type Category struct {
    gorm.Model
    Name string `json:"name" gorm:"unique"`
    Banner string `json:"banner"`
    Files []File `json:"files" gorm:"foreignKey:CatID"`
}

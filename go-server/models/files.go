package model

import "gorm.io/gorm"

type File struct {
    gorm.Model
    CatID uint
    Name string
}

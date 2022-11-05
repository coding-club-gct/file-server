package model

import "gorm.io/gorm"

type Category struct {
    gorm.Model
    Name string `json:"name"`
    Description string `json:"description"`
    UploadedBy string `json:"uploadedBy"`
    Banner string `json:"banner"`
}
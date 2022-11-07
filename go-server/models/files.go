package model

import "gorm.io/gorm"

type File struct {
    gorm.Model
    CatName string
    FileName string `json:"fileName"`
    FilePath string `json:"filePath"`
    FileBanner string `json:"fileBanner"`
}

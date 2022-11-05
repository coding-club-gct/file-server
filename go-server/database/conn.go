package database

import (
	model "go-server/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
    Db *gorm.DB
}

var Database DbInstance

func Conn () {
    db, err := gorm.Open(sqlite.Open("database/db/db.db"), &gorm.Config{})
    if err != nil {
	log.Fatalln(err)
    }
    db.Logger = logger.Default.LogMode(logger.Info)
    db.AutoMigrate(&model.Category{}, &model.File{})
    Database = DbInstance{Db: db}
}

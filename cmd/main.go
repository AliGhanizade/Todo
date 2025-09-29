package main

import (
	"todo/config"
	"todo/model"
	"todo/router"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	var err error
	config.Db, err = gorm.Open(sqlite.Open(config.DBName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	config.Db.AutoMigrate(&model.Task{})

	r := router.SetupRouter()

	r.Run(":9092")
}

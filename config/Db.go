package config

import "gorm.io/gorm"

var Db *gorm.DB

var DBName string = "../Todo.sqlite"
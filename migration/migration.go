package migration

import (
	"github.com/hesen/blog/entry/user"
	"github.com/hesen/blog/pool"

	"log"
)

// CreateTables CreateTable create table for models
func CreateTables() {
	pool.DB.CreateTable(&user.User{})

	log.Println(`==== @ complete tables creation ====`)
}

// Migrations AutoMigrate run auto migration for given models, will only add missing fields, won't delete/change current data
func Migrations() {
	pool.DB.AutoMigrate(&user.User{})

	log.Println(`==== @ finished tables migration ====`)
}

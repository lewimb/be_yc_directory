package main

import (
	"github/lewimb/be_yc_directory/config"
	"log"
)

func main() {
	var db config.Database
	db.Migrate()

	log.Print("Successfully Create tables")
}

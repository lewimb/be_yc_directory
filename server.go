package main

import (
	"database/sql"
	"fmt"
	database "github/lewimb/be_yc_directory/db"
	"github/lewimb/be_yc_directory/service/routes"
	"log"
	"net/http"
)

func main() {
	var db *sql.DB
	connection := &database.Connect{DB: db}

	db = connection.Open()

	defer connection.Close()

	fmt.Println("Connected!")
	log.Println("Server is starting at port 8080")

	if err := http.ListenAndServe(":8080", routes.RegisteredRoutes(db)); err != nil {
		log.Fatal(err)
	}

}

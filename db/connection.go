package db

import (
	"database/sql"
	"log"
)

type Connect struct {
	DB *sql.DB
}

func (conn *Connect) Open() *sql.DB {
	db, err := sql.Open("mysql", "root:Binusian1@/yc_directory")
	if err != nil {
		panic(err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Database not reachable:", err)
	}

	return db
}

func (conn *Connect) Close() error {
	err := conn.DB.Close()
	return err
}

package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	mysqlMigrate "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

type Database struct {
	DB       *sql.DB
	Cfg      mysql.Config
	DbDriver database.Driver
}

func (database *Database) GetConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	database.Cfg = mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("DBADDRESS"),
		DBName: os.Getenv("DBNAME"),
	}
}

func (database *Database) open() {
	database.GetConfig()
	var err error

	database.DB, err = sql.Open("mysql", database.Cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	if pingErr := database.DB.Ping(); pingErr != nil {
		log.Fatal(pingErr)
	}

}

func (database *Database) driver() {
	var err error
	database.open()

	database.DbDriver, err = mysqlMigrate.WithInstance(database.DB, &mysqlMigrate.Config{})

	if err != nil {
		log.Fatal(err)
	}
}

func (database *Database) Migrate() {
	database.driver()

	m, err := migrate.NewWithDatabaseInstance(
		"file://./db/tables",
		"mysql",
		database.DbDriver,
	)

	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}

func (database *Database) Rollback() {
	database.driver()

	m, err := migrate.NewWithDatabaseInstance(
		"file://./db/tables",
		"mysql",
		database.DbDriver,
	)

	if err != nil {
		log.Fatal(err)
	}

	if err := m.Down(); err != nil {
		log.Fatal(err)
	}
}

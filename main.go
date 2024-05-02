package main

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// InitDb initializes the database connection
func InitDb() *sql.DB {
	var err error

	// Connect to MariaDB assuming username/password is user
	Db, err := sql.Open("mysql", "user:user@tcp(localhost:3306)/our_database_name")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	return Db
}

// migrates the database to create tables
func UpdateDatabaseSchema() {
	// Initialize the database connection
	Db := InitDb()
	defer Db.Close() // Close the database connection when done

	// Create a new MySQL driver
	driver, _ := mysql.WithInstance(Db, &mysql.Config{})

	// Create a new migration instance
	m, _ := migrate.NewWithDatabaseInstance(
		"file://migrations", // Migration files location
		"our_database_name", // Current migration version
		driver,              // MySQL driver
	)

	// Run the migrations
	m.Up()
}

func main() {
	UpdateDatabaseSchema()
}

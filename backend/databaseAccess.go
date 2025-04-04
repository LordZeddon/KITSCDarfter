package backend

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"log"
)



func SetUpDB() {
	db, err := sql.Open("sqlite3", "file:./KITSCDrafterDB.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database successfully!")

	// Create a table
	createTables(db)

	//setting up Database scheme

	
}

func createTables(db *sql.DB) {
	_, err = db.Exec(	`CREATE TABLE IF NOT EXISTS test (
		id INT PRIMARY KEY,
		name VARCHAR(100)
	);`)
}


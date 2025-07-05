package config

import (
	"database/sql"
	"fmt" // Import fmt for printing
	"log"
	"os" // Import os for environment variables

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

var DB *sql.DB

func ConnectDB() {
	dsn := os.Getenv("MYSQL_DSN")
    // --- ADD THESE DEBUG LINES ---
	fmt.Printf("Attempting to connect with DSN: %s\n", dsn)
	if dsn == "" {
		log.Fatal("MYSQL_DSN environment variable not set.")
	}
    // --- END DEBUG LINES ---

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	// Ping the database to verify the connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err) // This is where your error currently comes from
	}

	log.Println("Conexión a la base de datos establecida correctamente!")
}
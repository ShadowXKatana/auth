package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func VerifyDatabaseConnection() {
	dsn := os.Getenv("DB_DSN")

	if dsn == "" {
		log.Fatal("Error: DB_DSN environment variable is not set")
	}

	// 2. เริ่มการเชื่อมต่อ
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// 3. ทดสอบการเชื่อมต่อจริง (Ping)
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Successfully connected to RDS PostgreSQL!")
}

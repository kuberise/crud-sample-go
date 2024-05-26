package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	connStr := os.Getenv("DB_CONN_STR")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS items (
        id SERIAL PRIMARY KEY,
        name TEXT,
        value TEXT
    );`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO items (name, value) VALUES
        ('item1', 'value1'),
        ('item2', 'value2')
    ON CONFLICT DO NOTHING;`)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database initialized")
}

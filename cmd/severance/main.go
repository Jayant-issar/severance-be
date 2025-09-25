package main

import (
	"database/sql"
	"log"

	// This is the blank import for the Postgres driver.
	// The underscore means we are importing it for its side effects only (the init function).
	_ "github.com/lib/pq"

	"github.com/Jayant-issar/severance-backend/internal/config"
	"github.com/Jayant-issar/severance-backend/internal/database/db"
	"github.com/Jayant-issar/severance-backend/internal/handler"
)

func main() {
	//load configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not lead config: %v", err)
	}
	//establish a conn to database.
	conn, err := sql.Open(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		log.Fatalf("cloud not connect to the database: %v", err)
	}
	defer conn.Close()

	//pinging database
	if err := conn.Ping(); err != nil {
		log.Fatalf("database is not reachable: %v", err)
	}

	log.Println("database connection established")

	//new db store
	store := db.NewStore(conn)

	//new server instance
	server := handler.NewServer(store)

	//start http request and listen for incoming requests.
	log.Printf("Server starting on %s", cfg.ServerAddress)
	if err := server.Start(cfg.ServerAddress); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}

package main

import (
	"flag"
	"log"
	"note-server/src/db"
	"os"
)

func main() {
	var (
		migrate = flag.Bool("migrate", false, "Run database migrations")
		seed    = flag.Bool("seed", false, "Seed the database with sample data")
		reset   = flag.Bool("reset", false, "Reset the database (drop and recreate all tables)")
	)
	flag.Parse()

	if err := db.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println("Error closing database:", err)
		}
	}()

	switch {
	case *reset:
		if err := db.DropAllTables(); err != nil {
			log.Fatal("Failed to drop all tables:", err)
		}
		if err := db.AutoMigrate(); err != nil {
			log.Fatal("Failed to run migrations:", err)
		}

		log.Println("Database reset completed")
		return

	case *migrate:
		if err := db.AutoMigrate(); err != nil {
			log.Fatal("Failed to run migrations:", err)
		}

		log.Println("Migrations completed")

	case *seed:
		if err := db.SeedDatabase(); err != nil {
			log.Fatal("Failed to seed database:", err)
		}

		log.Println("Database seeded")
		return
	}

	if len(os.Args) == 1 {
		if err := db.AutoMigrate(); err != nil {
			log.Fatal("Failed to run migrations:", err)
		}
		log.Println("Auto-migrations completed")
	}
}

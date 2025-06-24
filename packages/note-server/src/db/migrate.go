package db

import (
	"fmt"
	"log"

	"note-server/src/models"
)

func AutoMigrate() error {
	if DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	log.Println("Running auto-migrations...")

	if err := DB.AutoMigrate(&models.NoteModel{}); err != nil {
		return fmt.Errorf("failed to run auto-migrations: %w", err)
	}

	log.Println("Auto-migrations completed successfully")
	return nil
}

func DropAllTables() error {
	if DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	log.Println("Dropping all tables...")

	tables := []interface{}{
		&models.NoteModel{},
	}

	for _, table := range tables {
		if err := DB.Migrator().DropTable(table); err != nil {
			return fmt.Errorf("failed to drop table: %w", err)
		}
	}

	log.Println("All tables dropped successfully")
	return nil
}

func SeedDatabase() error {
	if DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	log.Println("Seeding database...")
	notes := []models.NoteModel{
		{
			Title:   "Title 1",
			Content: "Content 1",
		},
		{
			Title:   "Title 2",
			Content: "Content 2",
		},
	}

	for _, note := range notes {
		if err := DB.FirstOrCreate(&note, models.NoteModel{Title: note.Title}).Error; err != nil {
			return fmt.Errorf("failed to create note: %w", err)
		}
	}

	log.Println("Database seeded successfully")
	return nil
}

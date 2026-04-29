package main

import (
	"log"

	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/config"
	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/models"
	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/storage"

	"gorm.io/gorm"
)

func main() {
	if err := config.LoadAllEnv(); err != nil {
		log.Printf("Warning loading config: %v", err)
	}
	cfg, err := config.LoadNewBootCfg()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db := storage.InitDBConn(cfg)
	if db == nil {
		log.Fatal("Could not connect to database")
	}
	defer storage.CloseDBConn(db)

	log.Println("Starting GORM AutoMigration...")
	if err := storage.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Starting data seeding...")
	seedData(db)

	log.Println("Done!")
}

func seedData(db *gorm.DB) {
	// Инициализация базовых ролей
	roles := []models.Role{
		{ID: 1, Name: "Admin"},
		{ID: 2, Name: "User"},
	}

	for _, role := range roles {
		if err := db.FirstOrCreate(&role, models.Role{ID: role.ID}).Error; err != nil {
			log.Printf("Failed to seed role %s: %v", role.Name, err)
		}
	}
	log.Println("Roles seeded successfully.")

	// Инициализация базовых категорий
	categories := []models.Category{
		{ID: 1, Name: "Electronics", Description: "Electronic devices and accessories"},
		{ID: 2, Name: "Books", Description: "Various kinds of books"},
		{ID: 3, Name: "General", Description: "General items"},
	}

	for _, category := range categories {
		if err := db.FirstOrCreate(&category, models.Category{ID: category.ID}).Error; err != nil {
			log.Printf("Failed to seed category %s: %v", category.Name, err)
		}
	}
	log.Println("Categories seeded successfully.")
}

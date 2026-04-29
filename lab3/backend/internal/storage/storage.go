package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/config"
	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDBConn(config *config.BootConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
		return nil
	}

	sqlDB, err := db.DB()
	if err == nil {
		err = sqlDB.Ping()
		if err != nil {
			log.Fatal("failed to ping database: ", err)
			return nil
		}
	}

	log.Print("Successfully connected to the database!")
	return db
}

func RunMigrations(db *gorm.DB) error {
	log.Print("Running auto-migrations...")
	err := db.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Category{},
		&models.Product{},
	)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	log.Print("Migrations applied successfully.")
	return nil
}

func CloseDBConn(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

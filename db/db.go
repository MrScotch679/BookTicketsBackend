package db

import (
	"fmt"
	"github.com/MrScotch679/book-tickets-backend/config"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDb(config *config.EnvConfig, DBMigrator func(db *gorm.DB) error) *gorm.DB {
	uri := fmt.Sprintf(`
    host=%s user=%s password=%s dbname=%s sslmode=%s port=5432
`, config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBSSLMode)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	log.Info("Database connected")

	if err := DBMigrator(db); err != nil {
		log.Fatalf("Unable to migrate database: %v", err)
	}

	return db
}

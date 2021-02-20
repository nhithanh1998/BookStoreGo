package database

import (
	"bookstore/models"
	"bookstore/util"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// BookStoreDB is variable
var BookStoreDB *gorm.DB

// OpenConnectionToDatabase is function
func OpenConnectionToDatabase(databaseConfig *util.DatabaseConfig) error {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		databaseConfig.DatabaseHost, databaseConfig.DatabaseUsername, databaseConfig.DatabasePassword, databaseConfig.DatabaseName, databaseConfig.DatabasePort)
	var err error
	BookStoreDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

// MigrateModel is a function to migrate model
func MigrateModel() {
	BookStoreDB.AutoMigrate(&models.Genre{})
	genres := []models.Genre{{Name: "Romantic"}, {Name: "Drama"}}
	BookStoreDB.Create(&genres)
}

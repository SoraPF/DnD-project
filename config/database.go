package config

import (
	"DnD-project/helper"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB(config *Config) *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUsername, config.DBPassword, config.DBName)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{ /*Logger: logger.Default.LogMode(logger.Info)*/ })
	helper.ErrorPanic(err)

	fmt.Println("Connected successfully to the database!")
	return db
}

func AutoMigrate(DB *gorm.DB) error {
	// Migrate all tables
	if err := DB.AutoMigrate(
		//&model.Note{},
	); err != nil {
		return err
	}
	return nil
}
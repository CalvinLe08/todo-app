package initializers

import (
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var db *gorm.DB

func connectDB(config *Config) {
	var err error

	// https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL 
	dsn := fmt.Sprintf(`
		host=%s 
		user=%s 
		password=%s 
		dbname=%s 
		port=%s 
		sslmode=disable 
		TimeZone=Asia/Shanghai`,
		config.DBHost,
		config.DBUser,	
		config.DBPass,
		config.DBName,
		config.DBPort,
	)
	
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database.")
	}

	fmt.Println("Successfully connected to database")
}

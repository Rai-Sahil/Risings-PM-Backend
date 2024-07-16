package database

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"pm_backend/internals/config"
	"pm_backend/internals/models"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	if DB != nil {
		return DB, nil
	}

	cfg := config.GetConfig()
	connectionString := cfg.PostgresConnectionString

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Student{},
		&models.Calendar{},
		&models.Goal{},
		&models.Task{},
		&models.SubTask{},
		&models.Comment{},
		&models.Reminder{},
	)

	if err != nil {
		return nil, err
	}

	DB = db

	return db, nil
}

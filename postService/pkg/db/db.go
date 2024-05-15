package db

import (
	"fmt"
	"postservice/pkg/config"
	"postservice/pkg/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if dbErr != nil {
		return nil, dbErr
	}

	if err := db.AutoMigrate(&domain.Post{}); err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&domain.Upvote{}); err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&domain.Downvote{}); err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&domain.PostVote{}); err != nil {
		return nil, err
	}
	return db, nil
}

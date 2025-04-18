package db

import (
	"fmt"

	"example.com/review/v2/config"
	"example.com/review/v2/db/model"
	"example.com/review/v2/db/model/jobs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db connection failed: %w", err)
	}

	// auto migrate
	dbConn.AutoMigrate(&jobs.Jobs{})
	dbConn.AutoMigrate(&model.HotelReview{})

	return dbConn, nil
}

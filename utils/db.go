package utils

import (
	"fkedweb/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword)
	return gorm.Open("postgres", dsn)
}

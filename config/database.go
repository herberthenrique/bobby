package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// DBConfig Default struct for configuration
type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

func BuildDatabaseConfiguration() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DBName:   "example_db",
	}

	return &dbConfig
}

// DatabaseURl return formatted database UEL
func DabataseURL(dbConfig *DBConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.DBName, dbConfig.Password)
}

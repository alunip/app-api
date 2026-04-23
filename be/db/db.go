package db

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	Host            string
	Port            string
	User            string
	Password        string
	Database        string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func LoadConfig() *Config {
	return &Config{
		Host:            getEnv("DB_HOST", "localhost"),
		Port:            getEnv("DB_PORT", "5432"),
		User:            getEnv("DB_USER", "app_api"),
		Password:        getEnv("DB_PASSWORD", "changeme"),
		Database:        getEnv("DB_NAME", "app_api_db"),
		SSLMode:         getEnv("DB_SSLMODE", "disable"),
		MaxOpenConns:    getEnvInt("DB_MAX_OPEN_CONNS", 25),
		MaxIdleConns:    getEnvInt("DB_MAX_IDLE_CONNS", 5),
		ConnMaxLifetime: getEnvDuration("DB_CONN_MAX_LIFETIME", 5*time.Minute),
	}
}

func Connect(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.Database, config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying db: %w", err)
	}

	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)

	if err := sqlDB.Ping(); err != nil {
		sqlDB.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Printf("Database connection established (max: %d, idle: %d)", config.MaxOpenConns, config.MaxIdleConns)
	return db, nil
}

func Close() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err == nil {
			sqlDB.Close()
			log.Println("Database connection closed")
		}
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Invalid integer value for %s, using default %d", key, defaultValue)
		return defaultValue
	}
	return intValue
}

func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	duration, err := time.ParseDuration(value)
	if err != nil {
		log.Printf("Invalid duration value for %s, using default %v", key, defaultValue)
		return defaultValue
	}
	return duration
}

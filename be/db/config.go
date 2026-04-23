package db

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type AppConfig struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Version   string    `json:"version"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (AppConfig) TableName() string {
	return "app_config"
}

var ErrConfigNotFound = errors.New("config not found")

func GetAppConfig() (*AppConfig, error) {
	var config AppConfig
	result := DB.First(&config, 1)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrConfigNotFound
		}
		return nil, fmt.Errorf("query failed: %w", result.Error)
	}
	return &config, nil
}

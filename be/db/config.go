package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

// AppConfig represents the application configuration
type AppConfig struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Version   string    `json:"version"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var (
	ErrConfigNotFound = errors.New("config not found")
)

// GetAppConfig retrieves the singleton application configuration
func GetAppConfig(ctx context.Context) (*AppConfig, error) {
	query := `
		SELECT id, name, version, created_at, updated_at
		FROM app_config
		WHERE id = 1
	`

	var config AppConfig
	err := Pool.QueryRow(ctx, query).Scan(
		&config.ID, &config.Name, &config.Version,
		&config.CreatedAt, &config.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrConfigNotFound
		}
		return nil, fmt.Errorf("query failed: %w", err)
	}

	return &config, nil
}

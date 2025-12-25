CREATE TABLE app_config (
    id INTEGER PRIMARY KEY CHECK (id = 1),
    name VARCHAR(255) NOT NULL,
    version VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Insert initial configuration
INSERT INTO app_config (id, name, version) VALUES (1, 'app-api', '1.0.0');

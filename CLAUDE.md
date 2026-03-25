# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a full-stack web application with a Go backend (`be/`) and React TypeScript frontend (`fe/`). The application uses PostgreSQL for data persistence and Docker Compose for orchestration.

## Architecture

### Backend (Go)
- **Framework**: Standard library `net/http` with custom routing
- **Database**: PostgreSQL with pgx/v5 driver (connection pooling via pgxpool)
- **Migrations**: golang-migrate/migrate handles database schema versioning
- **Structure**:
  - `main.go`: Application entry point with server setup and graceful shutdown
  - `db/`: Database connection pool, configuration, and migration runner
  - `handlers/`: HTTP request handlers with standardized JSON responses
  - `middleware/`: CORS and other middleware
  - `migrations/`: SQL migration files (up/down pattern)
  - `models/`: Data models (currently minimal)

### Frontend (React + TypeScript)
- **Build Tool**: Vite with hot module replacement
- **HTTP Client**: Axios with interceptors configured in `services/api.ts`
- **Structure**:
  - `src/components/`: Reusable React components
  - `src/pages/`: Page-level components
  - `src/services/`: API communication layer
  - `src/hooks/`: Custom React hooks
  - `src/models/`: TypeScript type definitions
  - `src/routes/`: Routing configuration
  - `src/utils/`: Utility functions

### Key Design Patterns
- Backend uses standard HTTP handler functions with CORS middleware wrapping
- All API responses follow a consistent format: `{data, error, timestamp}`
- Database migrations run automatically on backend startup
- Backend implements graceful shutdown with signal handling (SIGINT/SIGTERM)
- Frontend uses environment variables for API URL configuration (`VITE_API_URL`)

## Development Commands

### Docker-based Development (Recommended)
```bash
# Start all services (backend, frontend, postgres)
docker-compose up

# Start in detached mode
docker-compose up -d

# View logs
docker-compose logs -f be-dev    # Backend logs
docker-compose logs -f fe-dev    # Frontend logs
docker-compose logs -f postgres  # Database logs

# Stop services
docker-compose down

# Rebuild containers after dependency changes
docker-compose up --build
```

### Backend (Go)
```bash
cd be

# Run locally (requires PostgreSQL)
go run .

# Build binary
go build -o app-api .

# Add dependencies
go get <package>
go mod tidy

# Run with Air (hot reload) - used by docker-compose
air -c .air.toml
```

### Frontend (React + TypeScript)
```bash
cd fe

# Install dependencies
npm install

# Development server (http://localhost:5173)
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview

# Lint code
npm run lint
```

### Database Migrations
Migrations run automatically on backend startup. To create new migrations:

```bash
cd be/migrations

# Create migration files (manual process)
# Follow naming: XXXXXX_description.up.sql and XXXXXX_description.down.sql
# Example: 000003_add_users_table.up.sql
```

## Environment Configuration

Copy `.env.example` to `.env` and adjust values as needed. Key variables:

**Backend:**
- `PORT`: Server port (default: 8080)
- `CORS_ORIGIN`: Allowed CORS origin
- `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`: PostgreSQL connection
- `DB_SSLMODE`: SSL mode for database (disable for local development)
- `DB_MAX_OPEN_CONNS`, `DB_MAX_IDLE_CONNS`: Connection pool settings

**Frontend:**
- `VITE_API_URL`: Backend API URL (e.g., http://localhost:8080)

## Production Deployment

```bash
# Build and run production containers
docker-compose -f docker-compose.prod.yml up -d

# Production backend runs on port 8080
# Production frontend (Nginx) runs on port 80
```

Note: The production setup does NOT include PostgreSQL in `docker-compose.prod.yml` - database must be configured separately.

## API Endpoints

- `GET /api/health` - Health check with database connectivity status
- `GET /api/config` - Fetch application configuration from database

All responses follow this structure:
```json
{
  "data": { ... },
  "error": "error message" | null,
  "timestamp": "2024-01-01T12:00:00Z"
}
```

## Important Notes

- Backend uses Air for hot reloading in development (configured via `.air.toml`)
- Frontend Vite dev server must bind to `0.0.0.0` inside Docker containers
- Database migrations are embedded and run automatically on startup
- Backend implements connection pooling with health checks and max connection limits
- All HTTP handlers should use context with timeout for database operations

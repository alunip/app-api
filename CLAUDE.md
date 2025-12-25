# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Structure

This is a full-stack car tracking application with:
- **Frontend (`fe/`)**: React 19 + TypeScript + Vite application
- **Backend (`be/`)**: Go 1.23 application (module: `andersonlira.com/base`)

### Frontend Architecture

The frontend follows a standard React project structure with organized directories:
- `src/components/` - Reusable React components
- `src/pages/` - Page-level components
- `src/routes/` - Routing configuration
- `src/hooks/` - Custom React hooks
- `src/services/` - API and external service integrations
- `src/utils/` - Utility functions and helpers

Currently these directories are empty placeholders, indicating early stage development.

## Development Commands

### Frontend (fe/)

Navigate to `fe/` directory before running these commands:

```bash
# Install dependencies (first time setup)
npm install

# Start development server with HMR
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview

# Lint code
npm run lint
```

The frontend build process runs TypeScript compilation (`tsc -b`) before Vite build.

### Backend (be/)

The backend uses Go 1.23. Standard Go commands apply:

```bash
# Run the application (from be/ directory)
go run .

# Build the application
go build

# Run tests
go test ./...

# Format code
go fmt ./...
```

## Technical Stack

### Frontend
- **React 19.2** with React DOM
- **Vite 7.2** for build tooling and dev server
- **TypeScript 5.9** for type safety
- **ESLint** with TypeScript support and React plugins
- Vite plugin using Babel for Fast Refresh

### Backend
- **Go 1.23** (module name: `andersonlira.com/base`)

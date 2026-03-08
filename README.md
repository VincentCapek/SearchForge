# SearchForge

SearchForge is a local search application built with a Go backend and a Vue 3 + TypeScript frontend.

The goal of the project is to provide a **SERP-style** (Search Engine Results Page) interface, with a backend API that fetches and formats search results, and a modern frontend that displays them cleanly.

## Features

- **Go** API for serving search results
- Frontend built with **Vue 3 + TypeScript**
- **Vite** integration for frontend development
- Styling planned with **UnoCSS**
- Clear separation between backend, business logic, and UI
- An ideal foundation for evolving the project into a real local search engine interface

## Tech Stack

### Backend
- Go
- Standard library HTTP server
- JSON API

### Frontend
- Vue 3
- TypeScript
- Vite
- UnoCSS
- pnpm

## Project Structure

```text
.
├── assets/              # Vue + Vite frontend
├── search/              # Search fetching / parsing logic
├── server/              # HTTP handlers and API
├── utils/               # Utility helpers
├── go.mod
└── main.go
```

> The exact structure may evolve slightly over time as the project is refactored.

## Installation

### 1. Clone the project

```bash
git clone <REPO_URL>
cd SearchForge
```

### 2. Install the Go backend

Make sure Go is installed:

```bash
go version
```

Then, from the project root:

```bash
go mod tidy
```

### 3. Install the frontend

The frontend is located in the `assets/` folder.

```bash
cd assets
pnpm install
```

## Run the project in development

### Backend

From the project root:

```bash
go run .
```

The backend will be available at:

```text
http://localhost:8080
```

### Frontend

In a second terminal:

```bash
cd assets
pnpm dev
```

The Vite frontend will be available at:

```text
http://localhost:5173
```

## Development Workflow

The frontend calls the Go API through routes such as:

```text
GET /api/search?q=luffy
```

In development, Vite can be configured with a proxy to automatically forward `/api` requests to `http://localhost:8080`.

## Main Endpoint

### Search

```http
GET /api/search?q=your-query
```

#### Example

```bash
curl "http://localhost:8080/api/search?q=luffy"
```

#### Example response

```json
[
  {
    "title": "Example result",
    "url": "https://example.com",
    "domain": "example.com",
    "snippet": "Example snippet"
  }
]
```

## Useful Scripts

### Backend

```bash
go run .
go test ./...
```

### Frontend

```bash
cd assets
pnpm dev
pnpm build
pnpm preview
```

## Project Goal

This project is intended as a foundation for building a modern, lightweight, and maintainable search interface, with:

- a Go API
- a Vue web interface
- a design system powered by UnoCSS
- an extensible architecture to later add more search engines, filters, or views

## Possible Roadmap

- improve the SERP UI
- add multiple search engines
- add filters
- improve the result fetching/parsing system
- add search history
- improve production mode to serve the frontend directly from Go

## Requirements

- Go installed
- Node.js installed
- pnpm installed

### Quick checks

```bash
go version
node -v
pnpm -v
```

## Notes

- The project is designed to evolve progressively.
- The backend and frontend are intentionally separated to keep the codebase maintainable.
- The name `SearchForge` can be changed depending on the final repository name.

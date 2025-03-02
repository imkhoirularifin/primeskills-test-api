# Primeskills Backend Test API

Simple Todo App API built with Golang and Gin.

## Table of Contents

- [Project structure](#project-structure)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [Documentation](#documentation)

## Project Structure

```bash
.
├── .github/
│   └── workflows/
│       ├── deploy.yml
│       └── migrate.yml
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── docs/
│   │   └── docs.controller.go
│   ├── domain/
│   │   ├── dto/
│   │   │   └── user.dto.go
│   │   ├── entity/
│   │   │   └── user.entity.go
│   │   └── interfaces/
│   │       └── user.interface.go
│   ├── infrastructure/
│   │   ├── container.go
│   │   ├── gin.go
│   │   └── gorm.go
│   ├── middleware/
│   │   └── auth.go
│   ├── user/
│   │   ├── user.controller.go
│   │   ├── user.repository.go
│   │   └── user.service.go
│   └── utilities/
│       └── validation.go
├── migrations/
│   └── sql_migration.sql
├── pkg/
│   └── xlogger/
│       └── xlogger.go
├── main.go
├── docker-compose.yml
├── Dockerfile
├── .gitignore
├── go.mod
└── go.sum
```

## Installation

### Prerequisites

- [Golang](https://go.dev/doc/install) (v1.23 or higher recommended)
- [Swaggo](https://github.com/swaggo/gin-swagger) (optional, for generating Swagger documentation)
- [Goose](https://github.com/pressly/goose) (optional, for database migration)

### Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/imkhoirularifin/primeskills-test-api
   cd primeskills-test-api
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

## Configuration

### Environment Variables

1. Create a `.env` file in the root directory of your project.
2. Add environment variables to the `.env` file based on the following template:

   | Name                             | Description                                            | Required | Default           |
   |----------------------------------|--------------------------------------------------------|----------|-------------------|
   | `HOST`                           | Server host                                            | No       | `localhost`       |
   | `PORT`                           | Port number for the server                             | No       | `3000`            |
   | `IS_DEVELOPMENT`                 | Application status                                     | No       | `true`            |
   | `DB_DRIVER`                      | Database Driver                                        | No       | `sqlite`          |
   | `DB_DSN`                         | Database Connection URL                                | No       | `fiile::memory:?` |
   | `JWT_SECRET_KEY`                 | JWT secret key                                         | Yes      |                   |
   | `JWT_EXPIRES_IN`                 | JWT expire duration                                    | No       | `24h`             |
   | `JWT_ISSUER`                     | JWT issuer                                             | No       | `localhost`       |
   | `GOOSE_DRIVER`                   | Database driver for goose migration                    | No       |                   |
   | `GOOSE_DBSTRING`                 | Database connection url                                | No       |                   |
   | `GOOSE_MIGRATION_DIR`            | Output migration dir                                   | No       | `./migrations`    |
   | `SWAGGER_HOST`                   | Swagger Host                                           | No       | `localhost:3000`  |
   | `GOOGLE_APPLICATION_CREDENTIALS` | Google ADC (Application Default Credentials) file path | No       | `""`              |

3. Update these values based on your setup (e.g., database credentials).

### Database Setup

Migrate database to the latest version:

```bash
goose up
```

## Running the Application

### Development Mode

Start the application in development mode with:

```bash
go run main.go
```

### Production Mode

Build the application with:

```bash
go build -o primeskills-test-api
```

## Documentation

This API is documented using Swagger. To access the documentation, navigate to `/api/v1/docs/index.html` in your browser (e.g., `http://localhost:3000/api/v1/docs/index.html`). The Swagger UI will be displayed with a list of available endpoints.

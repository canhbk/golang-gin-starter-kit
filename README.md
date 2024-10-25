# Golang Gin Starter Kit

A backend service built with Go using the Gin framework and following the MVC (Model-View-Controller) architectural pattern.

## Table of Contents

- [Golang Gin Starter Kit](#golang-gin-starter-kit)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Prerequisites](#prerequisites)
  - [Quick Start with Docker Compose](#quick-start-with-docker-compose)
  - [Docker Compose Services](#docker-compose-services)
    - [Main Services](#main-services)
      - [API Service](#api-service)
      - [MySQL Service](#mysql-service)
    - [Database Management Services](#database-management-services)
      - [Migration Service](#migration-service)
      - [Seed Service](#seed-service)
      - [Refresh Service](#refresh-service)
  - [Docker Compose Commands](#docker-compose-commands)
    - [Service Management](#service-management)
    - [Database Management](#database-management)
    - [Logs and Monitoring](#logs-and-monitoring)
    - [Container Management](#container-management)
  - [Traditional Installation](#traditional-installation)
  - [Project Structure](#project-structure)
    - [Directory Structure Explanation](#directory-structure-explanation)
      - [Core Directories](#core-directories)
      - [API Layer](#api-layer)
      - [Business Layer](#business-layer)
      - [Data Layer](#data-layer)
      - [Types and DTOs](#types-and-dtos)
      - [Support Directories](#support-directories)
    - [Key Design Principles](#key-design-principles)
  - [Database Setup](#database-setup)
    - [Prerequisites](#prerequisites-1)
    - [Database Management](#database-management-1)
  - [API Documentation](#api-documentation)
    - [Base URL](#base-url)
    - [Available Endpoints](#available-endpoints)
      - [Health Check](#health-check)
      - [User Management](#user-management)
  - [Error Handling](#error-handling)
  - [Development](#development)
    - [Generate Swagger Documentation](#generate-swagger-documentation)
    - [Running the Application](#running-the-application)
    - [Database Migrations](#database-migrations)
    - [Adding New Controllers](#adding-new-controllers)
  - [Database Migrations](#database-migrations-1)
  - [Testing](#testing)
  - [Contributing](#contributing)
    - [Commit Message Format](#commit-message-format)
  - [License](#license)
  - [Contact](#contact)

## Features

- MVC architecture
- RESTful API endpoints with versioning (/api/v1/...)
- MySQL database with GORM
- Database migrations and seeding
- Swagger API documentation
- Health check monitoring
- Environment-based configuration
- Docker support with multi-stage builds
- JWT authentication (coming soon)
- Clean and extensible structure

## Prerequisites

- Go 1.23.2
- Git
- MySQL 8.0 or higher
- Make (optional, for Makefile usage)
- Docker and Docker Compose (optional, for containerization)

## Quick Start with Docker Compose

1. Clone the repository

   ```bash
      git clone git@github.com:canhbk/golang-gin-starter-kit.git
      cd golang-gin-starter-kit
   ```

2. Configure environment variables

   ```bash
      cp .env.example .env
      # Edit .env file with your configurations
   ```

3. Start the services

   ```bash
      # Build and start API and MySQL services
      docker-compose up -d --build
   ```

4. Initialize the database

   ```bash
      # Run database migrations
      docker-compose --profile tools run migrate

      # Seed the database with initial data
      docker-compose --profile tools run seed
   ```

   The API will be available at <http://localhost:8080>

## Docker Compose Services

The application uses Docker Compose to manage multiple services:

### Main Services

#### API Service

- Main application service
- Built from local Dockerfile
- Exposes port 8080
- Connects to MySQL database

#### MySQL Service

- MySQL 8.0 database
- Persistent data storage
- Exposes port 3306
- Automatic initialization
- Health checking enabled

### Database Management Services

#### Migration Service

- Runs database migrations
- Part of the "tools" profile
- Only runs when explicitly called

#### Seed Service

- Seeds initial data
- Part of the "tools" profile
- Only runs when explicitly called

#### Refresh Service

- Combines migrations and seeds into one command
- Part of the "tools" profile
- Only runs when explicitly called

## Docker Compose Commands

### Service Management

```bash
    # Start all services
    docker-compose up -d

    # Build and start services
    docker-compose up -d --build

    # Stop all services
    docker-compose down

    # Stop services and remove volumes
    docker-compose down -v
```

### Database Management

```bash
    # Run database migrations
    docker-compose --profile tools run migrate

    # Seed the database
    docker-compose --profile tools run seed

    # Combined migration and seeding
    docker-compose --profile tools run migrate && \
    docker-compose --profile tools run seed
```

### Logs and Monitoring

```bash
  # View all logs
  docker-compose logs -f

  # View API service logs
  docker-compose logs -f api

  # View MySQL logs
  docker-compose logs -f mysql
```

### Container Management

```bash
    # List running containers
    docker-compose ps

    # Restart a specific service
    docker-compose restart api

    # Remove all containers
    docker-compose rm -f
```

## Traditional Installation

1. Clone the repository

   ```bash
   git clone git@github.com:canhbk/golang-gin-starter-kit.git
   cd golang-gin-starter-kit
   ```

2. Install dependencies

   ```bash
   go mod tidy
   ```

3. Install Swagger tools

   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

4. Configure environment variables

   ```bash
      cp .env.example .env
      # Edit .env file with your configurations
   ```

## Project Structure

```text
.
├── main.go                     # Application entry point
├── .env                        # Environment variables
├── .env.example               # Example environment variables
├── Dockerfile                 # Multi-stage Docker build file
├── docker-compose.yaml        # Docker Compose configuration
├── .dockerignore             # Docker ignore file
├── Makefile                   # Build and development commands
├── cmd/
│   └── db/
│       └── main.go            # Database CLI tool
├── config/
│   └── database.go            # Database configuration
├── controllers/
│   ├── health_controller.go   # Health check controller
│   └── v1/                    # Version 1 controllers
│       └── user_controller.go # User management
├── database/
│   ├── migration/
│   │   └── migration.go       # Database migrations
│   └── seeder/
│       └── seeder.go          # Database seeders
├── docs/                      # Swagger documentation
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── models/
│   └── user.go                # Database models
├── services/
│   └── v1/                    # Version 1 business logic
│       └── user_service.go
├── types/                     # API request/response types
│   └── v1/                    # Version 1 types
│       ├── common/            # Shared types
│       │   ├── error.go       # Common error responses
│       │   └── pagination.go  # Pagination types
│       └── user/              # User-related types
│           ├── request.go     # User request DTOs
│           └── response.go    # User response DTOs
├── middleware/                # Custom middleware
│   ├── auth.go               # Authentication middleware
│   └── logger.go             # Logging middleware
├── routes/
│   └── routes.go             # Route definitions
├── utils/                     # Utility functions
│   ├── validator.go          # Custom validators
│   └── helpers.go            # Helper functions
└── tests/                     # Test files
    ├── integration/          # Integration tests
    └── unit/                 # Unit tests
```

### Directory Structure Explanation

#### Core Directories

- `cmd/`: Contains executable applications
  - `db/`: Database management CLI tool
- `config/`: Configuration files and setup
  - Database configurations
  - Environment configurations
  - Other service configurations

#### API Layer

- `controllers/`: Request handlers
  - Organized by API version
  - Handles HTTP requests/responses
  - Input validation
  - Calls appropriate services
- `routes/`: Route definitions
  - API endpoint registration
  - Middleware attachment
  - Route grouping

#### Business Layer

- `services/`: Business logic
  - Organized by domain and version
  - Implements business rules
  - Handles data processing
  - Coordinates between different domains

#### Data Layer

- `models/`: Database models
  - Entity definitions
  - Database relationships
  - Model methods
- `database/`: Database management
  - `migration/`: Schema migrations
  - `seeder/`: Data seeders

#### Types and DTOs

- `types/`: Data Transfer Objects (DTOs)
  - Organized by version and domain
  - `common/`: Shared types across domains
  - Domain-specific request/response types
  - Input/Output data structures

#### Support Directories

- `middleware/`: Custom middleware
  - Authentication
  - Logging
  - Rate limiting
  - CORS
- `utils/`: Helper functions
  - Common utilities
  - Helper functions
  - Custom validators
- `docs/`: API documentation
  - Swagger files
  - API specifications
- `tests/`: Test files
  - Unit tests
  - Integration tests
  - Test utilities

### Key Design Principles

1. **Versioning**

   - All API-related code is versioned (`v1`, `v2`, etc.)
   - Enables smooth API evolution
   - Maintains backward compatibility

2. **Separation of Concerns**

   - Clear separation between layers
   - Each directory has a specific responsibility
   - Minimizes code coupling

3. **Domain-Driven Design**

   - Code organized by business domains
   - Each domain has its own types and logic
   - Clear boundaries between domains

4. **Type Safety**

   - Separate request/response types
   - Strong typing for API contracts
   - Clear data validation rules

5. **Maintainability**
   - Consistent file naming
   - Logical grouping of related code
   - Easy to locate and modify components

## Database Setup

### Prerequisites

Ensure you have MySQL installed and running. Create a database for the project:

```sql
CREATE DATABASE example CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### Database Management

Build the database CLI tool:

```bash
go build -o bin/db-cli cmd/db/main.go
```

Available commands:

```bash
# Run migrations
./bin/db-cli -migrate

# Seed the database with initial data
./bin/db-cli -seed

# Rollback migrations
./bin/db-cli -rollback

# Refresh database (rollback, migrate, and seed)
./bin/db-cli -refresh
```

## API Documentation

### Base URL

All API routes are prefixed with `/api/v1/` except for the health check endpoint.

### Available Endpoints

#### Health Check

```text
GET /health
```

Returns the current health status of the service.

#### User Management

```text
POST   /api/v1/users           # Create a new user
GET    /api/v1/users           # List users (with pagination)
GET    /api/v1/users/:id       # Get a specific user
PUT    /api/v1/users/:id       # Update a user
DELETE /api/v1/users/:id       # Delete a user
```

For detailed API documentation, visit the Swagger UI at `/swagger/index.html` when the server is running.

## Error Handling

The API uses standard HTTP status codes and returns errors in the following format:

```json
{
  "error": "Error type",
  "message": "Detailed error message"
}
```

Common status codes:

- 200: Success
- 201: Created
- 204: No Content
- 400: Bad Request
- 401: Unauthorized
- 403: Forbidden
- 404: Not Found
- 500: Internal Server Error

## Development

### Generate Swagger Documentation

```bash
# Generate/update Swagger docs
swag init
```

Access the Swagger UI at: `http://localhost:8080/swagger/index.html`

### Running the Application

1. Initialize the database (first time)

   ```bash
   ./bin/db-cli -refresh
   ```

2. Start the server

   ```bash
   # Direct start
   go run main.go

   # Or using air for hot reload (if installed)
   air
   ```

   The server will start on `http://localhost:8080`

### Database Migrations

1. Create a new model in the `models` directory
2. Add the model to migrations in `database/migration/migration.go`
3. Run migrations using the CLI tool

Example:

```go
// models/product.go
type Product struct {
    ID          uint           `gorm:"primarykey" json:"id"`
    Name        string         `gorm:"size:255;not null" json:"name"`
    Description string         `gorm:"type:text" json:"description"`
    Price       float64        `gorm:"not null" json:"price"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// Update migration.go to include the new model
func AutoMigrate() {
    err := config.DB.AutoMigrate(
        &models.User{},
        &models.Product{},
    )
    // ...
}
```

### Adding New Controllers

1. Create a new file in the `controllers` directory
2. Define your controller struct and methods
3. Register routes in `routes/routes.go`

Example:

```go
package controllers

type UserController struct{}

func NewUserController() *UserController {
    return &UserController{}
}

func (uc *UserController) HandleRequest(c *gin.Context) {
    // Implementation
}
```

## Database Migrations

When making changes to the database schema:

1. Update the relevant model in `models/`
2. Run migrations:

   ```bash
   ./bin/db-cli -migrate
   ```

To rollback changes:

```bash
./bin/db-cli -rollback
```

## Testing

To run tests:

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### Commit Message Format

Follow the [Conventional Commits](https://www.conventionalcommits.org/) specification:

```sh
<type>(<scope>): <description>

[optional body]

[optional footer(s)]
```

Types:

- feat: New feature
- fix: Bug fix
- docs: Documentation changes
- style: Code style changes (formatting, etc)
- refactor: Code refactoring
- test: Adding or updating tests
- chore: Maintenance tasks

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

Canh Nguyen - <canhcvp1998@gmail.com>

Project Link: [https://github.com/canhbk/golang-gin-starter-kit](https://github.com/canhbk/golang-gin-starter-kit)

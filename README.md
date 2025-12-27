# Kuchen

Kuchen is a small REST API project built while learning Go step by step.  
It focuses on understanding backend fundamentals by building something real, simple, and easy to reason about.

Rather than aiming to be a production-ready system, this project documents the process of learning Go, database integration, and API design through practice.

---

## Purpose of the Project

This project exists to support learning by doing.

It was created to:
- Learn Go as a backend language gradually
- Understand how ORMs work with relational databases
- Practice REST API design and HTTP handling
- Learn how to structure a Go project cleanly
- Gain confidence working with MySQL, GORM, and environment configuration

The code favors clarity and readability over cleverness.

---

## What the Application Does

Kuchen provides a REST API for managing cakes.

Using the API, you can:
- Create cake records
- View all cakes
- View a single cake by ID
- Update existing cakes
- Delete cakes (using soft deletes)

Each cake has a size and a price.  
Timestamps are handled automatically by the ORM.

---

## Features

- RESTful API built with Go
- MySQL database integration
- GORM ORM for database operations
- Automatic database migrations
- Soft delete support
- JSON request and response handling
- Environment-based configuration using `.env`
- Clear separation of concerns in project structure

---

## Technologies Used

- Go (backend language)
- GORM (ORM)
- MySQL (database)
- Gorilla Mux (HTTP routing)
- godotenv (environment variable loading)

No advanced Go knowledge is assumed.

---

## Prerequisites

To run this project, you need:
- Go 1.19 or later
- MySQL 5.7 or later
- Git

---

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/Joy-Phathomlezz/Kuchen.git
   cd Kuchen
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Set up environment variables by creating a `.env` file:
   ```
   DB_USER=your_mysql_username
   DB_PASS=your_mysql_password
   DB_HOST=127.0.0.1
   DB_PORT=3306
   DB_NAME=kuchen
   ```

4. Create the MySQL database:
   ```
   CREATE DATABASE kuchen;
   ```

## Usage

Run the application:
```
go run cmd/main/main.go
```

The server will start on port 9010.

## API Endpoints

### Create a Cake
```
POST /cake/
Content-Type: application/json

{
  "size": "Large",
  "price": 2.99
}
```

### Get All Cakes
```
GET /cake/
```

### Get a Cake by ID
```
GET /cake/{id}
```

### Update a Cake
```
PUT /cake/{id}
Content-Type: application/json

{
  "size": "Medium",
  "price": 1.99
}
```

### Delete a Cake
```
DELETE /cake/{id}
```

## Database Schema

The Cake model includes:
- ID (auto-incrementing primary key)
- Size (string)
- Price (float64)
- CreatedAt (timestamp)
- UpdatedAt (timestamp)
- DeletedAt (timestamp for soft deletes)

GORM automatically manages the timestamp fields and provides soft delete functionality.

## Architecture

- **Database Layer**: `internals/database/connection.go` - Establishes MySQL connection using GORM
- **Model Layer**: `internals/models/cake-model.go` - Defines the Cake struct with GORM tags
- **Repository Layer**: `internals/models/repository/cake-repository.go` - Contains database operations using GORM
- **Handler Layer**: `internals/handlers/cake-handler.go` - HTTP request handlers
- **Routes**: `internals/routes/cake-route.go` - API route definitions

## GORM and MySQL Integration

This project demonstrates GORM's capabilities with MySQL:

- Connection establishment with DSN configuration
- Automatic table creation via AutoMigrate
- CRUD operations using GORM methods
- Query building and execution
- Soft delete support
- Timestamp management

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

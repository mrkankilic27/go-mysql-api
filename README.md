# Go REST API with MySQL

This project is a simple RESTful API built using Go (Golang) with MySQL integration. It demonstrates how to build a backend service using Go’s `net/http` package and connect to a MySQL database for basic CRUD operations.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
- [Project Structure](#project-structure)
- [Endpoints](#endpoints)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Overview

This Go-based API connects to a MySQL database and provides REST endpoints to interact with data. It serves as a foundational backend service for applications that require a simple, fast, and scalable API using Go.

## Features

- RESTful architecture
- MySQL database connection
- Basic CRUD (Create, Read, Update, Delete) operations
- Lightweight and fast
- Easily extendable

## Technologies Used

- **Go** – Programming language used for building the API
- **MySQL** – Relational database for storing application data
- **net/http** – Go’s standard HTTP package for serving endpoints
- **database/sql** – Go's SQL database driver interface
- **github.com/go-sql-driver/mysql** – MySQL driver for Go

## Getting Started

### Prerequisites

- Go 1.18 or higher installed
- MySQL server running locally or remotely

### Installation

1. Clone the repository:

```bash
git clone https://github.com/mrkankilic27/go-mysql-api.git
cd go-mysql-api
```

2. Initialize Go modules:

```bash
go mod tidy
```

3. Set up your MySQL database and update connection credentials in `main.go`.

4. Run the application:

```bash
go run main.go
```

## Project Structure

```
go-mysql-api/
├── main.go          # Entry point with API routes and DB logic
├── go.mod           # Go module definition
└── go.sum           # Go dependencies hash
```

## Endpoints

- `GET /users` – List all users
- `GET /users/{id}` – Get user by ID
- `POST /users` – Create new user
- `PUT /users/{id}` – Update user by ID
- `DELETE /users/{id}` – Delete user by ID

*(Note: Endpoint routes may vary based on your implementation.)*

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change or improve.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

**Mertcan Kankılıç**  
Email: mertcankankilic27@gmail.com  
GitHub: [mrkankilic27](https://github.com/mrkankilic27)
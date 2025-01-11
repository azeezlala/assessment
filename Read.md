# Cloud Resource API

Cloud Resource Inventory Management System [Go](https://golang.org)

## Prerequisites

Before diving in, make sure you have:

- A Unix-like OS (Linux/macOS)
- Docker & Docker Compose
- Go 1.23+

## Quick Start

Ready to launch? Follow these steps:

- Sync your local database with the schema and run migration:
   
    ```bash
      go run cmd/migrate/main.go migrate
    ```

- Run the seeder
  
  ```bash
    go run cmd/migrate/main.go seed
  ```

- Start the server
  ```bash
    docker-compose up -d
  ```
  
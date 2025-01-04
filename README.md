# sqlc-api
sqlc-api is a Go-based API server that utilizes SQLC for generating type-safe database interactions from SQL queries.

## Features
- SQLC Integration: Automatically generates Go code from SQL queries, ensuring type safety and reducing boilerplate.
- Database Migrations: Manages schema changes using SQL migration files.
- HTTP API Endpoints: Provides a set of HTTP endpoints for interacting with the application.

## Prerequisites
Before setting up the project, ensure you have the following installed:

- Go: Version 1.16 or higher.
- MySQL: Ensure the database server is running and accessible.
- SQLC: Follow the installation guide to set up SQLC.

## Getting Started
1. Clone the Repository
```bash
git clone https://github.com/harliandi/sqlc-api.git
cd sqlc-api
```
2. Set Up Environment Variables
Create a .env file in the root directory with the following content:
```env
SERVER_PORT=3000
DB_USER=root
DB_PASSWORD=secret
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=sqlc_api
```
Replace with your MySQL credentials and database details.

3. Apply Database Migrations
Ensure your database is set up with the necessary schema by applying the migration files located in the migrations directory. Use a tool like golang-migrate to apply these migrations:
```bash
migrate -path migrations -database $DATABASE_URL up
```
4. Generate SQLC Code
Run SQLC to generate the Go code for database interactions:
```bash
sqlc generate
```
5. Install Dependencies
Ensure all Go dependencies are installed:
```bash
go mod tidy
```
6. Run the Server
Start the API server:
```bash
go run main.go
```
The server will start on http://localhost:3000.

## API Endpoints
The API provides the following endpoints:

- GET /author: Retrieves a list of authors.
- POST /author: Creates a new author.
- GET /author/{id}: Retrieves a specific author by ID.
- PUT /author/{id}: Updates an existing author by ID.
- DELETE /author/{id}: Deletes an author by ID.

For detailed request and response formats, refer to the API documentation.

## Acknowledgements
- SQLC for generating type-safe Go code from SQL queries.
- Chi for the HTTP router.
- Golang Migrate for database migrations.

For more information, visit the repository.
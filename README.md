# hyper-api

A Go API server for managing users, dashboard data, and map information.

# Full Project:
Golang API: https://github.com/HarisKordic/hyper-api
Next.js Frontend: https://github.com/HarisKordic/hyper6xhurmasice

## Development

1. Clone the repository
2. Create `.env` file with database configuration:
```sh
DATABASE_URL="postgres://username:password@localhost:5432/dbname"
```
3. Run the development server:
```
go run main.go
```

## Production Deployment

1. Build the production binary
```
go build -o prod_server
```
2. Start the server
```
./prod_server
```

## Routes:

### Get dashboard data
curl http://localhost:8080/api/dashboard

### Get map users
curl http://localhost:8080/api/map/users

### Get latest Google users
curl http://localhost:8080/api/users

### Get user by email 
curl http://localhost:8080/api/users/email/example@email.com

## E2E Tests:

1. Set up a test environment with database connection and test server  
2. Tests all major API endpoints:  
    - `/api/dashboard` - Checks carbon footprint and pollution data  
    - `/api/map/users` - Validates map user data  
    - `/api/users` - Verifies user listing with limit  
    - `/api/users/email[/email]` - Tests user lookup by email  
3. Uses `testify/assert` for cleaner assertions  
4. Includes data validation for each endpoint  
5. Handles error cases (e.g., non-existent email)  

## Running the Tests  
To run the tests:  
```bash  
go test -v ./tests/...
```

## Prerequisites
Make sure to have the env file with valid database credentials and have test data in the database

## Unit testing
Run the unit tests by executing: go test -v ./…
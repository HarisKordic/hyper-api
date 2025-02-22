# hyper-api

A Go API server for managing users, dashboard data, and map information.

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
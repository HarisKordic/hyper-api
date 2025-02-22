#!/bin/bash

# Pull latest changes
echo "Pulling latest changes..."
git pull

echo "Killing Go server if already running on port 8080..."
fuser -k 8080/tcp || true

# Build the project
echo "Building Go project..."
go build -o prod_server

# Start the server in background
echo "Starting Go server..."
nohup ./prod_server > /dev/null 2>&1 &

echo "Deployment complete! Go server is running on port 8080"
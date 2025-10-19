#!/bin/sh

cd /app

echo "Downloading dependencies..."
go mod download

echo "Building application..."
go build -o /app/main /app/main.go
go build -o /app/migrator /app/cli/migrator/main.go

echo "Starting application..."
exec /app/main
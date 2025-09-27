#!/bin/sh

echo "Downloading dependencies..."
go mod download

echo "Building application..."
go build -o /app/main /app/main.go

echo "Starting application..."
exec /app/main
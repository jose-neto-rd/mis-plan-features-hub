#!/bin/sh

echo "Running go mod tidy..."
go mod tidy

echo "Starting the service..."
exec go run cmd/main.go
#!/bin/bash

echo "************************"
echo "Starting testing input processing"
echo "************************"

echo "Testing 100KB input no error"
yes testing noerr | head -c 100KB | go run main.go

echo ""

echo "Testing 100KB input error"
yes testing error | head -c 100KB | go run main.go

echo ""

echo "Testing 10GB input error"
yes testing error | head -c 10GB | go run main.go

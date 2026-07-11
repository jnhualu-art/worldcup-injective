#!/usr/bin/env bash
# Run both backend and x402 gateway locally (macOS/Linux/WSL).
set -e

ROOT="$(cd "$(dirname "$0")" && pwd)"

echo "Starting Go backend..."
cd "$ROOT/backend"
go run main.go &
BACKEND_PID=$!

sleep 3

echo "Starting x402 gateway..."
cd "$ROOT/x402-gateway"
npm install
npm start &
GATEWAY_PID=$!

echo ""
echo "Backend: http://localhost:8080"
echo "Gateway: http://localhost:3000"
echo "Frontend: open $ROOT/frontend/index.html"
echo ""
echo "Press Ctrl+C to stop both services."
wait $BACKEND_PID $GATEWAY_PID

@echo off
REM Run both backend and x402 gateway locally on Windows.

echo Starting Go backend...
start "Backend" cmd /k "cd /d D:\WorkBuddy\InjectiveGlobalCup\backend && go run main.go"

timeout /t 3 >nul

echo Starting x402 gateway...
start "x402 Gateway" cmd /d "cd /d D:\WorkBuddy\InjectiveGlobalCup\x402-gateway && npm install && npm start"

echo.
echo Backend: http://localhost:8080
echo Gateway: http://localhost:3000
echo Frontend: open D:\WorkBuddy\InjectiveGlobalCup\frontend\index.html
pause

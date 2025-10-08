# Run AI Interviewer Locally (No Docker)
Write-Host "================================" -ForegroundColor Green
Write-Host "  Starting AI Interviewer      " -ForegroundColor Green
Write-Host "================================" -ForegroundColor Green
Write-Host ""

# Load environment variables
if (Test-Path ".env") {
    Get-Content ".env" | ForEach-Object {
        if ($_ -match '^([^=]+)=(.*)$') {
            [System.Environment]::SetEnvironmentVariable($matches[1], $matches[2], "Process")
        }
    }
}

# Update environment for local MySQL
$env:DB_HOST = "localhost"
$env:DB_PORT = "3306"

Write-Host "Starting Backend..." -ForegroundColor Yellow
Start-Process powershell -ArgumentList "-NoExit", "-Command", "cd '$PWD\backend'; go run cmd/server/main.go"

Start-Sleep -Seconds 2

Write-Host "Starting Frontend..." -ForegroundColor Yellow  
Start-Process powershell -ArgumentList "-NoExit", "-Command", "cd '$PWD\frontend'; npm run dev"

Write-Host ""
Write-Host "================================" -ForegroundColor Green
Write-Host "  Services Started!             " -ForegroundColor Green
Write-Host "================================" -ForegroundColor Green
Write-Host ""
Write-Host "Frontend: http://localhost:5173" -ForegroundColor Cyan
Write-Host "Backend:  http://localhost:8080" -ForegroundColor Cyan
Write-Host ""
Write-Host "Press Ctrl+C in each window to stop" -ForegroundColor Yellow

# PowerShell script to set up and run AI Interviewer
# Usage: .\start.ps1

Write-Host "==================================" -ForegroundColor Green
Write-Host "  AI Interviewer Setup & Start   " -ForegroundColor Green
Write-Host "==================================" -ForegroundColor Green
Write-Host ""

# Check if Docker is running
Write-Host "Checking Docker..." -ForegroundColor Yellow
$dockerRunning = docker info 2>&1 | Out-Null
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error: Docker is not running!" -ForegroundColor Red
    Write-Host "Please start Docker Desktop and try again." -ForegroundColor Red
    exit 1
}
Write-Host "Docker is running ✓" -ForegroundColor Green
Write-Host ""

# Check if .env file exists
if (-Not (Test-Path ".env")) {
    Write-Host "Creating .env file from template..." -ForegroundColor Yellow
    Copy-Item ".env.example" ".env"
    Write-Host "Please edit .env file and add your GEMINI_API_KEY" -ForegroundColor Red
    Write-Host "Then run this script again." -ForegroundColor Red
    notepad .env
    exit 0
}

# Check if GEMINI_API_KEY is set
$envContent = Get-Content ".env" -Raw
if ($envContent -match "GEMINI_API_KEY=your_gemini_api_key_here" -or $envContent -match "GEMINI_API_KEY=\s*$") {
    Write-Host "Warning: GEMINI_API_KEY not set in .env file!" -ForegroundColor Red
    Write-Host "Please add your Gemini API key to the .env file" -ForegroundColor Red
    $response = Read-Host "Do you want to edit .env now? (y or n)"
    if ($response -eq "y") {
        notepad .env
        exit 0
    }
    exit 1
}
Write-Host "Environment configured ✓" -ForegroundColor Green
Write-Host ""

# Stop any existing containers
Write-Host "Stopping existing containers..." -ForegroundColor Yellow
docker-compose down 2>&1 | Out-Null
Write-Host "Stopped ✓" -ForegroundColor Green
Write-Host ""

# Start services
Write-Host "Starting services..." -ForegroundColor Yellow
Write-Host "This may take a few minutes on first run..." -ForegroundColor Cyan
Write-Host ""

docker-compose up --build -d

if ($LASTEXITCODE -ne 0) {
    Write-Host "Error: Failed to start services!" -ForegroundColor Red
    exit 1
}

Write-Host ""
Write-Host "Services started successfully! ✓" -ForegroundColor Green
Write-Host ""
Write-Host "==================================" -ForegroundColor Green
Write-Host "  Application is now running!    " -ForegroundColor Green
Write-Host "==================================" -ForegroundColor Green
Write-Host ""
Write-Host "Frontend: http://localhost:3000" -ForegroundColor Cyan
Write-Host "Backend:  http://localhost:8080" -ForegroundColor Cyan
Write-Host "Database: localhost:3306" -ForegroundColor Cyan
Write-Host ""
Write-Host "To view logs: docker-compose logs -f" -ForegroundColor Yellow
Write-Host "To stop:      docker-compose down" -ForegroundColor Yellow
Write-Host ""
Write-Host "Opening browser..." -ForegroundColor Yellow
Start-Sleep -Seconds 3
Start-Process "http://localhost:3000"

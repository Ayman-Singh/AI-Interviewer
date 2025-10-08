# Local Setup Script (No Docker)
Write-Host "================================" -ForegroundColor Green
Write-Host "  AI Interviewer Local Setup   " -ForegroundColor Green
Write-Host "================================" -ForegroundColor Green
Write-Host ""

# Check if MySQL is installed
Write-Host "Checking MySQL..." -ForegroundColor Yellow
$mysqlPath = Get-Command mysql -ErrorAction SilentlyContinue
if (-not $mysqlPath) {
    Write-Host "MySQL not found. Please install MySQL first." -ForegroundColor Red
    Write-Host "Download from: https://dev.mysql.com/downloads/installer/" -ForegroundColor Yellow
    exit 1
}
Write-Host "MySQL found ✓" -ForegroundColor Green

# Check if .env file exists
if (-Not (Test-Path ".env")) {
    Write-Host "Creating .env file..." -ForegroundColor Yellow
    Copy-Item ".env.example" ".env"
    Write-Host ".env file created ✓" -ForegroundColor Green
}

Write-Host ""
Write-Host "Setup complete!" -ForegroundColor Green
Write-Host ""
Write-Host "Next steps:" -ForegroundColor Cyan
Write-Host "1. Make sure MySQL service is running" -ForegroundColor White
Write-Host "2. Run: .\run-local.ps1" -ForegroundColor White

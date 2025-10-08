# PowerShell script to stop AI Interviewer services
# Usage: .\stop.ps1

Write-Host "==================================" -ForegroundColor Red
Write-Host "  Stopping AI Interviewer        " -ForegroundColor Red
Write-Host "==================================" -ForegroundColor Red
Write-Host ""

Write-Host "Stopping all services..." -ForegroundColor Yellow
docker-compose down

if ($LASTEXITCODE -eq 0) {
    Write-Host ""
    Write-Host "All services stopped successfully! ✓" -ForegroundColor Green
} else {
    Write-Host ""
    Write-Host "Error stopping services" -ForegroundColor Red
    exit 1
}

$response = Read-Host "`nDo you want to remove all data (database volumes)? (y/n)"
if ($response -eq "y") {
    Write-Host "Removing volumes..." -ForegroundColor Yellow
    docker-compose down -v
    Write-Host "All data removed ✓" -ForegroundColor Green
}

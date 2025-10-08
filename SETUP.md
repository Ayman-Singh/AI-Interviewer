# Quick Setup Guide

## Prerequisites Check

Before starting, ensure you have:
- [ ] Docker Desktop installed and running
- [ ] Docker Compose installed (comes with Docker Desktop)
- [ ] Google Gemini API key

## Step-by-Step Setup

### 1. Get Your Gemini API Key

1. Visit https://makersuite.google.com/app/apikey
2. Sign in with your Google account
3. Click "Create API Key"
4. Copy the generated API key

### 2. Configure Environment

Open the `.env` file in the root directory and paste your API key:

```env
GEMINI_API_KEY=paste_your_key_here
```

### 3. Start the Application

Open PowerShell in the project directory and run:

```powershell
docker-compose up --build
```

Wait for all services to start (this may take a few minutes on first run).

### 4. Verify Services

Check that all services are running:

- Backend: http://localhost:8080/api/health
- Frontend: http://localhost:3000
- Database: localhost:3306

### 5. Use the Application

1. Open http://localhost:3000 in your browser
2. Fill in your details:
   - Name
   - Email (for tracking history)
   - Position (e.g., "Software Engineer")
   - Difficulty level
3. Click "Start Interview"
4. Answer questions and receive instant feedback
5. View your results and history

## Troubleshooting

### Port Already in Use

If you get a "port already in use" error:

```powershell
# Stop all containers
docker-compose down

# Check what's using the ports
netstat -ano | findstr :3000
netstat -ano | findstr :8080
netstat -ano | findstr :3306

# Kill the process using the port (replace PID with actual process ID)
taskkill /PID <PID> /F
```

### Database Connection Failed

If the backend can't connect to the database:

```powershell
# Check if MySQL container is running
docker ps

# View MySQL logs
docker-compose logs mysql

# Restart services
docker-compose restart
```

### Frontend Build Errors

```powershell
# Rebuild frontend only
docker-compose up --build frontend
```

### Backend Build Errors

```powershell
# Rebuild backend only
docker-compose up --build backend
```

## Development Mode

For development with hot-reload:

### Backend
```powershell
cd backend
go mod download
go run cmd/server/main.go
```

### Frontend
```powershell
cd frontend
npm install
npm run dev
```

## Stopping the Application

```powershell
# Stop all services
docker-compose down

# Stop and remove volumes (clears database)
docker-compose down -v
```

## Updating the Application

```powershell
# Pull latest changes
git pull

# Rebuild and restart
docker-compose down
docker-compose up --build
```

## Common Issues

### Issue: "GEMINI_API_KEY is required"
**Solution**: Make sure you've set the API key in the `.env` file

### Issue: Frontend shows "Failed to connect to backend"
**Solution**: Wait for backend to fully start (check logs with `docker-compose logs backend`)

### Issue: Database tables not created
**Solution**: Delete the volume and restart:
```powershell
docker-compose down -v
docker-compose up --build
```

## Next Steps

- Start your first interview
- Try different difficulty levels
- View your interview history
- Check the detailed results and feedback

For more information, see the main README.md file.

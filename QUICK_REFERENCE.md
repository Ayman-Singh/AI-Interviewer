# 🚀 Quick Reference Guide

## Essential Commands

### Starting the Application
```powershell
# Using the script (Recommended)
.\start.ps1

# Or manually
docker-compose up --build -d
```

### Stopping the Application
```powershell
# Using the script
.\stop.ps1

# Or manually
docker-compose down
```

### Viewing Logs
```powershell
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f backend
docker-compose logs -f frontend
docker-compose logs -f mysql
```

### Restarting Services
```powershell
# Restart all
docker-compose restart

# Restart specific service
docker-compose restart backend
```

## URLs

| Service  | URL                              |
|----------|----------------------------------|
| Frontend | http://localhost:3000            |
| Backend  | http://localhost:8080            |
| API      | http://localhost:8080/api        |
| Health   | http://localhost:8080/api/health |

## File Locations

### Configuration Files
- Environment: `.env`
- Docker: `docker-compose.yml`
- Backend Config: `backend/internal/config/config.go`
- Frontend Config: `frontend/vite.config.js`

### Key Backend Files
- Main Entry: `backend/cmd/server/main.go`
- AI Service: `backend/internal/ai/service.go`
- Handlers: `backend/internal/handlers/handlers.go`
- Models: `backend/internal/models/models.go`
- Database: `backend/db/init.sql`

### Key Frontend Files
- App Entry: `frontend/src/main.jsx`
- Home Page: `frontend/src/pages/Home.jsx`
- Interview: `frontend/src/pages/Interview.jsx`
- Results: `frontend/src/pages/Results.jsx`
- History: `frontend/src/pages/History.jsx`
- API Client: `frontend/src/services/api.js`

## Environment Variables

### Required
```env
GEMINI_API_KEY=your_api_key_here
```

### Optional
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=interviewer
DB_PASSWORD=interviewerpass
DB_NAME=ai_interviewer
PORT=8080
```

## Common Tasks

### Reset Database
```powershell
docker-compose down -v
docker-compose up --build
```

### View Database
```powershell
docker exec -it ai_interviewer_db mysql -u interviewer -pinterviewerpass ai_interviewer
```

### Backend Development Mode
```powershell
cd backend
go run cmd/server/main.go
```

### Frontend Development Mode
```powershell
cd frontend
npm run dev
```

### Build Frontend
```powershell
cd frontend
npm run build
```

### Run Tests
```powershell
# Backend
cd backend
go test ./...

# Frontend
cd frontend
npm test
```

## Troubleshooting

### Port Already in Use
```powershell
# Find process using port
netstat -ano | findstr :3000
netstat -ano | findstr :8080

# Kill process
taskkill /PID <process_id> /F
```

### Docker Issues
```powershell
# Clean everything
docker-compose down -v
docker system prune -a

# Restart Docker Desktop
```

### Database Connection Failed
```powershell
# Check MySQL logs
docker-compose logs mysql

# Verify MySQL is running
docker ps | findstr mysql
```

### Backend Won't Start
```powershell
# Check if API key is set
cat .env | findstr GEMINI_API_KEY

# View backend logs
docker-compose logs backend
```

### Frontend Build Errors
```powershell
# Clean and reinstall
cd frontend
Remove-Item -Recurse -Force node_modules
npm install
```

## API Quick Reference

### Start Interview
```bash
POST /api/interview/start
{
  "user_name": "John Doe",
  "email": "john@example.com",
  "position": "Software Engineer",
  "difficulty": "medium"
}
```

### Submit Answer
```bash
POST /api/interview/submit
{
  "question_id": 1,
  "response_text": "Your answer..."
}
```

### Get Interview
```bash
GET /api/interview/1
```

### Get History
```bash
GET /api/interviews?email=john@example.com
```

## Keyboard Shortcuts (Frontend)

- `Ctrl+Enter` - Submit answer (in textarea)
- `Ctrl+R` - Refresh page
- `Ctrl+P` - Print results

## Database Queries

### View All Interviews
```sql
SELECT * FROM interviews ORDER BY started_at DESC;
```

### View User's Interviews
```sql
SELECT * FROM interviews WHERE user_id = 1;
```

### View Questions for Interview
```sql
SELECT * FROM questions WHERE interview_id = 1 ORDER BY order_num;
```

### View Average Scores
```sql
SELECT AVG(score) as avg_score FROM interviews WHERE status = 'completed';
```

## Docker Commands

### View Running Containers
```powershell
docker ps
```

### Access Container Shell
```powershell
# Backend
docker exec -it ai_interviewer_backend sh

# Database
docker exec -it ai_interviewer_db mysql -u root -p
```

### View Container Stats
```powershell
docker stats
```

### Remove All Containers
```powershell
docker-compose down -v
docker system prune -a
```

## Git Commands

### Initial Commit
```bash
git init
git add .
git commit -m "Initial commit: AI Interviewer"
git branch -M main
git remote add origin <your-repo-url>
git push -u origin main
```

### Update
```bash
git add .
git commit -m "feat: add new feature"
git push
```

## Performance Tips

1. **Backend**: Increase Go routines for concurrent request handling
2. **Frontend**: Enable React production mode for better performance
3. **Database**: Add indexes on frequently queried columns
4. **Docker**: Allocate more memory to Docker Desktop
5. **Caching**: Implement Redis for API response caching

## Security Checklist

- [ ] Set strong database passwords
- [ ] Keep Gemini API key secure (never commit)
- [ ] Use environment variables for secrets
- [ ] Enable HTTPS in production
- [ ] Implement rate limiting
- [ ] Add user authentication
- [ ] Validate all inputs
- [ ] Use prepared statements for SQL

## Deployment Checklist

- [ ] Set production environment variables
- [ ] Configure CORS for production domains
- [ ] Enable HTTPS/SSL
- [ ] Set up database backups
- [ ] Configure logging and monitoring
- [ ] Set up error tracking
- [ ] Optimize Docker images
- [ ] Test all endpoints
- [ ] Set up CI/CD pipeline

## Getting Help

- **Documentation**: See README.md
- **Setup Guide**: See SETUP.md
- **API Docs**: See API.md
- **Contributing**: See CONTRIBUTING.md
- **Project Overview**: See PROJECT_OVERVIEW.md

## Version Info

- **Go**: 1.21+
- **Node**: 18+
- **Docker**: 20.10+
- **MySQL**: 8.0+
- **React**: 18.2+

---

**Quick Start**: `.\start.ps1` → Open `http://localhost:3000` → Start interviewing! 🚀

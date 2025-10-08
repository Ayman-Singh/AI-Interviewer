# AI Interviewer - Project Overview

## 🎯 Project Summary

A full-stack AI-powered interview preparation platform that helps candidates practice and improve their interview skills with real-time feedback.

## 📊 Technology Stack

```
┌─────────────────────────────────────────────────────────┐
│                     FRONTEND (React)                     │
│  - React 18 + Vite                                      │
│  - React Router for navigation                          │
│  - Axios for API calls                                  │
│  - Black & Green theme                                  │
│  - Port: 3000 (dev) / 80 (prod)                        │
└─────────────────────────────────────────────────────────┘
                            ↓ HTTP/REST
┌─────────────────────────────────────────────────────────┐
│                   BACKEND (Go + Gemini)                  │
│  - Go 1.21                                              │
│  - Gorilla Mux (routing)                                │
│  - LangChainGo + Gemini AI                              │
│  - Clean Architecture                                   │
│  - Port: 8080                                           │
└─────────────────────────────────────────────────────────┘
                            ↓ SQL
┌─────────────────────────────────────────────────────────┐
│                   DATABASE (MySQL 8.0)                   │
│  - Users, Interviews, Questions, Responses              │
│  - Persistent volumes                                   │
│  - Port: 3306                                           │
└─────────────────────────────────────────────────────────┘
```

## 🗂️ Project Structure

```
Interviewer/
│
├── 📁 backend/                    # Go backend application
│   ├── cmd/
│   │   └── server/
│   │       └── main.go           # Application entry point
│   ├── internal/
│   │   ├── ai/                   # AI service (Gemini integration)
│   │   ├── config/               # Configuration management
│   │   ├── database/             # Database connection
│   │   ├── handlers/             # HTTP request handlers
│   │   ├── models/               # Data models & DTOs
│   │   └── repository/           # Database operations
│   ├── db/
│   │   └── init.sql              # Database schema
│   ├── Dockerfile                # Backend container
│   ├── go.mod                    # Go dependencies
│   └── go.sum                    # Dependency checksums
│
├── 📁 frontend/                   # React frontend application
│   ├── src/
│   │   ├── pages/
│   │   │   ├── Home.jsx          # Landing page
│   │   │   ├── Interview.jsx    # Interview session
│   │   │   ├── Results.jsx      # Results display
│   │   │   └── History.jsx      # Interview history
│   │   ├── services/
│   │   │   └── api.js            # API client
│   │   ├── App.jsx               # Main app component
│   │   ├── App.css               # Global styles
│   │   ├── main.jsx              # React entry point
│   │   └── index.css             # Theme & base styles
│   ├── Dockerfile                # Frontend container
│   ├── nginx.conf                # Nginx configuration
│   ├── package.json              # NPM dependencies
│   └── vite.config.js            # Vite configuration
│
├── 📄 docker-compose.yml          # Multi-container orchestration
├── 📄 .env.example                # Environment template
├── 📄 .gitignore                  # Git ignore rules
├── 📄 README.md                   # Main documentation
├── 📄 SETUP.md                    # Setup guide
├── 📄 API.md                      # API documentation
├── 📄 CONTRIBUTING.md             # Contribution guidelines
├── 📄 Makefile                    # Build automation
├── 📄 start.ps1                   # Windows start script
└── 📄 stop.ps1                    # Windows stop script
```

## 🔄 Application Flow

```
1. User Starts Interview
   ↓
   [Home Page] → Submit form with name, email, position, difficulty
   ↓
   [POST /api/interview/start]
   ↓
   Backend creates user (if new) and interview record
   ↓
   Gemini AI generates 5 questions based on position & difficulty
   ↓
   Questions stored in database
   ↓
   Returns first question to frontend
   ↓
   [Interview Page] → Displays question

2. User Answers Question
   ↓
   [Interview Page] → User types answer and submits
   ↓
   [POST /api/interview/submit]
   ↓
   Gemini AI evaluates answer
   ↓
   Feedback & score calculated (0-10)
   ↓
   Response stored in database
   ↓
   Returns feedback, score, and next question (or completion)
   ↓
   [Interview Page] → Shows feedback, then next question

3. Interview Completion
   ↓
   All questions answered
   ↓
   Average score calculated
   ↓
   Interview marked as completed
   ↓
   [Results Page] → Display full results with breakdown

4. View History
   ↓
   [History Page] → Enter email
   ↓
   [GET /api/interviews?email=...]
   ↓
   Returns all interviews for user
   ↓
   Display list with scores and status
```

## 🎨 UI Flow

```
┌─────────────────────┐
│     Home Page       │
│  - Enter details    │
│  - Select difficulty│
│  - View history btn │
└──────────┬──────────┘
           │
           ├─────────────────────┐
           │                     │
           ↓                     ↓
┌─────────────────────┐  ┌─────────────────────┐
│  Interview Page     │  │   History Page      │
│  - Show question    │  │  - List interviews  │
│  - Answer input     │  │  - View results     │
│  - Submit answer    │  │  - Continue active  │
│  - See feedback     │  └─────────────────────┘
└──────────┬──────────┘
           │
           ↓
┌─────────────────────┐
│   Results Page      │
│  - Overall score    │
│  - Q&A breakdown    │
│  - Feedback details │
│  - Start new button │
└─────────────────────┘
```

## 🗄️ Database Schema

```sql
users
├── id (PK)
├── name
├── email (UNIQUE)
└── created_at

interviews
├── id (PK)
├── user_id (FK → users.id)
├── position
├── difficulty (easy/medium/hard)
├── status (in_progress/completed)
├── score (0-10, nullable)
├── started_at
└── completed_at (nullable)

questions
├── id (PK)
├── interview_id (FK → interviews.id)
├── question_text
├── question_type (technical/behavioral/coding)
├── order_num
└── created_at

responses
├── id (PK)
├── question_id (FK → questions.id)
├── response_text
├── feedback
├── score (0-10, nullable)
└── created_at
```

## 🔌 API Endpoints

| Method | Endpoint              | Description                    |
|--------|-----------------------|--------------------------------|
| GET    | /api/health          | Health check                   |
| POST   | /api/interview/start | Create new interview           |
| POST   | /api/interview/submit| Submit answer to question      |
| GET    | /api/interview/{id}  | Get interview details          |
| GET    | /api/interviews      | Get user's interview history   |

## 🚀 Quick Start Commands

```powershell
# Start everything
docker-compose up --build

# Or use the script
.\start.ps1

# View logs
docker-compose logs -f

# Stop everything
docker-compose down

# Or use the script
.\stop.ps1
```

## 🎯 Key Features

✅ **AI-Powered Questions** - Gemini generates contextual interview questions
✅ **Real-Time Evaluation** - Instant feedback on every answer
✅ **Smart Scoring** - 0-10 scale with detailed feedback
✅ **Interview History** - Track progress over time
✅ **Multiple Difficulty Levels** - Easy, Medium, Hard
✅ **Various Question Types** - Technical, Behavioral, Coding
✅ **Beautiful UI** - Black & green hacker aesthetic
✅ **Docker Ready** - One command to start everything
✅ **Responsive Design** - Works on all devices

## 📈 Scoring System

| Score Range | Level              | Description                    |
|-------------|--------------------|--------------------------------|
| 8.0 - 10.0  | Excellent          | Strong understanding           |
| 6.0 - 7.9   | Good               | Solid with minor improvements  |
| 4.0 - 5.9   | Fair               | Basic understanding            |
| 0.0 - 3.9   | Needs Improvement  | Significant gaps               |

## 🔐 Environment Variables

### Required
- `GEMINI_API_KEY` - Your Google Gemini API key

### Optional (with defaults)
- `DB_HOST` - Database host (default: localhost)
- `DB_PORT` - Database port (default: 3306)
- `DB_USER` - Database user (default: interviewer)
- `DB_PASSWORD` - Database password
- `DB_NAME` - Database name (default: ai_interviewer)
- `PORT` - Backend port (default: 8080)

## 🛠️ Development

### Backend Development
```powershell
cd backend
go mod download
go run cmd/server/main.go
```

### Frontend Development
```powershell
cd frontend
npm install
npm run dev
```

## 📦 Docker Services

- **mysql** - MySQL 8.0 database
- **backend** - Go API server
- **frontend** - React app with Nginx

All services are connected via `ai_interviewer_network`.

## 🎨 Theme Colors

```css
Black:        #000000
Dark Gray:    #0a0a0a
Darker Gray:  #121212
Neon Green:   #00ff41
Green Dark:   #00cc33
White:        #ffffff
Light Gray:   #cccccc
Muted Gray:   #888888
```

## 📝 Next Steps

1. Get Gemini API key
2. Configure .env file
3. Run `docker-compose up --build`
4. Visit http://localhost:3000
5. Start your first interview!

## 🤝 Contributing

See CONTRIBUTING.md for guidelines on how to contribute to this project.

## 📄 License

MIT License - See LICENSE file for details.

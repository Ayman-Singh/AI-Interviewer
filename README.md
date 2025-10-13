# AI Interviewer 

A full-stack AI-powered interview preparation platform that helps candidates practice and improve their interview skills with real-time feedback powered by Google Gemini AI.

##  Features

- **AI-Generated Questions**: Dynamic interview questions tailored to specific positions and difficulty levels
- **Real-Time Evaluation**: Instant feedback and scoring on answers using Gemini AI
- **Multiple Question Types**: Technical, behavioral, and coding questions
- **Interview History**: Track your progress across multiple interview sessions
- **Detailed Results**: Comprehensive breakdown of performance with question-by-question analysis

##  Architecture

### Backend
- **Language**: Go 1.21
- **Framework**: Gorilla Mux (HTTP routing)
- **AI Framework**: LangChainGo with Google Gemini AI
- **Database**: MySQL 8.0
- **Architecture**: Clean architecture with separation of concerns

### Frontend
- **Framework**: React 18 with Vite
- **Styling**: Custom CSS with black/green theme
- **Routing**: React Router v6
- **API Client**: Axios

### Infrastructure
- **Containerization**: Docker & Docker Compose
- **Database**: MySQL with persistent volumes
- **Web Server**: Nginx (for production frontend)

##  Prerequisites

- Docker and Docker Compose installed
- Google Gemini API Key ([Get one here](https://makersuite.google.com/app/apikey))
- Git


```bash
cp .env.example .env
```

Edit `.env` and add your Gemini API key:

```env
GEMINI_API_KEY=your_actual_api_key_here
```

```bash
docker-compose up --build
```

This will start:
- MySQL database on port 3306
- Go backend on port 8080
- React frontend on port 3000

Open your browser and navigate to:
```
http://localhost:3000
```

##  Development Setup

### Backend Development

```bash
cd backend

# Install dependencies
go mod download

# Set up environment variables
cp ../.env.example .env

# Run the server
go run cmd/server/main.go
```

The backend will be available at `http://localhost:8080`

### Frontend Development

```bash
cd frontend

# Install dependencies
npm install

# Copy environment file
cp .env.example .env

# Start development server
npm run dev
```

The frontend will be available at `http://localhost:5173`

## 🔌 API Endpoints

### Health Check
```
GET /api/health
```

### Start Interview
```
POST /api/interview/start
Content-Type: application/json

{
  "user_name": "John Doe",
  "email": "john@example.com",
  "position": "Software Engineer",
  "difficulty": "medium"
}
```

### Submit Answer
```
POST /api/interview/submit
Content-Type: application/json

{
  "question_id": 1,
  "response_text": "Your answer here..."
}
```

### Get Interview Details
```
GET /api/interview/{id}
```

### Get User Interview History
```
GET /api/interviews?email=user@example.com
```

##  Database Schema

### Users Table
- `id` - Primary key
- `name` - User's full name
- `email` - User's email (unique)
- `created_at` - Timestamp

### Interviews Table
- `id` - Primary key
- `user_id` - Foreign key to users
- `position` - Job position
- `difficulty` - easy/medium/hard
- `status` - in_progress/completed
- `score` - Overall score (0-10)
- `started_at` - Start timestamp
- `completed_at` - Completion timestamp

### Questions Table
- `id` - Primary key
- `interview_id` - Foreign key to interviews
- `question_text` - The question
- `question_type` - technical/behavioral/coding
- `order_num` - Question order
- `created_at` - Timestamp

### Responses Table
- `id` - Primary key
- `question_id` - Foreign key to questions
- `response_text` - User's answer
- `feedback` - AI-generated feedback
- `score` - Score for this answer (0-10)
- `created_at` - Timestamp


##  Testing

### Backend Tests
```bash
cd backend
go test ./...
```

### Frontend Tests
```bash
cd frontend
npm test
```

##  Docker Commands

### Build and start all services
```bash
docker-compose up --build
```

### Start services in background
```bash
docker-compose up -d
```

### Stop all services
```bash
docker-compose down
```

### View logs
```bash
docker-compose logs -f
```

### Rebuild specific service
```bash
docker-compose up --build backend
docker-compose up --build frontend
```

##  Performance Scoring

The AI evaluates responses on a 0-10 scale:

- **8-10**: Excellent - Strong understanding and communication
- **6-7.9**: Good - Solid answer with minor areas for improvement
- **4-5.9**: Fair - Basic understanding but needs more depth
- **0-3.9**: Needs Improvement - Significant gaps in knowledge

## 🚧 Future Enhancements

- [ ] User authentication and authorization
- [ ] Voice-based interview mode
- [ ] Code editor integration for coding questions
- [ ] Video recording of interview sessions
- [ ] Advanced analytics and performance trends
- [ ] Multiple language support
- [ ] Custom question banks
- [ ] Interview scheduling
- [ ] Peer comparison and benchmarking

## 📝 Environment Variables

### Backend (.env)
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=interviewer
DB_PASSWORD=interviewerpass
DB_NAME=ai_interviewer
GEMINI_API_KEY=your_api_key
PORT=8080
```

### Frontend (.env)
```env
VITE_API_URL=http://localhost:8080/api
```

## 👥 Authors

- Ayman Singh


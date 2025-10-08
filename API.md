# API Documentation

## Base URL
```
http://localhost:8080/api
```

## Authentication
Currently, no authentication is required. Future versions will implement JWT-based authentication.

## Endpoints

### 1. Health Check

Check if the API is running.

**Endpoint:** `GET /health`

**Response:**
```json
{
  "status": "healthy"
}
```

---

### 2. Start Interview

Create a new interview session with AI-generated questions.

**Endpoint:** `POST /interview/start`

**Request Body:**
```json
{
  "user_name": "John Doe",
  "email": "john@example.com",
  "position": "Software Engineer",
  "difficulty": "medium"
}
```

**Parameters:**
- `user_name` (string, required): Candidate's full name
- `email` (string, required): Candidate's email address
- `position` (string, required): Job position/role
- `difficulty` (string, required): One of: "easy", "medium", "hard"

**Response:** `200 OK`
```json
{
  "interview_id": 1,
  "question": {
    "id": 1,
    "interview_id": 1,
    "question_text": "What is your experience with React?",
    "question_type": "technical",
    "order": 1,
    "created_at": "2024-10-08T10:00:00Z"
  }
}
```

**Error Responses:**
- `400 Bad Request`: Missing or invalid fields
- `500 Internal Server Error`: Failed to create interview or generate questions

---

### 3. Submit Answer

Submit an answer to a question and receive AI-powered feedback.

**Endpoint:** `POST /interview/submit`

**Request Body:**
```json
{
  "question_id": 1,
  "response_text": "I have 3 years of experience building React applications..."
}
```

**Parameters:**
- `question_id` (integer, required): ID of the question being answered
- `response_text` (string, required): The candidate's answer

**Response:** `200 OK`
```json
{
  "feedback": "Great answer! You demonstrated solid understanding...",
  "score": 8.5,
  "next_question": {
    "id": 2,
    "interview_id": 1,
    "question_text": "Describe a challenging project you worked on.",
    "question_type": "behavioral",
    "order": 2,
    "created_at": "2024-10-08T10:00:00Z"
  },
  "completed": false
}
```

**When Interview is Complete:**
```json
{
  "feedback": "Excellent final answer...",
  "score": 9.0,
  "next_question": null,
  "completed": true
}
```

**Error Responses:**
- `400 Bad Request`: Invalid request payload
- `404 Not Found`: Question not found
- `500 Internal Server Error`: Failed to evaluate or store response

---

### 4. Get Interview Details

Retrieve complete details of an interview including all questions and responses.

**Endpoint:** `GET /interview/{id}`

**Parameters:**
- `id` (integer, path): Interview ID

**Response:** `200 OK`
```json
{
  "interview": {
    "id": 1,
    "user_id": 1,
    "position": "Software Engineer",
    "difficulty": "medium",
    "status": "completed",
    "score": 8.2,
    "started_at": "2024-10-08T10:00:00Z",
    "completed_at": "2024-10-08T10:30:00Z"
  },
  "questions": [
    {
      "id": 1,
      "interview_id": 1,
      "question_text": "What is your experience with React?",
      "question_type": "technical",
      "order": 1,
      "created_at": "2024-10-08T10:00:00Z"
    }
  ],
  "responses": [
    {
      "id": 1,
      "question_id": 1,
      "response_text": "I have 3 years of experience...",
      "feedback": "Great answer!...",
      "score": 8.5,
      "created_at": "2024-10-08T10:05:00Z"
    }
  ]
}
```

**Error Responses:**
- `400 Bad Request`: Invalid interview ID
- `404 Not Found`: Interview not found

---

### 5. Get User Interview History

Retrieve all interviews for a specific user.

**Endpoint:** `GET /interviews?email={email}`

**Query Parameters:**
- `email` (string, required): User's email address

**Response:** `200 OK`
```json
[
  {
    "id": 1,
    "user_id": 1,
    "position": "Software Engineer",
    "difficulty": "medium",
    "status": "completed",
    "score": 8.2,
    "started_at": "2024-10-08T10:00:00Z",
    "completed_at": "2024-10-08T10:30:00Z"
  },
  {
    "id": 2,
    "user_id": 1,
    "position": "Full Stack Developer",
    "difficulty": "hard",
    "status": "in_progress",
    "score": null,
    "started_at": "2024-10-08T14:00:00Z",
    "completed_at": null
  }
]
```

**Error Responses:**
- `400 Bad Request`: Email parameter missing
- `404 Not Found`: User not found
- `500 Internal Server Error`: Failed to retrieve interviews

---

## Data Models

### Interview Status
- `in_progress`: Interview is ongoing
- `completed`: Interview is finished

### Difficulty Levels
- `easy`: Entry level questions
- `medium`: Mid-level questions
- `hard`: Senior level questions

### Question Types
- `technical`: Technical knowledge questions
- `behavioral`: Behavioral and situational questions
- `coding`: Coding and problem-solving questions

### Score Range
- Scores are on a scale of 0-10
- 8-10: Excellent
- 6-7.9: Good
- 4-5.9: Fair
- 0-3.9: Needs Improvement

## Rate Limiting

Currently, no rate limiting is implemented. In production, consider implementing:
- Rate limiting per IP address
- Request throttling for AI API calls
- User-based quotas

## CORS

The API allows requests from:
- `http://localhost:3000` (React dev server)
- `http://localhost:5173` (Vite dev server)

## Error Handling

All errors follow this format:
```json
{
  "error": "Error message description"
}
```

Common HTTP status codes:
- `200 OK`: Successful request
- `400 Bad Request`: Invalid input
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server error

## Examples

### Starting an Interview with cURL

```bash
curl -X POST http://localhost:8080/api/interview/start \
  -H "Content-Type: application/json" \
  -d '{
    "user_name": "John Doe",
    "email": "john@example.com",
    "position": "Software Engineer",
    "difficulty": "medium"
  }'
```

### Submitting an Answer

```bash
curl -X POST http://localhost:8080/api/interview/submit \
  -H "Content-Type: application/json" \
  -d '{
    "question_id": 1,
    "response_text": "My answer goes here..."
  }'
```

### Getting Interview Details

```bash
curl http://localhost:8080/api/interview/1
```

### Getting User History

```bash
curl "http://localhost:8080/api/interviews?email=john@example.com"
```

## WebSocket Support (Future)

Future versions may include WebSocket support for:
- Real-time feedback
- Live interview sessions
- Progress updates

package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type Interview struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	Position    string     `json:"position"`
	Difficulty  string     `json:"difficulty"` // easy, medium, hard
	Status      string     `json:"status"`     // in_progress, completed
	Score       *float64   `json:"score,omitempty"`
	StartedAt   time.Time  `json:"started_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

type Question struct {
	ID           int       `json:"id"`
	InterviewID  int       `json:"interview_id"`
	QuestionText string    `json:"question_text"`
	QuestionType string    `json:"question_type"` // technical, behavioral, coding
	Order        int       `json:"order"`
	CreatedAt    time.Time `json:"created_at"`
}

type Response struct {
	ID           int       `json:"id"`
	QuestionID   int       `json:"question_id"`
	ResponseText string    `json:"response_text"`
	Feedback     string    `json:"feedback,omitempty"`
	Score        *float64  `json:"score,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}

// DTOs
type StartInterviewRequest struct {
	UserName   string `json:"user_name"`
	Email      string `json:"email"`
	Position   string `json:"position"`
	Difficulty string `json:"difficulty"`
}

type StartInterviewResponse struct {
	InterviewID int      `json:"interview_id"`
	Question    Question `json:"question"`
}

type SubmitAnswerRequest struct {
	QuestionID   int    `json:"question_id"`
	ResponseText string `json:"response_text"`
}

type SubmitAnswerResponse struct {
	Feedback     string    `json:"feedback"`
	Score        float64   `json:"score"`
	NextQuestion *Question `json:"next_question,omitempty"`
	Completed    bool      `json:"completed"`
}

type InterviewResult struct {
	Interview Interview  `json:"interview"`
	Questions []Question `json:"questions"`
	Responses []Response `json:"responses"`
}

package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ai-interviewer/backend/internal/models"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// User operations
func (r *Repository) CreateUser(name, email string) (*models.User, error) {
	// Check if user exists
	var existingUser models.User
	err := r.db.QueryRow("SELECT id, name, email, created_at FROM users WHERE email = ?", email).
		Scan(&existingUser.ID, &existingUser.Name, &existingUser.Email, &existingUser.CreatedAt)

	if err == nil {
		return &existingUser, nil
	}

	if err != sql.ErrNoRows {
		return nil, err
	}

	// Create new user
	result, err := r.db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", name, email)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:        int(id),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}

	return user, nil
}

// Interview operations
func (r *Repository) CreateInterview(userID int, position, difficulty string) (*models.Interview, error) {
	result, err := r.db.Exec(
		"INSERT INTO interviews (user_id, position, difficulty, status) VALUES (?, ?, ?, ?)",
		userID, position, difficulty, "in_progress",
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	interview := &models.Interview{
		ID:         int(id),
		UserID:     userID,
		Position:   position,
		Difficulty: difficulty,
		Status:     "in_progress",
		StartedAt:  time.Now(),
	}

	return interview, nil
}

func (r *Repository) GetInterview(id int) (*models.Interview, error) {
	var interview models.Interview
	var score sql.NullFloat64
	var completedAt sql.NullTime

	err := r.db.QueryRow(
		"SELECT id, user_id, position, difficulty, status, score, started_at, completed_at FROM interviews WHERE id = ?",
		id,
	).Scan(&interview.ID, &interview.UserID, &interview.Position, &interview.Difficulty,
		&interview.Status, &score, &interview.StartedAt, &completedAt)

	if err != nil {
		return nil, err
	}

	if score.Valid {
		interview.Score = &score.Float64
	}
	if completedAt.Valid {
		interview.CompletedAt = &completedAt.Time
	}

	return &interview, nil
}

func (r *Repository) UpdateInterviewStatus(id int, status string, score float64) error {
	now := time.Now()
	_, err := r.db.Exec(
		"UPDATE interviews SET status = ?, score = ?, completed_at = ? WHERE id = ?",
		status, score, now, id,
	)
	return err
}

func (r *Repository) GetUserInterviews(userID int) ([]models.Interview, error) {
	rows, err := r.db.Query(
		"SELECT id, user_id, position, difficulty, status, score, started_at, completed_at FROM interviews WHERE user_id = ? ORDER BY started_at DESC",
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var interviews []models.Interview
	for rows.Next() {
		var interview models.Interview
		var score sql.NullFloat64
		var completedAt sql.NullTime

		err := rows.Scan(&interview.ID, &interview.UserID, &interview.Position, &interview.Difficulty,
			&interview.Status, &score, &interview.StartedAt, &completedAt)
		if err != nil {
			return nil, err
		}

		if score.Valid {
			interview.Score = &score.Float64
		}
		if completedAt.Valid {
			interview.CompletedAt = &completedAt.Time
		}

		interviews = append(interviews, interview)
	}

	return interviews, nil
}

// Question operations
func (r *Repository) CreateQuestion(interviewID int, questionText, questionType string, order int) (*models.Question, error) {
	result, err := r.db.Exec(
		"INSERT INTO questions (interview_id, question_text, question_type, order_num) VALUES (?, ?, ?, ?)",
		interviewID, questionText, questionType, order,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	question := &models.Question{
		ID:           int(id),
		InterviewID:  interviewID,
		QuestionText: questionText,
		QuestionType: questionType,
		Order:        order,
		CreatedAt:    time.Now(),
	}

	return question, nil
}

func (r *Repository) GetQuestion(id int) (*models.Question, error) {
	var question models.Question
	err := r.db.QueryRow(
		"SELECT id, interview_id, question_text, question_type, order_num, created_at FROM questions WHERE id = ?",
		id,
	).Scan(&question.ID, &question.InterviewID, &question.QuestionText, &question.QuestionType,
		&question.Order, &question.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &question, nil
}

func (r *Repository) GetInterviewQuestions(interviewID int) ([]models.Question, error) {
	rows, err := r.db.Query(
		"SELECT id, interview_id, question_text, question_type, order_num, created_at FROM questions WHERE interview_id = ? ORDER BY order_num",
		interviewID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []models.Question
	for rows.Next() {
		var question models.Question
		err := rows.Scan(&question.ID, &question.InterviewID, &question.QuestionText,
			&question.QuestionType, &question.Order, &question.CreatedAt)
		if err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}

	return questions, nil
}

// Response operations
func (r *Repository) CreateResponse(questionID int, responseText, feedback string, score float64) (*models.Response, error) {
	result, err := r.db.Exec(
		"INSERT INTO responses (question_id, response_text, feedback, score) VALUES (?, ?, ?, ?)",
		questionID, responseText, feedback, score,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	response := &models.Response{
		ID:           int(id),
		QuestionID:   questionID,
		ResponseText: responseText,
		Feedback:     feedback,
		Score:        &score,
		CreatedAt:    time.Now(),
	}

	return response, nil
}

func (r *Repository) GetQuestionResponses(questionID int) ([]models.Response, error) {
	rows, err := r.db.Query(
		"SELECT id, question_id, response_text, feedback, score, created_at FROM responses WHERE question_id = ?",
		questionID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var responses []models.Response
	for rows.Next() {
		var response models.Response
		var score sql.NullFloat64
		err := rows.Scan(&response.ID, &response.QuestionID, &response.ResponseText,
			&response.Feedback, &score, &response.CreatedAt)
		if err != nil {
			return nil, err
		}

		if score.Valid {
			response.Score = &score.Float64
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (r *Repository) GetInterviewResult(interviewID int) (*models.InterviewResult, error) {
	interview, err := r.GetInterview(interviewID)
	if err != nil {
		return nil, fmt.Errorf("failed to get interview: %w", err)
	}

	questions, err := r.GetInterviewQuestions(interviewID)
	if err != nil {
		return nil, fmt.Errorf("failed to get questions: %w", err)
	}

	var responses []models.Response
	for _, q := range questions {
		qResponses, err := r.GetQuestionResponses(q.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get responses: %w", err)
		}
		responses = append(responses, qResponses...)
	}

	// Initialize empty array if no responses
	if responses == nil {
		responses = []models.Response{}
	}

	return &models.InterviewResult{
		Interview: *interview,
		Questions: questions,
		Responses: responses,
	}, nil
}

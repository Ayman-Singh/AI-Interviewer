package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ai-interviewer/backend/internal/ai"
	"github.com/ai-interviewer/backend/internal/models"
	"github.com/ai-interviewer/backend/internal/repository"
	"github.com/gorilla/mux"
)

type Handler struct {
	repo      *repository.Repository
	aiService *ai.AIService
}

func New(repo *repository.Repository, aiService *ai.AIService) *Handler {
	return &Handler{
		repo:      repo,
		aiService: aiService,
	}
}

func (h *Handler) StartInterview(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	
	var req models.StartInterviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate request
	if req.UserName == "" || req.Email == "" || req.Position == "" || req.Difficulty == "" {
		respondWithError(w, http.StatusBadRequest, "Missing required fields")
		return
	}

	// Create or get user
	user, err := h.repo.CreateUser(req.UserName, req.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	// Create interview
	interview, err := h.repo.CreateInterview(user.ID, req.Position, req.Difficulty)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create interview")
		return
	}

	// Generate questions using AI
	ctx := context.Background()
	questionTexts, err := h.aiService.GenerateQuestions(ctx, req.Position, req.Difficulty, 5)
	if err != nil {
		log.Printf("AI service error: %v", err)
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to generate questions: %v", err))
		return
	}

	// Store questions in database
	var questions []models.Question
	for i, qText := range questionTexts {
		qType := determineQuestionType(i)
		question, err := h.repo.CreateQuestion(interview.ID, qText, qType, i+1)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to store questions")
			return
		}
		questions = append(questions, *question)
	}

	// Return first question
	response := models.StartInterviewResponse{
		InterviewID: interview.ID,
		Question:    questions[0],
	}

	respondWithJSON(w, http.StatusOK, response)
}

func (h *Handler) SubmitAnswer(w http.ResponseWriter, r *http.Request) {
	var req models.SubmitAnswerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Get question
	question, err := h.repo.GetQuestion(req.QuestionID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Question not found")
		return
	}

	// Evaluate answer using AI
	ctx := context.Background()
	feedback, score, err := h.aiService.EvaluateAnswer(ctx, question.QuestionText, req.ResponseText)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to evaluate answer")
		return
	}

	// Store response
	_, err = h.repo.CreateResponse(req.QuestionID, req.ResponseText, feedback, score)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to store response")
		return
	}

	// Get all questions for this interview
	questions, err := h.repo.GetInterviewQuestions(question.InterviewID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to get questions")
		return
	}

	// Check if there are more questions
	var nextQuestion *models.Question
	completed := false

	for i, q := range questions {
		if q.ID == req.QuestionID && i < len(questions)-1 {
			nextQuestion = &questions[i+1]
			break
		}
	}

	if nextQuestion == nil {
		// Interview completed - calculate average score
		completed = true
		totalScore := 0.0
		count := 0

		for _, q := range questions {
			responses, err := h.repo.GetQuestionResponses(q.ID)
			if err == nil && len(responses) > 0 && responses[0].Score != nil {
				totalScore += *responses[0].Score
				count++
			}
		}

		avgScore := 0.0
		if count > 0 {
			avgScore = totalScore / float64(count)
		}

		// Update interview status
		err = h.repo.UpdateInterviewStatus(question.InterviewID, "completed", avgScore)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to update interview")
			return
		}
	}

	response := models.SubmitAnswerResponse{
		Feedback:     feedback,
		Score:        score,
		NextQuestion: nextQuestion,
		Completed:    completed,
	}

	respondWithJSON(w, http.StatusOK, response)
}

func (h *Handler) GetInterview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid interview ID")
		return
	}

	result, err := h.repo.GetInterviewResult(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Interview not found")
		return
	}

	respondWithJSON(w, http.StatusOK, result)
}

func (h *Handler) GetUserInterviews(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		respondWithError(w, http.StatusBadRequest, "Email parameter is required")
		return
	}

	// Get user by email
	user, err := h.repo.CreateUser("", email) // Will return existing user
	if err != nil {
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	interviews, err := h.repo.GetUserInterviews(user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to get interviews")
		return
	}

	respondWithJSON(w, http.StatusOK, interviews)
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "healthy"})
}

// Helper functions
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func determineQuestionType(index int) string {
	types := []string{"technical", "behavioral", "technical", "coding", "behavioral"}
	if index < len(types) {
		return types[index]
	}
	return "technical"
}

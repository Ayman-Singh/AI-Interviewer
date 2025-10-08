package ai

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type AIService struct {
	client *genai.Client
	model  *genai.GenerativeModel
}

func NewAIService(apiKey string) (*AIService, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Gemini: %w", err)
	}

	// Use gemini-2.5-flash - stable version released June 2025
	model := client.GenerativeModel("gemini-2.5-flash")

	return &AIService{
		client: client,
		model:  model,
	}, nil
}

func (s *AIService) GenerateQuestions(ctx context.Context, position, difficulty string, count int) ([]string, error) {
	prompt := fmt.Sprintf(`You are an expert technical interviewer. Generate %d interview questions for a %s position with %s difficulty level.

Mix the questions between:
- Technical knowledge questions
- Behavioral questions
- Problem-solving scenarios

Return ONLY the questions, one per line, numbered 1. 2. 3. etc.
Do not include any other text or explanations.

Position: %s
Difficulty: %s
Number of questions: %d`, count, position, difficulty, position, difficulty, count)

	resp, err := s.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, fmt.Errorf("failed to generate questions: %w", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("no response from AI")
	}

	response := fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0])
	questions := parseQuestions(response)
	if len(questions) == 0 {
		return nil, fmt.Errorf("no questions generated")
	}

	return questions, nil
}

func (s *AIService) EvaluateAnswer(ctx context.Context, question, answer string) (string, float64, error) {
	prompt := fmt.Sprintf(`You are an expert interviewer evaluating a candidate's response.

Question: %s

Candidate's Answer: %s

Please provide:
1. A score from 0-10 (where 10 is excellent)
2. Constructive feedback (2-3 sentences)

Format your response EXACTLY as:
Score: X
Feedback: Your feedback here

Be constructive and specific in your feedback.`, question, answer)

	resp, err := s.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", 0, fmt.Errorf("failed to evaluate answer: %w", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "Unable to evaluate answer at this time.", 5.0, nil
	}

	response := fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0])
	score, feedback := parseEvaluation(response)
	return feedback, score, nil
}

func (s *AIService) GenerateFinalFeedback(ctx context.Context, position string, averageScore float64, totalQuestions int) (string, error) {
	prompt := fmt.Sprintf(`You are an expert interviewer providing final feedback for a candidate.

Position: %s
Average Score: %.2f/10
Total Questions: %d

Provide a comprehensive summary (3-4 sentences) that includes:
1. Overall performance assessment
2. Key strengths observed
3. Areas for improvement
4. Recommendation (hire/consider/not recommended)

Be professional, constructive, and specific.`, position, averageScore, totalQuestions)

	resp, err := s.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("failed to generate final feedback: %w", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "Thank you for completing the interview.", nil
	}

	response := fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0])
	return strings.TrimSpace(response), nil
}

func parseQuestions(response string) []string {
	lines := strings.Split(response, "\n")
	var questions []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Remove numbering (1., 2., etc.)
		parts := strings.SplitN(line, ".", 2)
		if len(parts) == 2 {
			question := strings.TrimSpace(parts[1])
			if question != "" {
				questions = append(questions, question)
			}
		} else if line != "" {
			questions = append(questions, line)
		}
	}

	return questions
}

func parseEvaluation(response string) (float64, string) {
	lines := strings.Split(response, "\n")
	var score float64 = 5.0 // default score
	var feedback string

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(strings.ToLower(line), "score:") {
			scoreStr := strings.TrimSpace(strings.TrimPrefix(strings.ToLower(line), "score:"))
			// Parse score (handle formats like "8", "8.5", "8/10")
			var parsedScore float64
			if strings.Contains(scoreStr, "/") {
				fmt.Sscanf(scoreStr, "%f/", &parsedScore)
			} else {
				fmt.Sscanf(scoreStr, "%f", &parsedScore)
			}
			if parsedScore >= 0 && parsedScore <= 10 {
				score = parsedScore
			}
		} else if strings.HasPrefix(strings.ToLower(line), "feedback:") {
			feedback = strings.TrimSpace(strings.TrimPrefix(line, "Feedback:"))
			feedback = strings.TrimSpace(strings.TrimPrefix(feedback, "feedback:"))
		} else if feedback != "" {
			// Append additional feedback lines
			feedback += " " + line
		}
	}

	if feedback == "" {
		feedback = "Good effort. Keep practicing to improve your skills."
	}

	return score, strings.TrimSpace(feedback)
}

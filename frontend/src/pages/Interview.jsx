import { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { interviewAPI } from '../services/api';
import './Interview.css';

function Interview() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [loading, setLoading] = useState(true);
  const [submitting, setSubmitting] = useState(false);
  const [currentQuestion, setCurrentQuestion] = useState(null);
  const [answer, setAnswer] = useState('');
  const [feedback, setFeedback] = useState(null);
  const [questionNumber, setQuestionNumber] = useState(1);
  const [error, setError] = useState('');

  useEffect(() => {
    loadInterview();
  }, [id]);

  const loadInterview = async () => {
    try {
      const data = await interviewAPI.getInterview(id);
      console.log('Interview data received:', data);
      console.log('Questions:', data.questions);
      console.log('Responses:', data.responses);
      
      // Check if data structure is valid
      if (!data.questions || !Array.isArray(data.questions) || data.questions.length === 0) {
        console.error('No questions found in interview data');
        setError('No questions found for this interview');
        setLoading(false);
        return;
      }
      
      // Find the first unanswered question
      const responses = data.responses || []; // Handle null responses
      const unansweredQuestion = data.questions.find(q => {
        const hasResponse = responses.some(r => r.question_id === q.id);
        return !hasResponse;
      });

      console.log('Unanswered question:', unansweredQuestion);

      if (unansweredQuestion) {
        setCurrentQuestion(unansweredQuestion);
        const questionIndex = data.questions.findIndex(q => q.id === unansweredQuestion.id);
        setQuestionNumber(questionIndex + 1);
      } else {
        // All questions answered, go to results
        navigate(`/results/${id}`);
      }
    } catch (err) {
      console.error('Error loading interview:', err);
      console.error('Error response:', err.response);
      setError('Failed to load interview: ' + (err.response?.data?.error || err.message));
    } finally {
      setLoading(false);
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    
    if (!answer.trim()) {
      setError('Please provide an answer');
      return;
    }

    setSubmitting(true);
    setError('');

    try {
      const response = await interviewAPI.submitAnswer({
        question_id: currentQuestion.id,
        response_text: answer,
      });

      setFeedback({
        feedback: response.feedback,
        score: response.score,
      });

      // Wait a moment to show feedback, then move to next question or results
      setTimeout(() => {
        if (response.completed) {
          navigate(`/results/${id}`);
        } else if (response.next_question) {
          setCurrentQuestion(response.next_question);
          setQuestionNumber(questionNumber + 1);
          setAnswer('');
          setFeedback(null);
        }
      }, 3000);
    } catch (err) {
      setError(err.response?.data?.error || 'Failed to submit answer');
    } finally {
      setSubmitting(false);
    }
  };

  if (loading) {
    return (
      <div className="loading">
        <div className="spinner"></div>
      </div>
    );
  }

  if (error && !currentQuestion) {
    return (
      <div className="interview">
        <div className="error-message">{error}</div>
      </div>
    );
  }

  return (
    <div className="interview">
      <div className="interview-container">
        <div className="interview-header">
          <div className="progress-info">
            <span className="question-counter">Question {questionNumber}</span>
            <span className="question-type">{currentQuestion?.question_type}</span>
          </div>
        </div>

        <div className="question-card card">
          <h2 className="question-text">{currentQuestion?.question_text}</h2>
        </div>

        {feedback ? (
          <div className="feedback-section">
            <div className="feedback-card success-message">
              <div className="feedback-header">
                <h3>Feedback</h3>
                <div className="score-badge">
                  Score: {feedback.score.toFixed(1)}/10
                </div>
              </div>
              <p className="feedback-text">{feedback.feedback}</p>
              <p className="next-question-info">Moving to next question...</p>
            </div>
          </div>
        ) : (
          <div className="answer-section">
            <form onSubmit={handleSubmit}>
              {error && <div className="error-message">{error}</div>}
              
              <div className="form-group">
                <label className="form-label" htmlFor="answer">
                  Your Answer
                </label>
                <textarea
                  id="answer"
                  className="form-textarea"
                  value={answer}
                  onChange={(e) => setAnswer(e.target.value)}
                  placeholder="Type your answer here..."
                  rows="8"
                  disabled={submitting}
                  required
                />
              </div>

              <div className="answer-actions">
                <button
                  type="submit"
                  className="btn btn-primary"
                  disabled={submitting || !answer.trim()}
                >
                  {submitting ? 'Submitting...' : 'Submit Answer'}
                </button>
              </div>
            </form>
          </div>
        )}
      </div>
    </div>
  );
}

export default Interview;

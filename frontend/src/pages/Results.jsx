import { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { interviewAPI } from '../services/api';
import './Results.css';

function Results() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [loading, setLoading] = useState(true);
  const [result, setResult] = useState(null);
  const [error, setError] = useState('');

  useEffect(() => {
    loadResults();
  }, [id]);

  const loadResults = async () => {
    try {
      const data = await interviewAPI.getInterview(id);
      setResult(data);
    } catch (err) {
      setError('Failed to load results');
    } finally {
      setLoading(false);
    }
  };

  const getScoreColor = (score) => {
    if (score >= 8) return '#00ff41';
    if (score >= 6) return '#ffff00';
    return '#ff4444';
  };

  const getPerformanceLevel = (score) => {
    if (score >= 8) return 'Excellent';
    if (score >= 6) return 'Good';
    if (score >= 4) return 'Fair';
    return 'Needs Improvement';
  };

  if (loading) {
    return (
      <div className="loading">
        <div className="spinner"></div>
      </div>
    );
  }

  if (error || !result) {
    return (
      <div className="results">
        <div className="error-message">{error || 'No results found'}</div>
      </div>
    );
  }

  const averageScore = result.interview.score || 0;

  return (
    <div className="results">
      <div className="results-container">
        <div className="results-header card">
          <h1 className="results-title">Interview Complete!</h1>
          <div className="overall-score">
            <div 
              className="score-circle"
              style={{ borderColor: getScoreColor(averageScore) }}
            >
              <span className="score-value" style={{ color: getScoreColor(averageScore) }}>
                {averageScore.toFixed(1)}
              </span>
              <span className="score-max">/10</span>
            </div>
            <div className="score-info">
              <h2 style={{ color: getScoreColor(averageScore) }}>
                {getPerformanceLevel(averageScore)}
              </h2>
              <p className="interview-meta">
                Position: {result.interview.position} | 
                Difficulty: {result.interview.difficulty}
              </p>
            </div>
          </div>
        </div>

        <div className="questions-results">
          <h2 className="section-title">Question-by-Question Breakdown</h2>
          
          {result.questions.map((question, index) => {
            const response = result.responses.find(r => r.question_id === question.id);
            
            return (
              <div key={question.id} className="question-result card">
                <div className="question-result-header">
                  <span className="question-number">Question {index + 1}</span>
                  <span className="question-type-badge">{question.question_type}</span>
                  {response?.score && (
                    <span 
                      className="question-score"
                      style={{ color: getScoreColor(response.score) }}
                    >
                      {response.score.toFixed(1)}/10
                    </span>
                  )}
                </div>
                
                <div className="question-content">
                  <h3 className="question-text-small">{question.question_text}</h3>
                  
                  {response && (
                    <>
                      <div className="answer-section-result">
                        <h4 className="subsection-title">Your Answer:</h4>
                        <p className="answer-text">{response.response_text}</p>
                      </div>
                      
                      <div className="feedback-section-result">
                        <h4 className="subsection-title">Feedback:</h4>
                        <p className="feedback-text-result">{response.feedback}</p>
                      </div>
                    </>
                  )}
                </div>
              </div>
            );
          })}
        </div>

        <div className="results-actions">
          <button
            className="btn btn-primary"
            onClick={() => navigate('/')}
          >
            Start New Interview
          </button>
          <button
            className="btn"
            onClick={() => window.print()}
          >
            Print Results
          </button>
        </div>
      </div>
    </div>
  );
}

export default Results;

import { useState, useEffect } from 'react';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { interviewAPI } from '../services/api';
import './History.css';

function History() {
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();
  const [email, setEmail] = useState(searchParams.get('email') || '');
  const [interviews, setInterviews] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  useEffect(() => {
    if (email) {
      loadHistory();
    }
  }, []);

  const loadHistory = async () => {
    if (!email) {
      setError('Please enter an email address');
      return;
    }

    setLoading(true);
    setError('');

    try {
      const data = await interviewAPI.getUserInterviews(email);
      setInterviews(data || []);
      
      if (!data || data.length === 0) {
        setError('No interview history found for this email');
      }
    } catch (err) {
      setError(err.response?.data?.error || 'Failed to load history');
    } finally {
      setLoading(false);
    }
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    loadHistory();
  };

  const formatDate = (dateString) => {
    const date = new Date(dateString);
    return date.toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    });
  };

  const getStatusBadge = (status) => {
    return status === 'completed' ? 'status-completed' : 'status-progress';
  };

  const getScoreColor = (score) => {
    if (score >= 8) return '#00ff41';
    if (score >= 6) return '#ffff00';
    return '#ff4444';
  };

  return (
    <div className="history">
      <div className="history-container">
        <h1 className="history-title">Interview History</h1>

        <div className="search-section card">
          <form onSubmit={handleSubmit}>
            <div className="search-form">
              <input
                type="email"
                className="form-input"
                placeholder="Enter your email address"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                required
              />
              <button type="submit" className="btn btn-primary" disabled={loading}>
                {loading ? 'Loading...' : 'Search'}
              </button>
            </div>
          </form>
        </div>

        {error && !loading && (
          <div className={interviews.length === 0 ? "error-message" : "info-message"}>
            {error}
          </div>
        )}

        {loading && (
          <div className="loading">
            <div className="spinner"></div>
          </div>
        )}

        {!loading && interviews.length > 0 && (
          <div className="interviews-list">
            {interviews.map((interview) => (
              <div key={interview.id} className="interview-card card">
                <div className="interview-card-header">
                  <div className="interview-info">
                    <h3 className="interview-position">{interview.position}</h3>
                    <p className="interview-date">{formatDate(interview.started_at)}</p>
                  </div>
                  <span className={`status-badge ${getStatusBadge(interview.status)}`}>
                    {interview.status}
                  </span>
                </div>

                <div className="interview-card-body">
                  <div className="interview-details">
                    <div className="detail-item">
                      <span className="detail-label">Difficulty:</span>
                      <span className="detail-value">{interview.difficulty}</span>
                    </div>
                    
                    {interview.status === 'completed' && interview.score && (
                      <div className="detail-item">
                        <span className="detail-label">Score:</span>
                        <span 
                          className="detail-value score-value"
                          style={{ color: getScoreColor(interview.score) }}
                        >
                          {interview.score.toFixed(1)}/10
                        </span>
                      </div>
                    )}
                  </div>

                  <div className="interview-actions">
                    {interview.status === 'completed' ? (
                      <button
                        className="btn"
                        onClick={() => navigate(`/results/${interview.id}`)}
                      >
                        View Results
                      </button>
                    ) : (
                      <button
                        className="btn btn-primary"
                        onClick={() => navigate(`/interview/${interview.id}`)}
                      >
                        Continue Interview
                      </button>
                    )}
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}

        <div className="back-action">
          <button className="btn" onClick={() => navigate('/')}>
            Back to Home
          </button>
        </div>
      </div>
    </div>
  );
}

export default History;

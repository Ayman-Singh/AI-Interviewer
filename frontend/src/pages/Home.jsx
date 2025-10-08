import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { interviewAPI } from '../services/api';
import './Home.css';

function Home() {
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    user_name: '',
    email: '',
    position: '',
    difficulty: 'medium',
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setLoading(true);

    try {
      const response = await interviewAPI.startInterview(formData);
      navigate(`/interview/${response.interview_id}`);
    } catch (err) {
      setError(err.response?.data?.error || 'Failed to start interview. Please try again.');
    } finally {
      setLoading(false);
    }
  };

  const handleViewHistory = () => {
    if (formData.email) {
      navigate(`/history?email=${encodeURIComponent(formData.email)}`);
    } else {
      setError('Please enter your email to view history');
    }
  };

  return (
    <div className="home">
      <div className="home-container">
        <div className="welcome-section">
          <div className="brand-logo">
            <svg viewBox="0 0 200 200" className="logo-icon">
              <defs>
                <linearGradient id="logoGradient" x1="0%" y1="0%" x2="100%" y2="100%">
                  <stop offset="0%" stopColor="#a855f7" />
                  <stop offset="50%" stopColor="#7c3aed" />
                  <stop offset="100%" stopColor="#6366f1" />
                </linearGradient>
              </defs>
              
              {/* Microphone base */}
              <circle cx="100" cy="75" r="25" fill="url(#logoGradient)" opacity="0.2"/>
              <rect x="85" y="50" width="30" height="50" rx="15" fill="url(#logoGradient)"/>
              
              {/* Microphone stand */}
              <path d="M 100 100 Q 100 115 85 125" stroke="url(#logoGradient)" strokeWidth="4" fill="none" strokeLinecap="round"/>
              <line x1="70" y1="125" x2="100" y2="125" stroke="url(#logoGradient)" strokeWidth="4" strokeLinecap="round"/>
              
              {/* AI Brain waves */}
              <path d="M 130 60 Q 145 60 145 75" stroke="url(#logoGradient)" strokeWidth="3" fill="none" opacity="0.6" strokeLinecap="round"/>
              <path d="M 135 50 Q 155 50 155 70" stroke="url(#logoGradient)" strokeWidth="2.5" fill="none" opacity="0.4" strokeLinecap="round"/>
              <path d="M 130 90 Q 145 90 145 75" stroke="url(#logoGradient)" strokeWidth="3" fill="none" opacity="0.6" strokeLinecap="round"/>
              <path d="M 135 100 Q 155 100 155 80" stroke="url(#logoGradient)" strokeWidth="2.5" fill="none" opacity="0.4" strokeLinecap="round"/>
            </svg>
            <h1 className="brand-name">InterviewAI</h1>
          </div>
          
          <h2 className="welcome-title">Master Your Next Interview</h2>
          <p className="welcome-text">
            Experience intelligent interview preparation powered by advanced AI. 
            Get personalized questions, real-time feedback, and actionable insights 
            to elevate your interview performance.
          </p>
          
          <div className="features">
            <div className="feature">
              <div className="feature-icon-wrapper">
                <svg viewBox="0 0 24 24" className="feature-svg">
                  <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm-1-13h2v6h-2zm0 8h2v2h-2z" fill="currentColor"/>
                </svg>
              </div>
              <h3>Adaptive Intelligence</h3>
              <p>Dynamic questions tailored to your role and experience level</p>
            </div>
            
            <div className="feature">
              <div className="feature-icon-wrapper">
                <svg viewBox="0 0 24 24" className="feature-svg">
                  <path d="M19 3H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zM9 17H7v-7h2v7zm4 0h-2V7h2v10zm4 0h-2v-4h2v4z" fill="currentColor"/>
                </svg>
              </div>
              <h3>Instant Analytics</h3>
              <p>Comprehensive scoring with detailed performance breakdown</p>
            </div>
            
            <div className="feature">
              <div className="feature-icon-wrapper">
                <svg viewBox="0 0 24 24" className="feature-svg">
                  <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z" fill="currentColor"/>
                </svg>
              </div>
              <h3>Industry Standards</h3>
              <p>Questions aligned with real-world interview scenarios</p>
            </div>
          </div>
        </div>

        <div className="form-section">
          <div className="card">
            <h2 className="form-title">Start Your Interview</h2>
            
            {error && <div className="error-message">{error}</div>}

            <form onSubmit={handleSubmit}>
              <div className="form-group">
                <label className="form-label" htmlFor="user_name">
                  Full Name
                </label>
                <input
                  type="text"
                  id="user_name"
                  name="user_name"
                  className="form-input"
                  value={formData.user_name}
                  onChange={handleChange}
                  required
                  placeholder="John Doe"
                />
              </div>

              <div className="form-group">
                <label className="form-label" htmlFor="email">
                  Email Address
                </label>
                <input
                  type="email"
                  id="email"
                  name="email"
                  className="form-input"
                  value={formData.email}
                  onChange={handleChange}
                  required
                  placeholder="john@example.com"
                />
              </div>

              <div className="form-group">
                <label className="form-label" htmlFor="position">
                  Position / Role
                </label>
                <input
                  type="text"
                  id="position"
                  name="position"
                  className="form-input"
                  value={formData.position}
                  onChange={handleChange}
                  required
                  placeholder="e.g., Software Engineer, Data Scientist"
                />
              </div>

              <div className="form-group">
                <label className="form-label" htmlFor="difficulty">
                  Difficulty Level
                </label>
                <select
                  id="difficulty"
                  name="difficulty"
                  className="form-select"
                  value={formData.difficulty}
                  onChange={handleChange}
                  required
                >
                  <option value="easy">Easy - Entry Level</option>
                  <option value="medium">Medium - Mid Level</option>
                  <option value="hard">Hard - Senior Level</option>
                </select>
              </div>

              <div className="form-actions">
                <button
                  type="submit"
                  className="btn btn-primary"
                  disabled={loading}
                >
                  {loading ? 'Starting...' : 'Start Interview'}
                </button>
                
                <button
                  type="button"
                  className="btn"
                  onClick={handleViewHistory}
                  disabled={loading || !formData.email}
                >
                  View History
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Home;

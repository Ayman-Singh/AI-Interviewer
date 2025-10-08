import { Routes, Route, Link, useLocation } from 'react-router-dom';
import './App.css';
import Home from './pages/Home';
import Interview from './pages/Interview';
import Results from './pages/Results';
import History from './pages/History';

function App() {
  const location = useLocation();
  const isHomePage = location.pathname === '/';

  return (
    <div className="app">
      <nav className="app-nav">
        <div className="nav-container">
          <Link to="/" className="nav-brand">
            <svg viewBox="0 0 200 200" className="nav-logo-icon">
              <defs>
                <linearGradient id="navLogoGradient" x1="0%" y1="0%" x2="100%" y2="100%">
                  <stop offset="0%" stopColor="#3b82f6" />
                  <stop offset="100%" stopColor="#8b5cf6" />
                </linearGradient>
              </defs>
              <circle cx="100" cy="75" r="25" fill="url(#navLogoGradient)" opacity="0.2"/>
              <rect x="85" y="50" width="30" height="50" rx="15" fill="url(#navLogoGradient)"/>
              <path d="M 100 100 Q 100 115 85 125" stroke="url(#navLogoGradient)" strokeWidth="4" fill="none" strokeLinecap="round"/>
              <line x1="70" y1="125" x2="100" y2="125" stroke="url(#navLogoGradient)" strokeWidth="4" strokeLinecap="round"/>
              <path d="M 130 60 Q 145 60 145 75" stroke="url(#navLogoGradient)" strokeWidth="3" fill="none" opacity="0.6" strokeLinecap="round"/>
              <path d="M 135 50 Q 155 50 155 70" stroke="url(#navLogoGradient)" strokeWidth="2.5" fill="none" opacity="0.4" strokeLinecap="round"/>
              <path d="M 130 90 Q 145 90 145 75" stroke="url(#navLogoGradient)" strokeWidth="3" fill="none" opacity="0.6" strokeLinecap="round"/>
              <path d="M 135 100 Q 155 100 155 80" stroke="url(#navLogoGradient)" strokeWidth="2.5" fill="none" opacity="0.4" strokeLinecap="round"/>
            </svg>
            <span className="nav-brand-name">InterviewAI</span>
          </Link>
          
          <div className="nav-links">
            {!isHomePage && (
              <>
                <Link to="/" className="nav-link">
                  <svg viewBox="0 0 24 24" className="nav-link-icon">
                    <path d="M10 20v-6h4v6h5v-8h3L12 3 2 12h3v8z" fill="currentColor"/>
                  </svg>
                  Home
                </Link>
              </>
            )}
          </div>
        </div>
      </nav>
      
      <main className="app-main">
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/interview/:id" element={<Interview />} />
          <Route path="/results/:id" element={<Results />} />
          <Route path="/history" element={<History />} />
        </Routes>
      </main>

      <footer className="app-footer">
        <p>&copy; 2024 AI Interviewer. Prepare smarter with AI.</p>
      </footer>
    </div>
  );
}

export default App;

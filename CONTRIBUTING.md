# Contributing to AI Interviewer

Thank you for your interest in contributing to AI Interviewer! This document provides guidelines and instructions for contributing.

## Code of Conduct

- Be respectful and inclusive
- Provide constructive feedback
- Focus on the issue, not the person
- Help others learn and grow

## How to Contribute

### Reporting Bugs

Before submitting a bug report:
1. Check if the issue already exists
2. Verify you're using the latest version
3. Reproduce the bug in a clean environment

When submitting a bug report, include:
- Clear description of the issue
- Steps to reproduce
- Expected vs actual behavior
- Environment details (OS, Docker version, etc.)
- Screenshots if applicable
- Error logs

### Suggesting Enhancements

When suggesting enhancements:
1. Check if it's already been suggested
2. Explain the use case clearly
3. Describe the expected behavior
4. Consider backwards compatibility

### Pull Requests

#### Setup Development Environment

1. Fork the repository
2. Clone your fork
3. Set up the development environment:

```powershell
# Install backend dependencies
cd backend
go mod download

# Install frontend dependencies
cd ../frontend
npm install
```

#### Making Changes

1. Create a feature branch:
```bash
git checkout -b feature/your-feature-name
```

2. Make your changes following the coding standards

3. Test your changes thoroughly:
```powershell
# Test backend
cd backend
go test ./...

# Test frontend
cd ../frontend
npm test
```

4. Commit your changes:
```bash
git commit -m "feat: add amazing feature"
```

#### Commit Message Format

Follow conventional commits:

- `feat:` New feature
- `fix:` Bug fix
- `docs:` Documentation changes
- `style:` Code style changes (formatting, etc.)
- `refactor:` Code refactoring
- `test:` Adding or updating tests
- `chore:` Maintenance tasks

Examples:
```
feat: add voice recording for interviews
fix: correct scoring calculation bug
docs: update API documentation
refactor: reorganize component structure
```

#### Pull Request Process

1. Update documentation if needed
2. Add tests for new features
3. Ensure all tests pass
4. Update the CHANGELOG.md
5. Submit the pull request with:
   - Clear title and description
   - Reference to related issues
   - Screenshots for UI changes
   - Test coverage information

## Coding Standards

### Go (Backend)

- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use `gofmt` for formatting
- Write meaningful variable and function names
- Add comments for exported functions
- Handle errors explicitly
- Use interfaces where appropriate

Example:
```go
// EvaluateAnswer evaluates a candidate's answer using AI
func (s *AIService) EvaluateAnswer(ctx context.Context, question, answer string) (string, float64, error) {
    if question == "" || answer == "" {
        return "", 0, fmt.Errorf("question and answer are required")
    }
    // Implementation...
}
```

### JavaScript/React (Frontend)

- Use functional components and hooks
- Follow React best practices
- Use meaningful component and variable names
- Keep components small and focused
- Add PropTypes or TypeScript for type checking
- Use CSS modules or styled-components

Example:
```jsx
function QuestionCard({ question, onSubmit }) {
  const [answer, setAnswer] = useState('');
  
  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit(answer);
  };
  
  return (
    <form onSubmit={handleSubmit}>
      {/* Component content */}
    </form>
  );
}
```

### CSS

- Follow BEM naming convention
- Use CSS custom properties for theming
- Mobile-first approach
- Consistent spacing and sizing

### Database

- Use migrations for schema changes
- Write efficient queries
- Add appropriate indexes
- Document complex queries

## Testing

### Backend Tests

Create tests in `*_test.go` files:

```go
func TestEvaluateAnswer(t *testing.T) {
    service := NewAIService("test-key")
    
    feedback, score, err := service.EvaluateAnswer(
        context.Background(),
        "What is React?",
        "React is a JavaScript library...",
    )
    
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    
    if score < 0 || score > 10 {
        t.Errorf("Score out of range: %f", score)
    }
}
```

### Frontend Tests

Use React Testing Library:

```jsx
import { render, screen } from '@testing-library/react';
import QuestionCard from './QuestionCard';

test('renders question text', () => {
  render(<QuestionCard question="What is React?" />);
  expect(screen.getByText('What is React?')).toBeInTheDocument();
});
```

## Documentation

- Update README.md for major changes
- Update API.md for API changes
- Add inline comments for complex logic
- Include examples in documentation

## Release Process

1. Update version numbers
2. Update CHANGELOG.md
3. Create a release tag
4. Build and test Docker images
5. Update documentation
6. Announce the release

## Questions?

If you have questions:
- Open a discussion on GitHub
- Check existing issues and discussions
- Reach out to maintainers

## Thank You!

Your contributions make this project better for everyone. Thank you for taking the time to contribute!

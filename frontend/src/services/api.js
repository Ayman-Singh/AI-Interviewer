import axios from 'axios';

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

export const interviewAPI = {
  // Start a new interview
  startInterview: async (data) => {
    const response = await api.post('/interview/start', data);
    return response.data;
  },

  // Submit an answer
  submitAnswer: async (data) => {
    const response = await api.post('/interview/submit', data);
    return response.data;
  },

  // Get interview details
  getInterview: async (id) => {
    const response = await api.get(`/interview/${id}`);
    return response.data;
  },

  // Get user interviews
  getUserInterviews: async (email) => {
    const response = await api.get('/interviews', {
      params: { email }
    });
    return response.data;
  },

  // Health check
  healthCheck: async () => {
    const response = await api.get('/health');
    return response.data;
  },
};

export default api;

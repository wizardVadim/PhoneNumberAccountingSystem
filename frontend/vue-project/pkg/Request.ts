const API_BASE_URL = '/api'

interface RequestOptions {
  method?: 'GET' | 'POST' | 'PUT' | 'DELETE'
  body?: any
  headers?: Record<string, string>
}

export const request = async (endpoint: string, options: RequestOptions = {}) => {
  const { method = 'POST', body, headers = {} } = options
  
  const token = localStorage.getItem('token')
  
  const defaultHeaders: Record<string, string> = {
    'Content-Type': 'application/json',
    ...headers
  }
  
  if (token) {
    defaultHeaders['Authorization'] = `Bearer ${token}`
  }
  
  const config: RequestInit = {
    method,
    headers: defaultHeaders,
    credentials: 'include'
  }
  
  if (body && method !== 'GET') {
    config.body = JSON.stringify(body)
  }
  
  const response = await fetch(`${API_BASE_URL}${endpoint}`, config)
  
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`)
  }
  
  return response.json()
}
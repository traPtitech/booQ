import axios from 'axios'

axios.defaults.baseURL = process.env.NODE_ENV === 'development' ? 'http://localhost:3000' : process.env.VUE_APP_API_ENDPOINT

export function getMe () {
  return axios.get('/api/users/me')
}

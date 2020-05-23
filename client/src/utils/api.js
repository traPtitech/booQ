import axios from 'axios'
import { randomString, pkce } from './hash'

// axios.defaults.withCredentials = true
export const traQBaseURL = 'https://q.trap.jp/api/1.0'
export const traQClientID = process.env.VUE_APP_API_CLIENT_ID
axios.defaults.baseURL = process.env.NODE_ENV === 'development' ? 'http://localhost:3000' : process.env.VUE_APP_API_ENDPOINT

export function getMe () {
  return axios.get('/api/users/me')
}

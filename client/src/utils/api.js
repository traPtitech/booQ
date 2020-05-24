import axios from 'axios'

axios.defaults.baseURL = '/'

export function getMe () {
  return axios.get('/api/users/me')
}

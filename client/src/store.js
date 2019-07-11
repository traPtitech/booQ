import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from 'vuex-persistedstate'
import { setAuthToken } from './utils/api'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    me: null,
    drawer: null,
    color: 'success',
    sidebarBackgroundColor: 'rgba(27, 27, 27, 0.74)',
    loginDialog: false,
    authToken: null
  },
  mutations: {
    setMe (state, data) {
      state.me = data
    },
    setDrawer (state, data) {
      state.drawer = data
    },
    setColor (state, data) {
      state.color = data
    },
    setToken (state, data) {
      state.authToken = data
      setAuthToken(data)
    },
    toggleDrawer (state) {
      state.drawer = !state.drawer
    },
    toggleLoginDialog (state) {
      state.loginDialog = !state.loginDialog
    }
  },
  actions: {
    //
  },
  plugins: [createPersistedState({
    paths: ['authToken']
  })]
})

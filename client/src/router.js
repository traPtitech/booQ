import Vue from 'vue'
import Router from 'vue-router'
import store from './store'
import Home from './views/Home.vue'
import { fetchAuthToken, setAuthToken, getMe } from './utils/api'

setAuthToken(store.state.authToken)

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/About.vue')
    },
    {
      path: '/callback',
      name: 'callback',
      component: () => import('./views/Home.vue'),
      beforeEnter: async (to, from, next) => {
        const code = to.query.code
        const state = to.query.state
        const codeVerifier = sessionStorage.getItem(`login-code-verifier-${state}`)
        if (!code || !codeVerifier) {
          next('/')
        }
        try {
          const res = await fetchAuthToken(code, codeVerifier)
          await store.commit('setToken', res.data.access_token)
          await setAuthToken(res.data.access_token)
          const resp = await getMe()
          await store.commit('setMe', resp.data)
          next('/')
        } catch (e) {
          console.error(e)
        }
      }
    }
  ]
})

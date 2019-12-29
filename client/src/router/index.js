import Vue from 'vue'
import Router from 'vue-router'
import store from '../store'
import About from '../components/About.vue'
import DashBoard from '../components/DashBoard.vue'
import UserPage from '../components/UserPage.vue'
import RegisterItemPage from '../components/RegisterItemPage'
import ItemDetailPage from '../components/ItemDetailPage'
import AdminPage from '../components/AdminPage'
import ItemListPage from '../components/ItemListPage'
import { fetchAuthToken, setAuthToken, getMe } from '../utils/api'

setAuthToken(store.state.authToken)

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  scrollBehavior (to, from, savedPosition) {
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        resolve({ x: 0, y: 0 })
      }, 250)
    })
  },
  routes: [
    {
      path: '/',
      name: 'Dashboard',
      component: DashBoard
    },
    {
      path: '/users',
      name: 'All Users',
      component: UserPage
    },
    {
      path: '/users/:name',
      name: 'User Page',
      component: UserPage
    },
    {
      path: '/items/new',
      name: 'Register Item Page',
      component: RegisterItemPage
    },
    {
      path: '/items/equipment',
      name: 'Equipments',
      component: ItemListPage
    },
    {
      path: '/items/property',
      name: 'Personal Property Items',
      component: ItemListPage
    },
    {
      path: '/items/:id',
      name: 'Item',
      component: ItemDetailPage
    },
    {
      path: '/items',
      name: 'All Items',
      component: ItemListPage
    },
    {
      path: '/admin',
      name: 'Admin Page',
      component: AdminPage
    },
    {
      path: '/about',
      name: 'About',
      component: About
    },
    {
      path: '/callback',
      name: 'callback',
      component: () => import('../components/Home.vue'),
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
          // eslint-disable-next-line no-console
          console.error(e)
        }
      }
    }
  ]
})

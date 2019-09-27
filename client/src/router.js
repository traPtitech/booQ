import Vue from 'vue'
import Router from 'vue-router'
import store from './store'
import Home from './components/Home.vue'
import About from './components/About.vue'
import DashBoard from './components/DashBoard.vue'
import UserPage from './components/UserPage.vue'
import RegisterItemPage from './components/RegisterItemPage'
import ItemDetailPage from './components/ItemDetailPage'
import RegisterOwnerForm from './components/RegisterOwnerForm'
import RentalForm from './components/RentalForm'
import AdminPage from './components/AdminPage'
import AllItemPage from './components/AllItemPage'
import { fetchAuthToken, setAuthToken, getMe } from './utils/api'

setAuthToken(store.state.authToken)

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
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
      path: '/items/:id',
      name: 'Item',
      component: ItemDetailPage
    },
    {
      path: '/items/:id/owner/new',
      component: RegisterOwnerForm
    },
    {
      path: '/items/:id/rental',
      component: RentalForm
    },
    {
      path: '/items',
      name: 'All Items',
      component: AllItemPage
    },
    // ここから
    {
      path: '/items_test',
      name: 'Item',
      component: ItemDetailPage
    },
    {
      path: '/register_owner_form',
      component: RegisterOwnerForm
    },
    {
      path: '/rental_form',
      component: RentalForm
    },
    // ここまで消す
    {
      path: '/admin',
      component: AdminPage
    },
    {
      // TODO: 初期ページなのである程度検証したら消す
      path: '/home',
      name: 'home',
      component: Home
    },
    {
      path: '/about',
      name: 'About',
      component: About
    },
    {
      path: '/callback',
      name: 'callback',
      component: () => import('./components/Home.vue'),
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

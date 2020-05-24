import Vue from 'vue'
import Router from 'vue-router'
import store from '../store'
import { getMe } from '../utils/api'

Vue.use(Router)

const UserPage = () => import('../components/UserPage.vue')
const ItemListPage = () => import('../components/ItemListPage')

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
      component: () => import('../components/DashBoard.vue')
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
      component: () => import('../components/RegisterItemPage')
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
      component: () => import('../components/ItemDetailPage')
    },
    {
      path: '/items',
      name: 'All Items',
      component: ItemListPage
    },
    {
      path: '/admin',
      name: 'Admin Page',
      component: () => import('../components/AdminPage')
    },
    {
      path: '/about',
      name: 'About',
      component: () => import('../components/About.vue')
    },
    {
      path: '/callback',
      name: 'callback',
      component: () => import('../components/Home.vue'),
      beforeEnter: async (to, from, next) => {
        try {
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

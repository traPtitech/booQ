import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from 'vuex-persistedstate'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    me: null,
    drawer: null,
    color: 'success',
    sidebarBackgroundColor: 'rgba(27, 27, 27, 0.74)',
    aboutDialog: false,
    authToken: null,
    cart: [],
    navBarTitle: ''
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
    toggleDrawer (state) {
      state.drawer = !state.drawer
    },
    toggleAboutDialog (state) {
      state.aboutDialog = !state.aboutDialog
    },
    item2Cart (state, data) {
      const targetItem = state.cart.filter(element => {
        return element.ID === data.ID
      })
      if (targetItem.length === 1) {
        targetItem[0].rentalCount += data.rentalCount
      } else {
        state.cart.push(data)
      }
    },
    removeItemFromCart (state, i) {
      state.cart.splice(i, 1)
    },
    resetCart (state) {
      state.cart = []
    },
    setNavBarTitle (state, data) {
      state.navBarTitle = data
    },
    resetNavBarTitle (state) {
      state.navBarTitle = ''
    }
  },
  actions: {
    //
  },
  plugins: [createPersistedState({
    paths: ['authToken']
  })]
})

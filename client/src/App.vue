<template>
  <v-app>
    <core-nav-bar />

    <core-side-bar />

    <core-view />

    <core-login-dialog />
  </v-app>
</template>

<script>
import CoreNavBar from './components/core/NavBar'
import CoreSideBar from './components/core/SideBar'
import CoreView from './components/core/View'
import CoreLoginDialog from './components/core/LoginDialog'

export default {
  name: 'App',
  components: {
    CoreNavBar,
    CoreSideBar,
    CoreView,
    CoreLoginDialog
  },
  data () {
    return {
      //
    }
  }
}
</script>

<style>
</style>

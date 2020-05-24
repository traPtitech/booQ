<template>
  <v-app>
    <core-nav-bar />

    <core-side-bar />

    <core-view />

  </v-app>
</template>

<script>
import CoreNavBar from './components/core/NavBar'
import CoreSideBar from './components/core/SideBar'
import CoreView from './components/core/View'

export default {
  name: 'App',
  components: {
    CoreNavBar,
    CoreSideBar,
    CoreView
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

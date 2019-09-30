<template>
  <v-app-bar
    id="core-app-bar"
    fixed
    app
    height="70"
    style="background: #eee;"
  >
    <v-toolbar-title
      v-if="responsive"
      class="tertiary--text align-self-center"
    >
      <v-btn
        dark
        icon
        @click.stop="onClickBtn"
      >
        <v-avatar>
          <v-img :src="logo" />
        </v-avatar>
      </v-btn>
      booQ
    </v-toolbar-title>
    <v-toolbar-title
      v-else
      class="tertiary--text align-self-center"
    >
      {{ title }}
    </v-toolbar-title>
    <v-spacer />
    <v-toolbar-items>
      <v-flex
        layout
        py-2
      >
        <v-btn
          v-if="$store.state.me"
          dark
          icon
          :to="`/users/${$store.state.me.name}`"
        >
          <v-avatar size="40">
            <img :src="`https://q.trap.jp/api/1.0/public/icon/${$store.state.me.name}`">
          </v-avatar>
        </v-btn>
        <router-link
          v-else
          v-ripple
          to="/about"
        >
          <v-icon color="tertiary">mdi-account</v-icon>
        </router-link>
      </v-flex>
    </v-toolbar-items>
  </v-app-bar>
</template>

<script>
import { mapMutations } from 'vuex'
import { getMe } from '@/utils/api'

export default {
  data () {
    return {
      logo: '/img/logo.png',
      title: null,
      responsive: false,
      responsiveInput: false
    }
  },
  watch: {
    '$route' (val) {
      this.title = val.name
    }
  },
  async created () {
    try {
      if (!this.$store.state.me) {
        if (location.pathname === '/callback') return
        const resp = await getMe()
        await this.$store.commit('setMe', resp.data)
      }
    } catch (e) {
      this.toggleLoginDialog()
    }
  },
  mounted () {
    this.onResponsiveInverted()
    window.addEventListener('resize', this.onResponsiveInverted)
    this.title = this.$route.name
  },
  beforeDestroy () {
    window.removeEventListener('resize', this.onResponsiveInverted)
  },
  methods: {
    ...mapMutations(['setDrawer', 'toggleDrawer', 'toggleLoginDialog']),
    onClickBtn () {
      this.setDrawer(!this.$store.state.drawer)
    },
    onClick () {
      //
    },
    onResponsiveInverted () {
      if (window.innerWidth < 991) {
        this.responsive = true
        this.responsiveInput = false
      } else {
        this.responsive = false
        this.responsiveInput = true
      }
    }
  }
}
</script>

<style>
  #core-app-bar {
    width: auto;
  }
  #core-app-bar a {
    text-decoration: none;
  }
</style>

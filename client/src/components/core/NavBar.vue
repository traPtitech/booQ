<template>
  <v-toolbar
    id="core-toolbar"
    fixed
    prominent
    style="background: #eee;"
  >
    <div class="v-toolbar-title">
      <v-toolbar-title
        v-if="responsive"
        class="tertiary--text title"
      >
        <v-btn
          class="default v-btn--simple"
          dark
          icon
          @click.stop="onClickBtn"
        >
          <v-img
            :src="logo"
            height="35"
          />
        </v-btn>
        booQ
      </v-toolbar-title>
      <v-toolbar-title
        v-else
        class="tertiary--text title"
        style="padding-left: 260px;"
      >
        {{ title }}
      </v-toolbar-title>
    </div>

    <v-spacer />
    <v-toolbar-items>
      <v-flex
        align-center
        layout
        py-2
      >
        <router-link
          v-ripple
          class="toolbar-items"
          to="/"
        >
          <v-icon color="tertiary">mdi-view-dashboard</v-icon>
        </router-link>
        <v-btn
          v-if="$store.state.me"
          class="default v-btn--simple"
          dark
          icon
          @click="$router.push('/about')"
        >
          <v-avatar size="40">
            <img :src="`https://q.trap.jp/api/1.0/public/icon/${$store.state.me.name}`">
          </v-avatar>
        </v-btn>
        <router-link
          v-else
          v-ripple
          class="toolbar-items"
          to="/about"
        >
          <v-icon color="tertiary">mdi-account</v-icon>
        </router-link>
      </v-flex>
    </v-toolbar-items>
  </v-toolbar>
</template>

<script>
import { mapMutations } from 'vuex'
import { getMe } from '@/utils/api'

export default {
  data () {
    return {
      logo: './img/logo.png',
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
  #core-toolbar a {
    text-decoration: none;
  }
</style>

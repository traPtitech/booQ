<template>
  <v-navigation-drawer
    id="app-drawer"
    v-model="inputValue"
    app
    dark
    floating
    persistent
    mobile-break-point="991"
    width="260"
  >
    <v-layout
      class="fill-height"
      column
    >
      <v-list
        style="display: contents;"
      >
        <v-list-tile avatar>
          <v-list-tile-avatar>
            <v-img
              :src="logo"
              height="45"
            />
          </v-list-tile-avatar>
          <v-list-tile-title class="title">
            booQ
          </v-list-tile-title>
        </v-list-tile>
        <v-divider/>
        <v-list-tile>
          <v-text-field
            class="purple-input search-input"
            label="Search..."
            color="purple"
          />
        </v-list-tile>
        <v-list-tile
          v-for="(link, i) in links"
          :key="i"
          :to="link.to"
          :active-class="color"
          avatar
          class="v-list-item"
        >
          <v-list-tile-action>
            <v-icon>{{ link.icon }}</v-icon>
          </v-list-tile-action>
          <v-list-tile-title
            v-text="link.text"
          />
        </v-list-tile>
        <v-list-tile
          active-class="primary"
          class="nav-footer"
        >
          <div class="font-weight-light body-1">
            <a href="https://github.com/traPtitech/booQ" target="_blank">booQ Project</a>
          </div>
        </v-list-tile>
      </v-list>
    </v-layout>
  </v-navigation-drawer>
</template>

<script>
// Utilities
import { mapMutations, mapState } from 'vuex'
export default {
  name: 'SideBar',
  data () {
    return {
      logo: './img/logo.png',
      links: [
        {
          to: '/',
          icon: 'mdi-view-dashboard',
          text: 'Dashboard'
        },
        {
          to: '/items',
          icon: 'mdi-view-list',
          text: 'Item List'
        },
        {
          to: `/user/wip`,
          icon: 'mdi-account',
          text: 'User Profile'
        },
        {
          to: `/items/new`,
          // icon: 'mdi-account', あとで
          text: ' Register Item'
        }
      ],
      responsive: false
    }
  },
  computed: {
    ...mapState(['color']),
    inputValue: {
      get () {
        return this.$store.state.drawer
      },
      set (val) {
        this.setDrawer(val)
      }
    },
    items () {
      return this.$t('Layout.View.items')
    },
    sidebarOverlayGradiant () {
      return `${this.$store.state.sidebarBackgroundColor}, ${this.$store.state.sidebarBackgroundColor}`
    }
  },
  mounted () {
    this.onResponsiveInverted()
    window.addEventListener('resize', this.onResponsiveInverted)
  },
  beforeDestroy () {
    window.removeEventListener('resize', this.onResponsiveInverted)
  },
  methods: {
    ...mapMutations(['setDrawer', 'toggleDrawer']),
    onResponsiveInverted () {
      if (window.innerWidth < 991) {
        this.responsive = true
      } else {
        this.responsive = false
      }
    }
  }
}
</script>

<style lang="scss">
  #app-drawer {
    .v-list__tile {
      border-radius: 4px;
    }
    .nav-footer {
      margin-top: auto;
      margin-bottom: 5px;
    }
    .v-image__image--contain {
      top: 9px;
      height: 60%;
    }
    .search-input {
      margin-bottom: 30px !important;
      padding-left: 15px;
      padding-right: 15px;
    }
    div.v-responsive.v-image > div.v-responsive__content {
      overflow-y: auto;
    }
  }
</style>

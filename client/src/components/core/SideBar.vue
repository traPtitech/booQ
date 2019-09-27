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
    <v-list-item two-line>
      <v-list-item-avatar>
        <v-img
          :src="logo"
          height="45"
        />
      </v-list-item-avatar>
      <v-list-item-title class="title">
        booQ
      </v-list-item-title>
    </v-list-item>
    <v-divider/>
    <v-list nav>
      <v-list-item>
        <v-text-field
          class="purple-input search-input"
          label="Search..."
          color="purple"
        />
      </v-list-item>
      <v-list-item
        v-for="(link, i) in links"
        :key="i"
        :to="link.to"
        :active-class="color"
        avatar
        class="v-list-item"
      >
        <v-list-item-action>
          <v-icon>{{ link.icon }}</v-icon>
        </v-list-item-action>
        <v-list-item-title
          v-text="link.text"
        />
      </v-list-item>
      <v-list-item
        active-class="primary"
        class="nav-footer"
      >
        <div class="font-weight-light body-1">
          
        </div>
      </v-list-item>
    </v-list>
    <template v-slot:append>
      <v-list nav>
        <v-list-item
          href="https://github.com/traPtitech/booQ"
        >
          <v-list-item-action>
            <v-icon>mdi-book</v-icon>
          </v-list-item-action>

          <v-list-item-title class="font-weight-light">
            booQ Project
          </v-list-item-title>
        </v-list-item>
      </v-list>
    </template>
  </v-navigation-drawer>
</template>

<script>
// Utilities
import { mapMutations, mapState } from 'vuex'
export default {
  name: 'SideBar',
  data () {
    return {
      logo: '/img/logo.png',
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

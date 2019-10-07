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
        <v-menu v-model="value">
          <template v-slot:activator="{ on }">
            <v-btn dark icon v-on="on">
             <v-icon>mdi-cart</v-icon>
            </v-btn>
          </template>
          <v-list>
            <v-list-item v-for="(item, i) in $store.state.cart" :key="i">
              <v-list-item-title>
                {{ item.name }}
              </v-list-item-title>
              <v-list-item-action>
                <v-btn icon @click="items.splice( i, 1)">
                  <v-icon>mdi-minus-circle</v-icon>
                </v-btn>
              </v-list-item-action>
            </v-list-item>
          </v-list>
        </v-menu>
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
import axios from 'axios'

export default {
  data () {
    return {
      logo: '/img/logo.png',
      title: null,
      responsive: false,
      responsiveInput: false,
      cartPurpose: '',
      cartDueDate: '' ,
      cartError: null
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
    },
     async rentCartItem () {
       for ( var i = 0; i < this.$store.state.cart.length;  i++ ) {
         var names = []
         names = names.push(this.$store.state.cart[i].name)
         await axios.post(`/api/items/` + this.$store.state.cart[i].id + `/logs`, { owner_id: 1, type: 0, purpose: this.cartPurpose, due_date: this.cartDueDate, count: this.$store.state.cart[i].rentalCount })
           .catch(e => {
             alert(e)
             this.cartError = e
           })
       }
       if (!this.cartError) { alert(names) }
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

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
        style="width: 30px;"
      >
        <v-avatar>
          <img
            :src="logo"
            style="width: 30px;"
          />
        </v-avatar>
      </v-btn>
      <img
        :src="logoTitle"
        class="logo-title"
      />
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
        <v-menu min-width="250" max-width="400">
          <template v-slot:activator="{ on }">
            <v-btn :disabled="$store.state.cart.length == 0" icon v-on="on">
              <mdi-icon name="mdi-cart" dark />
            </v-btn>
          </template>
          <v-list>
            <v-list-item v-for="(item, i) in $store.state.cart" :key="i">
              <v-list-item-title>
                {{ item.name }} × {{ item.rentalCount }}
              </v-list-item-title>
              <v-list-item-action>
                <v-btn icon @click="$store.commit('removeItemFromCart', i)">
                  <mdi-icon name="mdi-minus-circle" dark />
                </v-btn>
              </v-list-item-action>
            </v-list-item>
            <v-list-item>
              <v-list-item-action>
                <v-btn @click.stop="cartDialog = !cartDialog" primary>まとめて借りる</v-btn>
                <div class="text-center">
                  <v-dialog max-width="320" v-model="cartDialog">
                    <v-card width="320">
                      <v-card-title class="headline">備品をまとめて借りる</v-card-title>
                      <v-card-actions>
                        <div>
                          <v-form ref="form">
                            <v-textarea outlined v-model="cartPurpose" :rules="[() => !!cartPurpose || 'This field is required']" label="目的"/>
                          </v-form>
                        </div>
                      </v-card-actions>
                      <v-card-actions max-width="320">
                        <v-date-picker v-model="cartDueDate"></v-date-picker>
                      </v-card-actions>
                      <v-divider></v-divider>
                      <v-card-actions>
                        <v-btn v-on:click="rentCartItem()">借りる</v-btn>
                      </v-card-actions>
                    </v-card>
                  </v-dialog>
                </div>
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
            <img :src="`https://q.trap.jp/api/v3/public/icon/${$store.state.me.name}`">
          </v-avatar>
        </v-btn>
        <router-link
          v-else
          v-ripple
          to="/about"
        >
          <mdi-icon color="tertiary" name="mdi-account" />
        </router-link>
      </v-flex>
    </v-toolbar-items>
  </v-app-bar>
</template>

<script>
import { mapMutations } from 'vuex'
import { getMe } from '@/utils/api'
import axios from 'axios'
import MdiIcon from '../shared/MdiIcon.vue'

export default {
  components: {
    MdiIcon
  },
  data () {
    return {
      logo: '/img/logo-main.svg',
      logoTitle: '/img/logo-type.svg',
      title: null,
      responsive: false,
      responsiveInput: false,
      cartPurpose: '',
      cartDueDate: '',
      cartDialog: false,
      error: null
    }
  },
  watch: {
    '$route' (val) {
      this.title = val.name
    },
    '$store.state.navBarTitle' (val) {
      if (this.$store.state.navBarTitle) {
        this.title = val
      }
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
    ...mapMutations(['setDrawer', 'toggleDrawer']),
    onClickBtn () {
      this.setDrawer(!this.$store.state.drawer)
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
      if (!this.cartPurpose) {
        alert('目的を入力してください')
        return
      }
      if (!this.cartDueDate) {
        alert('返却日を入力してください')
        return
      }
      for (let i = 0; i < this.$store.state.cart.length; i++) {
        let names = []
        names = names.push(this.$store.state.cart[i].name)
        await axios.post(
          '/api/items/' + this.$store.state.cart[i].id + '/logs',
          {
            ownerId: 1,
            type: 0,
            purpose: this.cartPurpose,
            dueDate: this.cartDueDate,
            count: this.$store.state.cart[i].rentalCount
          }
        ).catch(e => {
          this.error = e
          alert(e)
        })
      }
      if (!this.error) {
        this.cartDialog = !this.cartDialog
        alert('まとめて借りることに成功しました。')
        this.$store.commit('resetCart')
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
  .logo-title {
    position: absolute;
    padding-left: 10px;
    padding-top: 15px;
    width: 60px;
  }
</style>

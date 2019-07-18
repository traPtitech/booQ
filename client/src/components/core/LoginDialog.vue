<template>
  <v-layout
    row
    justify-center
  >
    <v-dialog
      v-model="loginDialog"
      persistent
      max-width="290"
    >
      <v-card>
        <v-card-title
          class="headline"
        >
          ログインしてください
        </v-card-title>
        <v-card-text>
          OKを押すとtraQに飛びます。<br>
          「承認」を押すとログインできます。
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="green darken-1" @click="login">OK</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-layout>
</template>

<script>
import {
  mapMutations,
  mapState
} from 'vuex'
import { redirectAuthorizationEndpoint } from '@/utils/api'
export default {
  computed: {
    ...mapState(['loginDialog'])
  },
  methods: {
    ...mapMutations(['toggleLoginDialog']),
    login () {
      this.toggleLoginDialog()
      redirectAuthorizationEndpoint()
    }
  }
}
</script>

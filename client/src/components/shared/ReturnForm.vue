<template>
  <div>
    <v-btn block color="warning" @click.stop="open" :disabled="data.rental_users.filter(function (element) {return element.user_id = $store.state.me.ID}).length === 0">返却する</v-btn>
    <div class="text-center">
      <v-dialog light v-model="isOpenReturnForm" max-width="320">
        <v-card width="320">
          <v-card-title class="headline">物品を返却する</v-card-title>
          <v-card-actions>
            <v-menu bottom origin="center center" transition="scale-transition" open-on-hover>
              <template v-slot:activator="{ on }">
                <v-btn color="primary" dark v-on="on">返却する所有者を選ぶ</v-btn>
              </template>
              <v-list>
                <v-list-item
                v-for="(rentalUser, i) in data.rental_users.filter(function (element) {return element.user_id = $store.state.me.ID})"
                :key="i"
                @click="returnOwnerID = rentalUser.owner.ID">
                  <v-list-item-title>{{ rentalUser.owner.name }}</v-list-item-title>
                </v-list-item>
              </v-list>
            </v-menu>
          </v-card-actions>
          <div>{{returnOwnerID}}{{getRentalCount(returnOwnerID)}}</div>
          <v-card-actions v-if="getRentalCount(returnOwnerID) > 1">
            <v-slider :max="getRentalCount(returnOwnerID)" v-model="data.returnCount" thumb-label="always" />
          </v-card-actions>
          <v-divider></v-divider>
          <v-card-actions>
            <div class="flex-grow-1"></div>
            <v-btn v-on:click="returnItem()">返却する</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'ReturnForm',
  props: {
    data: Object
  },
  data () {
    return {
      returnOwnerID: 0,
      returnCount: 0,
      error: '',
      isOpenReturnForm: false
    }
  },
  methods: {
    getRentalCount (ownerID) {
      if (this.data.rental_users.length === 0) {
        return '借りていません'
      }
      let rentalUsers = this.data.rental_users.filter(element => {
        return (element.user.ID = this.$store.state.me.ID)
      })
      const rentalUser = rentalUsers.find(function (element) {
        return (element.owner_id = ownerID)
      })
      if (!rentalUser) {
        return '借りていません'
      }
      return rentalUsers[0].count * -1
    },
    async returnItem () {
      if (this.returnOwnerID === 0) {
        alert('所有者を選択してください')
        return false
      }
      const today = new Date()
      await axios.post(`/api/items/` + this.$route.params.id + `/logs`, { owner_id: this.returnOwnerID, type: 1, count: this.returnCount, purpose: '', due_date: today.getFullYear() + '-' + ('00' + (today.getMonth() + 1)).slice(-2) + '-' + ('00' + today.getDate()).slice(-2) })
        .catch(e => {
          alert(e)
          this.error = e
        })
      if (!this.error) { alert('あなたは”' + this.data.ID + '”を返しました。') }
      this.isOpenReturnForm = !this.isOpenReturnForm
    },
    open () {
      this.isOpenReturnForm = !this.isOpenReturnForm
    }
  }
}
</script>

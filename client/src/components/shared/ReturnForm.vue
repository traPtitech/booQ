<template>
  <div>
    <v-btn block color="warning" @click.stop="open" :disabled="propItem.rental_users.filter(element => {return element.userId === $store.state.me.id}).length === 0">返却する</v-btn>
    <div class="text-center">
      <v-dialog light v-model="isOpenReturnForm" max-width="320">
        <v-card width="320">
          <v-card-title class="headline">物品を返却する</v-card-title>
          <v-card-actions>
            <v-menu bottom origin="center center" transition="scale-transition">
              <template v-slot:activator="{ on }">
                <v-btn color="primary" dark v-on="on">返却する所有者を選ぶ</v-btn>
              </template>
              <v-list>
                <v-list-item
                v-for="(rentalUser, i) in propItem.rental_users.filter(function (element) {return element.userId === $store.state.me.id})"
                :key="i"
                @click="returnOwnerID = rentalUser.owner.id; returnOwnerName = rentalUser.owner.name">
                  <v-list-item-title>{{ rentalUser.owner.name }}</v-list-item-title>
                </v-list-item>
              </v-list>
            </v-menu>
              <div>
                  {{returnOwnerName}}
              </div>
          </v-card-actions>
          <v-card-actions v-if="getRentalCount(returnOwnerID) > 1">
            <v-slider :max="getRentalCount(returnOwnerID)" min="1" v-model="returnCount" thumb-label="always" />
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
    propItem: Object
  },
  data () {
    return {
      returnOwnerID: 0,
      returnOwnerName: '',
      returnCount: 1,
      error: '',
      isOpenReturnForm: false
    }
  },
  methods: {
    getRentalCount (ownerID) {
      if (this.propItem.rental_users.length === 0) {
        return 0
      }
      const rentalUsers = this.propItem.rental_users.filter(element => {
        return element.user.id === this.$store.state.me.id
      })
      const rentalUser = rentalUsers.find(element => {
        return element.owner_id === ownerID
      })
      if (!rentalUser) {
        return 0
      }
      return rentalUser.count * -1
    },
    async returnItem () {
      this.error = null
      if (this.returnOwnerID === 0) {
        alert('所有者を選択してください')
        return
      }
      const today = new Date()
      await axios.post('/api/items/' + this.$route.params.id + '/logs', { owner_id: this.returnOwnerID, type: 1, count: this.returnCount, purpose: '', due_date: today.getFullYear() + '-' + ('0' + (today.getMonth() + 1)).slice(-2) + '-' + ('0' + today.getDate()).slice(-2) })
        .catch(e => {
          alert(e)
          this.error = e
          return false
        })
      if (this.error) { return false }
      if (!this.error) { alert('あなたは”' + this.propItem.name + '”を返しました。') }
      this.isOpenReturnForm = !this.isOpenReturnForm
      this.$emit('reload')
    },
    open () {
      this.isOpenReturnForm = !this.isOpenReturnForm
    }
  }
}
</script>

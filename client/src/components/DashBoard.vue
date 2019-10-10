<template>
  <div>
    <h2>あなたが借りている物品</h2>
    <v-card
      class="mx-auto"
      width="1500"
      elevation="5"
      tile
    >
      <v-list
        two-line
        avatar
        nav
      >
        <v-list-item-group color="primary">
          <v-list-item
            v-for="item in items"
            :key="item.id"
            @click.stop="$router.push({path: `/items/${item.ID}`})"
            style="height: 100px;"
          >
            <img
              :src="item.img_url.length ? item.img_url : 'https://q.trap.jp/api/1.0/files/3380fbc6-6141-4b60-99ae-a1d270842d60/thumbnail'"
              class="item-list-image"
            />
            <v-list-item-content style="padding-left: 15px;" :to="`/items/${item.ID}`">
              <v-list-item-title class="headline mb-1">{{ item.name }}</v-list-item-title>
              <v-list-item-subtitle>{{ item.owners.map(i => i.user.name).join(', ') }}</v-list-item-subtitle>
            </v-list-item-content>
            <v-list-item-action class="item-list-icons">
              <v-btn :disabled="getBihinMyRentalCount(item.ID) === 1" icon v-if="item.type !== 0" @click.stop="click2Cart(item)">
                <v-icon>mdi-cart-arrow-down</v-icon>
              </v-btn>
            </v-list-item-action>
          </v-list-item>
        </v-list-item-group>
      </v-list>
    </v-card>
    <div v-if="returnCart.length === 0">
      まとめて返却する物品
    </div>
    <div>
      <v-list-item v-for="itemInCart in returnCart" :key="itemInCart.ID">
        <div>{{itemInCart.name}} × {{itemInCart.returnCount}}</div>
      </v-list-item>
    </div>
    <v-btn @click="returnItems">まとめて返却</v-btn>
    <div class="text-center">
      <v-dialog
        v-if="isOpen2Cart"
        v-model="isOpen2Cart"
        max-width="290"
      >
        <v-card width="290">
          <v-card-title class="headline grey lighten-2" primary-title>
            個数を選択
          </v-card-title>
          <v-card-actions>
            <v-slider :max="maxCount" v-model="itemCount" thumb-label="always" />
          </v-card-actions>
          <v-divider />
          <v-card-actions>
            <v-btn @click="putItem2Cart" primary>
              Put Cart
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DashBoard',
  data () {
    return {
      isOpen2Cart: false,
      returnCart: [],
      itemCount: 0,
      maxCount: 0,
      item: {},
      items: []
    }
  },
  mounted () {
    axios
      .get(`/api/items/` + this.$store.state.me.name)
      .then(res => (this.items = res.data))
      .catch(e => { alert(e) })
  },
  methods: {
    getBihinMyRentalCount (itemID) {
      const item = this.items.find(element => {
        return element.ID === itemID
      })
      if (!item) {
        alert('対象itemがありません')
        return 0
      }
      let targetRentalUser = item.rental_users.find(rentalUser => {
        return rentalUser.count < 0
      })
      if (!targetRentalUser) {
        return 0
      }
      return targetRentalUser.count * -1
    },
    putItem2Cart () {
      if (this.itemCount <= 0) {
        alert('0以上を入力してください')
        return
      }
      this.item.returnCount = this.itemCount
      this.returnCart.push()
      this.itemCount = 0
      this.maxCount = 0
      this.item = {}
      this.isOpen2Cart = !this.isOpen2Cart
    },
    click2Cart (item) {
      this.item = item
      this.maxCount = this.getBihinMyRentalCount(item.ID)
      this.isOpen2Cart = !this.isOpen2Cart
    },
    async rentCartItem () {
      for (let i = 0; i < this.returnCart.length; i++) {
        let names = []
        names = names.push(this.returnCart[i].name)
        await axios.post(`/api/items/` + this.returnCart[i].ID + `/logs`, { owner_id: 1, type: 1, count: this.returnCart[i].returnCount, purpose: '', due_date: today.getFullYear() + '-' + ('00' + (today.getMonth() + 1)).slice(-2) + '-' + ('00' + today.getDate()).slice(-2) })
          .catch(e => {
            this.error = e
            alert(e)
          })
      }
      if (!this.error) {
        this.cartDialog = !this.cartDialog
        alert('まとめて返すことに成功しました。')
        let message = '入'
        console.log(this.returnCart)
        console.log(this.returnCart.length)
        for (let i = 0; i < this.returnCart.length; i++) {
          message = message + '\n[' + this.returnCart[i].name + '](' + process.env.VUE_APP_API_ENDPOINT + '/items/' + this.returnCart[i].ID + ') × ' + this.returnCart[i].returnCount
        }
        await axios.post(`${traQBaseURL}/channels/` + process.env.VUE_APP_EQUIPMENT_CHANNEL_ID + `/messages?embed=1`, { text: message }).catch(e => { alert(e) })
        this.returnCart = []
      }
    }
  }
}
</script>

<style>
.item-list-image {
  max-height: 100px;
  max-width: 120px;
}
.item-list-icons {
  max-height: 80px;
  max-width: 80px;
}
</style>

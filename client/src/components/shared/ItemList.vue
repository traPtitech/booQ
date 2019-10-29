<template>
  <div>
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
              :src="item.img_url.length ? item.img_url : '/img/no-image.svg'"
              class="item-list-image"
            />
            <v-list-item-content style="padding-left: 15px;" :to="`/items/${item.ID}`">
              <v-list-item-title class="headline mb-1">{{ item.name }}</v-list-item-title>
              <v-list-item-subtitle>{{ item.owners.map(i => i.user.name).join(', ') }}</v-list-item-subtitle>
            </v-list-item-content>
            <v-list-item-action class="item-list-icons">
              <div v-if="item.like_counts > 0">
                <v-icon>thumb_up_alt</v-icon>
                {{ item.like_counts }}
              </div>
              <v-btn :disabled="getBihinLatestCount(item.ID) < 1" icon v-if="item.type !== 0" @click.stop="click2Cart(item)">
                <v-icon>mdi-cart-arrow-down</v-icon>
              </v-btn>
            </v-list-item-action>
          </v-list-item>
        </v-list-item-group>
      </v-list>
    </v-card>
    <div class="text-center">
      <v-dialog
        v-if="isOpen2Cart"
        v-model="isOpen2Cart"
        max-width="290"
      >
        <v-card width="270">
          <v-card-title class="headline grey lighten-2" primary-title>
            個数を選択
          </v-card-title>
          <v-card-actions>
            <v-slider :max="maxCount" min="1" v-model="itemCount" thumb-label="always" />
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
  name: 'ItemList',
  props: [
    'items'
  ],
  data () {
    return {
      isOpen2Cart: false,
      itemCount: 1,
      maxCount: 0,
      item: {}
    }
  },
  methods: {
    getBihinLatestCount (itemID) {
      const item = this.items.find(element => {
        return element.ID === itemID
      })
      if (!item) {
        alert('対象itemがありません')
        return 0
      }
      let targetLog = item.latest_logs.find(log => {
        return log.owner.name === 'traP' || log.owner.name === 'sienka'
      })
      if (!targetLog) {
        // logが存在しない場合
        const targetOwner = item.owners.find(owner => {
          return owner.owner.name === 'traP' || owner.owner.name === 'sienka'
        })
        if (!targetOwner) {
          // OwnerにtraPがいない状態
          return 0
        }
        return targetOwner.count
      }
      return targetLog.count
    },
    putItem2Cart () {
      if (this.itemCount <= 0) {
        alert('0以上を入力してください')
        return
      }
      this.item.rentalCount = this.itemCount
      this.$store.commit('item2Cart', this.item)
      this.itemCount = 1
      this.maxCount = 0
      this.item = {}
      this.isOpen2Cart = !this.isOpen2Cart
    },
    click2Cart (item) {
      this.item = item
      this.maxCount = this.getBihinLatestCount(item.ID) - this.searchItemCountInCart(item)
      this.isOpen2Cart = !this.isOpen2Cart
    },
    searchItemCountInCart (item) {
      const targetItem = this.$store.state.cart.find(element => {
        return element.ID === item.ID
      })
      if (targetItem) {
        return targetItem.rentalCount
      }
      return 0
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

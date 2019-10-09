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
        <v-list-item-group v-model="items" color="primary">
          <v-list-item
            v-for="item in items"
            :key="item.ID"
            style="height: 100px;"
          >
            <v-list-item :to="`/items/${item.ID}`">
              <img
                :src="item.img_url === '' ? 'https://q.trap.jp/api/1.0/files/3380fbc6-6141-4b60-99ae-a1d270842d60/thumbnail' : item.img_url"
                class="item-list-image"
              />
              <v-list-item-content style="padding-left: 15px;" :to="`/items/${item.ID}`">
                <v-list-item-title class="headline mb-1">{{ item.name }}</v-list-item-title>
                <v-list-item-subtitle>{{ item.owners.map(i => i.owner.name).join(', ') }}</v-list-item-subtitle>
              </v-list-item-content>
              <v-list-item-action class="item-list-icons">
                <v-btn icon v-if="item.type === 1" @click.stop="click2Cart()" absolute>
                  <v-icon>mdi-cart-arrow-down</v-icon>
                </v-btn>
                <div class="text-center">
                  <v-dialog v-model="isOpen2Cart" max-width="290">
                    <v-card width="290">
                      <v-card-title class="headline grey lighten-2" primary-title>
                        個数を選択
                      </v-card-title>
                      <v-card-actions v-if="item.type === 1">
                        <v-slider :max="getBihinLatestCount(item.ID)" v-model="item.rentalCount" thumb-label="always" />
                      </v-card-actions>
                      <v-divider></v-divider>
                      <v-card-actions>
                        <v-btn @click="putItem2Cart(item)" primary>
                          Put Cart
                        </v-btn>
                      </v-card-actions>
                    </v-card>
                  </v-dialog>
                </div>
              </v-list-item-action>
              <!-- TODO: まだ/itemsのレスポンスにlike_countがないので保留 -->
              <!-- <v-list-item-action>
                <v-icon>thumb_up_alt</v-icon>
                {{ item.like_counts }}
              </v-list-item-action> -->
            </v-list-item>
          </v-list-item>
        </v-list-item-group>
      </v-list>
    </v-card>
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
      isOpen2Cart: false
    }
  },
  methods: {
    getBihinLatestCount (itemID) {
      var item = this.items.filter(function (element) {
        return (element.ID = itemID)
      })
      var targetLog = item[0].latest_logs.filter(function (log) {
        return (log.owner.name = 'trap')
      })
      if (targetLog === []) {
        targetLog = targetLog.push(item[0].owners.filter(function (owner) {
          return (owner.owner.name = 'trap')
        }))
      }
      if (targetLog.length === 0) return 0
      return targetLog[0].count
    },
    putItem2Cart (item) {
      this.$store.commit('item2Cart', item)
      this.isOpen2Cart = !this.isOpen2Cart
    },
    click2Cart () {
      this.isOpen2Cart = !this.isOpen2Cart
      return false
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

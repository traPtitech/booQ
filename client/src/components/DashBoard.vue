<template>
  <div>
    <div>
      使い方ページは
      <router-link :to="`/about`">こちら</router-link>
    </div>
    <div v-if="items !== null">
      <div v-if="items.length === 0">
        <h2>あなたが現在借りている物品はありません</h2>
      </div>
      <div v-else style="padding-bottom: 30px;">
        <div>
          <div
            v-for="item in items"
            :key="item.ID"
          >
            <v-alert type="error" v-if="checkDueDate(item)">
              {{ item.name }}が未返却です
            </v-alert>
          </div>
        </div>
        <div>
          <v-row class="fill-height">
            <v-col>
              <v-sheet height="64">
                <v-toolbar flat color="white">
                  <v-btn fab text small @click="prev">
                    <v-icon small>mdi-chevron-left</v-icon>
                  </v-btn>
                  <v-btn fab text small @click="next">
                    <v-icon small>mdi-chevron-right</v-icon>
                  </v-btn>
                  <v-toolbar-title>あなたが借りている物品</v-toolbar-title>
                  <v-spacer />
                  <v-btn @click="returnItems" :disabled="!returnCart.length">まとめて返却</v-btn>
                </v-toolbar>
              </v-sheet>
              <v-sheet height="600">
                <v-calendar
                  ref="calendar"
                  v-model="focus"
                  color="primary"
                  :type="type"
                  :events="items"
                  :event-color="getItemColor"
                  :event-margin-bottom="3"
                  @click:event="showEvent"
                  @click:more="viewDay"
                  @change="updateRange"
                />
                <v-menu
                  v-model="selectedOpen"
                  :close-on-content-click="false"
                  :activator="selectedElement"
                  full-width
                  offset-x
                >
                  <v-card
                    color="grey lighten-4"
                    min-width="350px"
                    max-width="600px"
                    flat
                  >
                    <v-toolbar
                      :color="selectedItem.type === 0 ? 'green' : 'grey darken-1'"
                      dark
                    >
                      <v-toolbar-title>{{ selectedItem.name }}</v-toolbar-title>
                    </v-toolbar>
                    <br>
                    <div class="text-center">
                      <v-img v-if="selectedItem.img_url" contain :src="selectedItem.img_url.length ? selectedItem.img_url : '/img/no-image.svg'" height="194" />
                    </div>
                    <v-card-text>
                      <span>{{ selectedItem.description }}</span>
                    </v-card-text>
                    <v-card-actions>
                      <v-btn
                        text
                        color="secondary"
                        @click="click2Cart(selectedItem)"
                        outlined
                        v-if="selectedItem.type != 0"
                      >
                        返却するものにまとめる
                      </v-btn>
                      <v-btn
                        block
                        text
                        color="secondary"
                        @click.stop="$router.push({path: `/items/${selectedItem.ID}`})"
                        outlined
                        v-else
                      >
                        詳細
                      </v-btn>
                    </v-card-actions>
                  </v-card>
                </v-menu>
              </v-sheet>
            </v-col>
          </v-row>
        </div>
      </div>
      <div v-if="returnCart.length">
        <div>まとめて返却する備品</div>
        <div>
          <v-list-item v-for="itemInCart in returnCart" :key="itemInCart.ID">
            <div>{{itemInCart.name}} × {{itemInCart.returnCount}}</div>
          </v-list-item>
        </div>
        <v-btn @click="returnItems">まとめて返却</v-btn>
      </div>
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
              <v-slider :max="maxCount" min="1" v-model="itemCount" thumb-label="always" />
            </v-card-actions>
            <v-divider />
            <v-card-actions>
              <v-btn @click="putItem2Cart" primary>
                返す物品としてまとめる
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </div>
    </div>
    <div v-else>
      読み込み中...
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import { traQBaseURL, getMe } from '../utils/api.js'

export default {
  name: 'DashBoard',
  data () {
    return {
      isOpen2Cart: false,
      returnCart: [],
      itemCount: 0,
      maxCount: 0,
      item: {},
      items: null,
      error: null,
      // 以下カレンダー
      type: 'month',
      today: null,
      focus: null,
      start: null,
      end: null,
      selectedItem: {},
      selectedElement: null,
      selectedOpen: false
    }
  },
  computed: {
    title () {
      const { start, end } = this
      if (!start || !end) {
        return ''
      }

      const startMonth = this.monthFormatter(start)
      const endMonth = this.monthFormatter(end)
      const suffixMonth = startMonth === endMonth ? '' : endMonth

      const startYear = start.year
      const endYear = end.year
      const suffixYear = startYear === endYear ? '' : endYear

      const startDay = start.day + this.nth(start.day)
      const endDay = end.day + this.nth(end.day)

      switch (this.type) {
        case 'month':
          return `${startMonth} ${startYear}`
        case 'week':
        case '4day':
          return `${startMonth} ${startDay} ${startYear} - ${suffixMonth} ${endDay} ${suffixYear}`
        case 'day':
          return `${startMonth} ${startDay} ${startYear}`
      }
      return ''
    },
    monthFormatter () {
      return this.$refs.calendar.getFormatter({
        timeZone: 'GMT', month: 'long'
      })
    }
  },
  mounted () {
    this.mount()
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
      const targetRentalUser = item.rental_users.find(rentalUser => {
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
      const targetItem = this.returnCart.filter(element => {
        return element.ID === this.item.ID
      })
      if (targetItem.length === 1) {
        targetItem[0].returnCount += this.item.returnCount
      } else {
        this.returnCart.push(this.item)
      }
      this.itemCount = 0
      this.maxCount = 0
      this.item = {}
      this.isOpen2Cart = !this.isOpen2Cart
    },
    click2Cart (item) {
      this.item = item
      this.maxCount = this.getBihinMyRentalCount(item.ID) - this.searchItemCountInCart(item)
      this.isOpen2Cart = !this.isOpen2Cart
    },
    searchItemCountInCart (item) {
      const targetItem = this.returnCart.find(element => {
        return element.ID === item.ID
      })
      if (targetItem) {
        return targetItem.returnCount
      } else {
        return 0
      }
    },
    async returnItems () {
      if (!this.returnCart) return false
      for (let i = 0; i < this.returnCart.length; i++) {
        let names = []
        const myLatest = this.returnCart[i].latest_logs.find(element => {
          return element.user_id === this.$store.state.me.ID
        })
        const dueDate = myLatest.due_date
        names = names.push(this.returnCart[i].name)
        const ownerID = this.returnCart[i].rental_users.length === 1 ? this.returnCart[i].rental_users[0].owner_id : 1
        await axios.post('/api/items/' + this.returnCart[i].ID + '/logs', { owner_id: ownerID, type: 1, count: this.returnCart[i].returnCount, purpose: '', due_date: dueDate.substr(0, 10) })
          .catch(e => {
            this.error = e
            alert(e)
          })
      }
      if (!this.error) {
        this.cartDialog = !this.cartDialog
        alert('まとめて返すことに成功しました。')
        let message = '入'
        for (let i = 0; i < this.returnCart.length; i++) {
          message = message + '\n[' + this.returnCart[i].name + '](' + process.env.VUE_APP_API_ENDPOINT + '/items/' + this.returnCart[i].ID + ') × ' + this.returnCart[i].returnCount
        }
        await axios.post(`${traQBaseURL}/channels/` + process.env.VUE_APP_EQUIPMENT_CHANNEL_ID + '/messages?embed=1', { text: message }).catch(e => { alert(e) })
        this.returnCart = []
        this.mount()
      }
    },
    async mount () {
      const me = await getMe()
      const res = await axios.get('/api/items?rental=' + me.data.name)
        .catch(e => {
          alert(e)
        })
      for (let i = 0; i < res.data.length; i++) {
        res.data[i].start = res.data[i].latest_logs[0].due_date.substr(0, 10)
        // res.data[i].stop = res.data[i].due_date
      }
      this.items = res.data
      let day = new Date()
      day = day.getFullYear() + '-' + ('00' + (day.getMonth() + 1)).slice(-2) + '-' + ('00' + day.getDate()).slice(-2)
      this.today = day
      this.focus = day
    },
    viewDay ({ date }) {
      this.focus = date
      this.type = 'day'
    },
    getItemColor (item) {
      if (item.type === 0) {
        return 'green'
      } else {
        return 'grey darken-1'
      }
    },
    setToday () {
      this.focus = this.today
    },
    prev () {
      this.$refs.calendar.prev()
    },
    next () {
      this.$refs.calendar.next()
    },
    showEvent ({ nativeEvent, event }) {
      const open = () => {
        this.selectedItem = event
        this.selectedElement = nativeEvent.target
        setTimeout(this.selectedOpen = true, 10)
        return true
      }

      if (this.selectedOpen) {
        this.selectedOpen = false
        setTimeout(open, 10)
      } else {
        open()
      }

      nativeEvent.stopPropagation()
    },
    updateRange ({ start, end }) {
      // You could load events from an outside source (like database) now that we have the start and end dates on the calendar
      this.start = start
      this.end = end
    },
    nth (d) {
      return d > 3 && d < 21
        ? 'th'
        : ['th', 'st', 'nd', 'rd', 'th', 'th', 'th', 'th', 'th', 'th'][d % 10]
    },
    checkDueDate (item) {
      let today = new Date()
      today = today.getFullYear() + '-' + ('00' + (today.getMonth() + 1)).slice(-2) + '-' + ('00' + today.getDate()).slice(-2)
      return today > item.start
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

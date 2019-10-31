<template>
  <nobr>
    <nobr
      class="buttonform"
      v-if="propOwner.user.name === $store.state.me.name || propOwner.user.ID === 1"
      @click.stop="open"
    >
      <v-icon style="color: #66CC33;" left>mdi-pencil</v-icon>編集
    </nobr>
    <div class="text-center">
      <v-dialog light v-model="isOpenEditOwner" max-width="290">
        <v-card width="290">
          <v-card-title class="headline">所有者の情報を更新する</v-card-title>
          <v-card-actions>
            <v-checkbox v-model="rentalable" label="貸し出し可"/>
          </v-card-actions>
          <v-card-actions>
            <div>
              <v-form ref="form">
                <v-text-field outlined v-model.number="count" type="number" label="個数"/>
              </v-form>
            </div>
          </v-card-actions>
          <v-divider/>
          <v-card-actions>
            <div class="flex-grow-1"></div>
            <v-btn @click="updateOwner()">更新</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
  </nobr>
</template>

<script>
import axios from 'axios'

export default {
  name: 'EditOwner',
  props: {
    propOwner: Object,
    propLatestLogs: Array,
    itemID: Number
  },
  data () {
    return {
      rentalable: true,
      count: 1,
      error: null,
      isOpenEditOwner: false,
      message: ''
    }
  },
  mounted () {
    this.count = this.propOwner.count
  },
  methods: {
    async updateOwner () {
      this.error = null
      if (!Number.isInteger(this.count) || this.count < 0) {
        alert('個数を正常にしてください')
        return false
      }
      if (this.count === this.propOwner.count && this.rentalable === this.propOwner.rentalable) {
        alert('変更されていません')
        return false
      }
      const latestLog = this.propLatestLogs.find(element => {
        return element.owner_id === this.propOwner.owner_id
      })
      if (latestLog) {
        if (this.count - this.propOwner.count + latestLog.count < 0) {
          alert('現在貸し出し中の物品が存在するのでそれよりも少ない数にはできません')
          return
        }
      }
      await axios.put(`/api/items/` + this.itemID + `/owners`, { user_id: this.propOwner.owner_id, rentalable: this.rentalable, count: this.count })
        .catch(e => {
          alert(e)
          this.error = e
        })
      if (this.error) {
        alert('何らかの原因で処理が完了しませんでした')
        return
      }
      if (this.count - this.propOwner.count > 0) {
        this.message = this.count - this.propOwner.count + '個追加しました'
      } else if (this.count - this.propOwner.count < 0) {
        this.message = this.propOwner.count - this.count + '個減らしました'
      }
      if (this.rentalable !== this.propOwner.rentalable) {
        this.message = this.rentalable === true ? '貸し出し可' : '貸し出し不可' + 'な物品を' + this.count + '個登録しました'
      }
      if (!this.error) { alert(this.message) }
      this.isOpenEditOwner = !this.isOpenEditOwner
      this.$emit('reload')
    },
    open () {
      this.isOpenEditOwner = !this.isOpenEditOwner
    }
  }
}
</script>

<style scoped>
.buttonform {
  color: #66CC33;
  border: solid 2px #66CC33;
  margin: 2px;
}
</style>

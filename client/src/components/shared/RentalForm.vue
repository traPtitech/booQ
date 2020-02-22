<template>
  <div>
    <v-btn block color="primary" @click.stop="open">借りる</v-btn>
    <div class="text-center">
      <v-dialog light v-model="isOpenRentalForm" max-width="320">
        <v-card width="320">
          <v-card-title class="headline">物品を借りる</v-card-title>
          <v-card-actions>
            <v-menu bottom origin="center center" transition="scale-transition" >
              <template v-slot:activator="{ on }">
                <v-btn color="primary" dark v-on="on">所有者を選ぶ</v-btn>
              </template>
              <v-list>
                <v-list-item
                v-for="(owner, i) in propItem.owners"
                :key="i"
                @click="rentOwnerID = owner.user.ID; rentOwnerName = owner.user.name"
                :disabled="$emit('checkRentalable', owner) === '貸し出しできません' || $emit('checkRentalable', owner) === '現在すべて貸しだし中'">
                  <v-list-item-title>{{ owner.user.name }}</v-list-item-title>
                </v-list-item>
              </v-list>
            </v-menu>
              <div>
                 {{rentOwnerName}}
              </div>
          </v-card-actions>
          <v-card-actions>
            <div v-if="propItem.type == 1">
              <v-form ref="form">
                <v-textarea outlined v-model="purpose" :rules="[() => !!purpose || 'This field is required']" label="目的"/>
              </v-form>
            </div>
            <div v-else>
              <v-form ref="form">
                <v-textarea outlined v-model="purpose" label="目的"/>
              </v-form>
            </div>
          </v-card-actions>
          <div v-if="propItem.type === 1">
            <div style="padding-bottom: 30px;">
              個数
            </div>
            <v-card-actions >
              <v-slider :max="getBihinLatestCount()" min="1" v-model="rentalCount" thumb-label="always" />
            </v-card-actions>
          </div>
          <div>返却日</div>
          <v-card-actions max-width="320">
            <v-date-picker v-model="dueDate"></v-date-picker>
          </v-card-actions>
          <v-divider></v-divider>
          <v-card-actions>
            <div class="flex-grow-1"></div>
            <v-btn v-on:click.stop="rental()">借りる</v-btn>
            <v-dialog light v-model="isOpenConfirm" max-width="320">
              <v-card width="320">
                <v-card-title class="headline">注意</v-card-title>
                <div style="margin-left: 10px">
                  役員には確認しましたか？
                </div>
                <div style="margin-left: 10px">
                  <a href="https://wiki.trapti.tech/general/%E5%80%89%E5%BA%AB">
                    倉庫について
                  </a>
                </div>
                <v-card-actions>
                  <v-btn @click="cancel()">Cancel</v-btn>
                  <v-btn class="green green-text" @click="rental()">Confirm</v-btn>
                </v-card-actions>
              </v-card>
            </v-dialog>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import { traQBaseURL } from '../../utils/api.js'

export default {
  name: 'RentalForm',
  props: {
    propItem: Object
  },
  data () {
    return {
      isOpenConfirm: false,
      purpose: null,
      rentalCount: 1,
      dueDate: null,
      rentOwnerID: 0,
      rentOwnerName: '',
      error: '',
      isOpenRentalForm: false
    }
  },
  methods: {
    getBihinLatestCount (itemID) {
      const item = this.propItem
      const targetLog = item.latest_logs.find(log => {
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
    async rental () {
      this.error = null
      if (this.rentOwnerID === 0) {
        alert('所有者を選択してください')
        return false
      }
      if (this.propItem.type === 1 && this.purpose === null) {
        alert('目的を入力してください')
        return false
      }
      if (this.dueDate === null) {
        alert('返却日を入力してください')
        return false
      }
      if (!this.isOpenConfirm && !this.$store.state.me.admin) {
        this.isOpenConfirm = true
        return
      }
      await axios.post(`/api/items/` + this.$route.params.id + `/logs`, { owner_id: this.rentOwnerID, type: 0, purpose: this.purpose, due_date: this.dueDate, count: this.rentalCount })
        .catch(e => {
          alert(e)
          this.error = e
          return false
        })
      if (this.error) { return false }
      if (!this.error) { alert('あなたは「' + this.propItem.name + '」を' + this.rentalCount + '個借りました。') }
      this.isOpenRentalForm = !this.isOpenRentalForm
      this.$emit('reload')
      if (this.propItem.type === 0) {
        const traQmessage = '@' + this.rentOwnerName + ' の「' + this.propItem.name + '」を借りました。\n' + process.env.VUE_APP_API_ENDPOINT + '/items/' + this.propItem.ID
        await axios.post(`${traQBaseURL}/channels/` + process.env.VUE_APP_ACTIVITY_CHANNEL_ID + '/messages?embed=1', { text: traQmessage })
          .catch(e => {
            alert(e)
            return false
          })
      } else {
        const traQmessage = '出\n[' + this.propItem.name + '](' + process.env.VUE_APP_API_ENDPOINT + '/items/' + this.propItem.ID + ')×' + this.rentalCount + '\n目的：' + this.purpose
        await axios.post(`${traQBaseURL}/channels/` + process.env.VUE_APP_EQUIPMENT_CHANNEL_ID + '/messages?embed=1', { text: traQmessage })
          .catch(e => {
            alert(e)
            return false
          })
      }
      if (this.isOpenConfirm) {
        this.isOpenConfirm = false
      }
    },
    cancel () {
      this.isOpenConfirm = false
      this.isOpenRentalForm = false
    },
    open () {
      this.isOpenRentalForm = !this.isOpenRentalForm
    }
  }
}
</script>

<template>
  <div>
    <v-btn block color="primary" @click.stop="open">借りる</v-btn>
    <div class="text-center">
      <v-dialog light v-model="isOpenRentalForm" max-width="320">
        <v-card width="320">
          <v-card-title class="headline">物品を借りる</v-card-title>
          <v-card-actions>
            <v-menu bottom origin="center center" transition="scale-transition" open-on-hover>
              <template v-slot:activator="{ on }">
                <v-btn color="primary" dark v-on="on">所有者を選ぶ</v-btn>
              </template>
              <v-list>
                <v-list-item
                v-for="(owner, i) in data.owners"
                :key="i"
                @click="rentOwnerID = owner.user.ID"
                :disabled="!owner.rentalable">
                  <v-list-item-title>{{ owner.user.name }}</v-list-item-title>
                </v-list-item>
              </v-list>
            </v-menu>
          </v-card-actions>
          <v-card-actions>
            <div v-if="data.type == 1">
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
          <v-card-actions v-if="data.type === 1">
            <v-slider :max="getBihinLatestCount()" v-model="data.rentalCount" thumb-label="always" />
          </v-card-actions>
          <div>返却日</div>
          <v-card-actions max-width="320">
            <v-date-picker v-model="dueDate"></v-date-picker>
          </v-card-actions>
          <v-divider></v-divider>
          <v-card-actions>
            <div class="flex-grow-1"></div>
            <v-btn v-on:click="rental()">借りる</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'RentalForm',
  props: {
    data: Object
  },
  data () {
    return {
      purpose: null,
      rentalCount: 1,
      dueDate: null,
      rentOwnerID: 0,
      error: '',
      isOpenRentalForm: false
    }
  },
  methods: {
    getBihinLatestCount () {
      var targetLog = this.data.latest_logs.filter(function (log) {
        return (log.owner.name = 'trap')
      })
      if (targetLog === []) {
        targetLog = targetLog.push(this.data.owners.filter(function (owner) {
          return (owner.owner.name = 'trap')
        }))
      }
      return targetLog[0].count
    },
    async rental () {
      if (this.data.type === 1 && this.purpose === null) {
        alert('目的を入力してください')
        return false
      }
      if (this.dueDate === null) {
        alert('返却日を入力してください')
        return false
      }
      if (this.rentOwnerID === 0) {
        alert('所有者を選択してください')
        return false
      }
      await axios.post(`/api/items/` + this.$route.params.id + `/logs`, { owner_id: this.rentOwnerID, type: 0, purpose: this.purpose, due_date: this.dueDate, count: this.rentalCount })
        .catch(e => {
          alert(e)
          this.error = e
        })
      if (!this.error) { alert('あなたは”' + this.data.name + '”を' + this.rentalCount + '個借りました。') }
      this.isOpenRentalForm = !this.isOpenRentalForm
    },
    open () {
      this.isOpenRentalForm = !this.isOpenRentalForm
    }
  }
}
</script>

<template>
  <div>
    <v-btn block color="success" @click.stop="open">借りたい</v-btn>
    <div class="text-center">
      <v-dialog light v-model="isOpenWannaRentalForm" max-width="320">
        <v-card width="320">
          <v-card-title class="headline">所有者にメッセージを送る</v-card-title>
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
                <v-textarea outlined v-model="message" :rules="[() => !!message || 'This field is required']" label="目的"/>
              </v-form>
            </div>
            <div v-else>
              <v-form ref="form">
                <v-textarea outlined v-model="message" label="文面"/>
              </v-form>
            </div>
          </v-card-actions>
          <v-divider/>
          <v-card-actions>
            <div class="flex-grow-1"></div>
            <v-btn v-on:click="send()">送る</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
  </div>
</template>

<script>
export default {
  name: 'WannaRental',
  props: {
    propItem: Object
  },
  data () {
    return {
      message: null,
      rentalCount: 1,
      dueDate: null,
      rentOwnerID: 0,
      rentOwnerName: '',
      error: '',
      isOpenWannaRentalForm: false
    }
  },
  mounted () {
    this.message = '[' + this.propItem.name + '](https://' + location.host + '/items/' + this.propItem.ID + ')を借りたいです。'
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
    async send () {
      this.error = null
      if (this.rentOwnerID === 0) {
        alert('所有者を選択してください')
        return false
      }
      window.location.href = `https://q.trap.jp/users/${this.rentOwnerName}`
    },
    open () {
      this.isOpenWannaRentalForm = !this.isOpenWannaRentalForm
    }
  }
}
</script>

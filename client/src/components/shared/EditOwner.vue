<template>
  <nobr>
    <v-btn class="ma-2" tile outlined color="success" v-if="propOwner.user.name === $store.state.me.name || propOwner.user.ID === 1" @click.stop="open">
      <v-icon left>mdi-pencil</v-icon>編集
    </v-btn>
    <div class="text-center">
      <v-dialog light v-model="isOpenEditOwner" max-width="290">
        <v-card width="290">
          <v-card-title class="headline">所有者の情報を更新する</v-card-title>
          <v-card-actions>
            <div>
              <input type="checkbox" id="checkbox" v-model="rentalable">
              <label for="checkbox">貸し出し可</label>
            </div>
          </v-card-actions>
          <v-card-actions>
            <div>
              <v-form ref="form">
                <v-text-field outlined v-model.number="count" type="number" label="個数"/>
              </v-form>
            </div>
          </v-card-actions>
          <v-divider></v-divider>
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
    propOwner: Object
  },
  data () {
    return {
      rentalable: true,
      count: 1,
      error: '',
      isOpenEditOwner: false,
      message: ''
    }
  },
  mounted () {
    this.count = this.propOwner.count
  },
  methods: {
    async updateOwner () {
      if (!Number.isInteger(this.count) || this.count < 0 || this.count === this.propOwner.count) {
        alert('個数を正常にしてください')
        return false
      }
      await axios.post(`/api/items/` + this.$route.params.id + `/owners`, { user_id: this.$store.state.me.ID, rentalable: this.rentalable, count: this.count - this.propOwner.count })
        .catch(e => {
          alert(e)
          this.error = e
        })
      if (this.count - this.propOwner.count > 0) {
        this.message = this.count - this.propOwner.count + "個追加しました"
      }if (this.count - this.propOwner.count < 0) {
        this.message = this.propOwner.count - this.count + "個減らしました"
      }
      if (this.error) { alert("何らかの原因で処理が完了しませんでした") }
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

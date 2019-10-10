<template>
  <nobr>
    <v-btn x-small outlined fab dark color="primary" @click.stop="open">
      <v-icon dark>mdi-plus</v-icon>
    </v-btn>
    <div class="text-center">
      <v-dialog light v-model="isOpenAddOwner" max-width="290">
        <v-card width="290">
          <v-card-title class="headline">所有者を追加する</v-card-title>
          <v-card-actions>
            <div v-if="$store.state.me && $store.state.me.admin" >
              <label v-for="owner in ownerOptions" :key="owner.id">
                <div><input type="radio" name="owner" :value="owner.id" v-model="ownerID">{{ owner.name }}</div>
              </label>
            </div>
          </v-card-actions>
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
            <v-btn @click="add()">追加</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
  </nobr>
</template>

<script>
import axios from 'axios'

export default {
  name: 'RegisterOwnerForm',
  data () {
    return {
      ownerID: 0,
      ownerOptions: [
        { id: 0, name: '自身' },
        { id: 1, name: 'traP' },
        { id: 2, name: '支援課' }
      ],
      rentalable: true,
      count: 1,
      error: '',
      isOpenAddOwner: false,
      message: ''
    }
  },
  methods: {
    async add () {
      if (!Number.isInteger(this.count) || this.count < 0) {
        alert('個数を正常にしてください')
        return false
      }
      let newOwnerID = 0
      if (Number(this.ownerID) === 0) {
        this.message = '所有者に' + this.$store.state.me.name + 'を追加しました。'
        newOwnerID = this.$store.state.me.ID
      } else {
        this.message = '所有者に' + this.ownerOptions.find(element => element.id === this.ownerID).name + 'を追加しました。'
        newOwnerID = this.ownerID
      }
      await axios.post(`/api/items/` + this.$route.params.id + `/owners`, { user_id: newOwnerID, rentalable: this.rentalable, count: this.count })
        .catch(e => {
          alert(e)
          this.error = e
        })
      if (!this.error) { alert(this.message) }
      this.isOpenAddOwner = !this.isOpenAddOwner
      this.$emit('reload')
    },
    open () {
      this.isOpenAddOwner = !this.isOpenAddOwner
    }
  }
}
</script>

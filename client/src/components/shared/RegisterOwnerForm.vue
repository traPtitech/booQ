<template>
  <div>
    <v-btn color="primary" @click.stop="open">所有者を追加</v-btn>
    <div class="text-center">
      <v-dialog light v-model="isOpenAddOwner" max-width="290">
        <v-card width="290">
          <v-card-title class="headline">所有者を追加する</v-card-title>
          <v-card-actions>
            <div v-if="$store.state.me.admin" >
              <label v-for="(label,id) in ownerOptions" v-bind:key="label">
                <div><input type="radio" name="owner" :value="id" v-model="ownerID">{{ label }}</div>
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
          <v-alert type="error" v-if="count<0">個数が負になっています</v-alert>
          <v-divider></v-divider>
          <v-card-actions>
            <div class="flex-grow-1"></div>
            <v-btn v-on:click="add()">追加</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'RegisterOwnerForm',
  data () {
    return {
      ownerID: 0,
      ownerOptions: {
        0: '自身',
        1: 'traP',
        2: '支援課'
      },
      rentalable: true,
      count: 1,
      error: '',
      isOpenAddOwner: false
    }
  },
  methods: {
    async add () {
      if (this.ownerID === 0) {
        await axios.post(`/api/items/` + this.$route.params.id + `/owners`, { user_id: this.$store.state.me.ID, rentalable: this.rentalable, count: this.count })
          .catch(e => {
            alert(e)
            this.error = e
          })
        if (!this.error) { alert('”' + this.data.name + '”の所有者に' + this.$store.state.me.name + 'を追加しました。') }
      } else {
        await axios.post(`/api/items/` + this.$route.params.id + `/owners`, { user_id: this.ownerID, rentalable: this.rentalable, count: this.count })
          .catch(e => {
            alert(e)
            this.error = e
          })
        if (!this.error) { alert('”' + this.data.name + '”の所有者に' + this.ownerOptions[this.ownerID] + 'を追加しました。') }
      }
      this.isOpenAddOwner = !this.isOpenAddOwner
    },
    open () {
      this.isOpenAddOwner = !this.isOpenAddOwner
    }
  }
}
</script>

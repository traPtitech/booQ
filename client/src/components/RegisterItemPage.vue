<template>
  <div>
    <h1>物品登録ページ</h1>
    <div v-if="$store.state.me.admin">
      <span>登録する物品の所有者: {{ ownerOptions[ownerID] }}</span>
      <br>
      <label v-for="(label,id) in ownerOptions" v-bind:key="label">
        <input type="radio" name="owner" :value="id" v-model="ownerID">{{ label }}
      </label>
    </div>
    <div>
      <v-text-field solo v-model="code" placeholder="ISBN-10 or ASIN"/>
      <v-btn class="green green-text darken-2" v-on:click="img = 'http://images-jp.amazon.com/images/P/' + code + '.09.LZZZZZZZ.jpg';img_name = 'by amazon'">MakeImage</v-btn>
    </div>
    <div>
      <div>物品名</div>
      <v-text-field class="mt-0" solo required v-model="name" placeholder="Name"/>
    </div>
    <div>
      <div>物品詳細</div>
      <v-textarea class="mt-0" solo label="Solo textarea" rows="1" auto-grow v-model="description" placeholder="Description"></v-textarea>
    </div>
    <div>
      <div>物品イメージ</div>
      <label class="input-item__label">
        <input type="file" @change="onFileChange" />
      </label>
      <div class="preview-item">
        <v-img
          :src="img"
          aspect-ratio="1"
          position="left"
          :contain="true"
          max-height="500px"
        />
        <div>
          <p>{{ img_name }}</p>
        </div>
        <v-btn class="red" @click="remove">削除</v-btn>
      </div>
    </div>
    <div>
      <p>個数</p>
      <v-text-field class="mt-0" solo required v-model.number="count" type="number"/>
    </div>
    <div>
      <input type="checkbox" id="checkbox" v-model="rentalable">
      <label for="checkbox">貸し出し可</label>
    </div>
    <v-btn class="blue" @click="register">登録</v-btn>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'RegisterItemPage',
  data () {
    return {
      ownerID: 0,
      ownerOptions: {
        0: '個人',
        1: 'traP',
        2: '支援課'
      },
      rentalable: true,
      code: '',
      img: 'https://q.trap.jp/api/1.0/files/3380fbc6-6141-4b60-99ae-a1d270842d60/thumbnail', // ここは適当に変えてください
      name: '',
      description: '',
      img_name: '',
      img_url: '',
      count: 1
    }
  },
  methods: {
    async register () {
      if (this.img_name) {
        // img_urlに画像のURLをセットしてください
      }
      const res = await axios.post(`/api/items`, { name: this.name, code: this.code, type: this.ownerID, description: this.description, img_url: this.img_url }).catch(e => { alert(e) })
      if (this.ownerID === 0) {
        const res2 = await axios.post(`/api/items/` + res.data.ID + `/owners`, { user_id: this.$store.state.me.ID, rentalable: this.rentalable, count: this.count }).catch(e => { alert(e) })
        alert('Registered ”' + res2.data.name + '”!所有者は' + this.$store.state.me.name + 'です。"')
      } else {
        const res2 = await axios.post(`/api/items/` + res.data.ID + `/owners`, { user_id: this.ownerID, rentalable: this.rentalable, count: this.count }).catch(e => { alert(e) })
        alert('Registered ”' + res2.data.name + '”!所有者は' + this.ownerOptions[this.ownerID] + 'です。')
      }
    },
    onFileChange (e) {
      const files = e.target.files || e.dataTransfer.files
      this.createImage(files[0])
      this.img_name = files[0].name
    },
    createImage (file) {
      const reader = new FileReader()
      reader.onload = e => {
        this.img = e.target.result
      }
      reader.readAsDataURL(file)
    },
    remove () {
      this.img = 'https://q.trap.jp/api/1.0/files/3380fbc6-6141-4b60-99ae-a1d270842d60/thumbnail'
      this.img_name = ''
    }
  }
}
</script>

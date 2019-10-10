<template>
  <div>
    <h1>物品登録ページ</h1>
    <div
      v-if="$store.state.me && $store.state.me.admin"
      class="contents"
    >
      <span>登録する物品の所有者: {{ ownerOptions[ownerID].name }}</span>
      <br>
      <label v-for="owner in ownerOptions" :key="owner.id">
        <div><input type="radio" name="owner" :value="owner.id" v-model="ownerID">{{ owner.name }}</div>
      </label>
    </div>
    <div class="contents">
      <div>物品コードもしくはISBNコード</div>
      <v-text-field v-model="code" placeholder="ISBN-13 or ASIN"/>
      <v-btn class="green green-text" @click="getBookInformation">
        自動入力
      </v-btn>
    </div>
    <div>
      <div>物品名</div>
      <v-text-field class="mt-0" required v-model="name" placeholder="Name" />
    </div>
    <div>
      <div>物品詳細</div>
      <v-textarea class="mt-0" rows="1" auto-grow v-model="description" placeholder="Description" />
    </div>
    <div class="contents">
      <div>物品イメージ</div>
      <label class="input-item__label">
        <input type="file" @change="onFileChange" />
      </label>
      <div class="preview-item">
        <v-img
          :src="img_url"
          aspect-ratio="1"
          position="left"
          :contain="true"
          max-height="400px"
        />
        <div>
          <p>{{ img_name }}</p>
        </div>
        <v-btn class="error" @click="remove">画像削除</v-btn>
      </div>
    </div>
    <div>
      個数
      <v-text-field class="mt-0" required v-model.number="count" type="number"/>
    </div>
    <div class="contents">
      <input type="checkbox" id="checkbox" v-model="rentalable">
      <label for="checkbox">貸し出し可</label>
    </div>
    <div>
      <v-btn class="primary" @click="register">登録</v-btn>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import { traQBaseURL } from '../utils/api.js'

export default {
  name: 'RegisterItemPage',
  data () {
    return {
      ownerID: 0,
      ownerOptions: [
        { id: 0, name: '自身' },
        { id: 1, name: 'traP' },
        { id: 2, name: '支援課' }
      ],
      rentalable: true,
      code: '',
      name: '',
      description: '',
      img_name: '',
      img_url: 'https://q.trap.jp/api/1.0/files/3380fbc6-6141-4b60-99ae-a1d270842d60/thumbnail', // NoImage
      count: 1
    }
  },
  watch: {
    code (val) {
      this.code = this.code.replace('-', '')
    }
  },
  methods: {
    async register () {
      const res = await axios.post(`/api/items`, { name: this.name, code: this.code, type: Number(this.ownerID), description: this.description, img_url: this.img_url }).catch(e => { alert(e) })
      if (!res) {
        alert('エラーが発生したため物品の登録が行われませんでした。')
        return
      }
      const itemID = res.data.ID
      const userID = Number(this.ownerID) === 0 ? Number(this.$store.state.me.ID) : Number(this.ownerID)
      const res2 = await axios.post(`/api/items/` + itemID + `/owners`, { user_id: userID, rentalable: this.rentalable, count: Number(this.count) }).catch(e => { alert(e) })
      if (!res2) {
        alert('エラーが発生したため所有者の登録が行われませんでした。')
        return
      }
      await axios.post(`${traQBaseURL}/channels/` + process.env.VUE_APP_ACTIVITY_CHANNEL_ID + `/messages?embed=1`, {
        text: '@' + this.$store.state.me.name + ' が「' + this.name + '」を登録しました。\n' + process.env.VUE_APP_API_ENDPOINT + '/items/' + itemID
      }).catch(e => { alert(e) })
      this.$router.push({ path: `/items/${itemID}` })
    },
    onFileChange (e) {
      const files = e.target.files || e.dataTransfer.files
      this.createImage(files[0])
      this.img_name = files[0].name
    },
    async createImage (file) {
      let form = new FormData()
      form.enctype = 'multipart/form-data'
      form.append('file', file)
      const data = await axios.post(`${traQBaseURL}/files`, form).catch(e => alert(e))
      if (!data) {
        alert('画像の投稿に失敗しました。')
      }
      this.img_url = `${traQBaseURL}/files/${data.data.fileId}/thumbnail`
    },
    remove () {
      this.img_url = 'https://q.trap.jp/api/1.0/files/3380fbc6-6141-4b60-99ae-a1d270842d60/thumbnail' // NoImage
      this.img_name = ''
    },
    isbn13Toisbn10 (isbn13) {
      let isbn13Array = []
      for (let i = 0; i < isbn13.length; i++) {
        isbn13Array.push(parseInt(isbn13.charAt(i), 10))
      }
      let checkDegit = 11 - ((isbn13Array[3] * 10 + isbn13Array[4] * 9 + isbn13Array[5] * 8 + isbn13Array[6] * 7 + isbn13Array[7] * 6 + isbn13Array[8] * 5 + isbn13Array[9] * 4 + isbn13Array[10] * 3 + isbn13Array[11] * 2) % 11)
      if (checkDegit === 10) {
        checkDegit = 'X'
      } else if (checkDegit === 11) {
        checkDegit = 0
      }
      let res = isbn13Array.slice(3, 12)
      res.push(checkDegit)
      return res.join('')
    },
    async getBookInformation () {
      if (this.code.length === 10 || this.code.length === 13) {
        const resp = await axios.get(`https://api.openbd.jp/v1/get?isbn=${this.code}`)
        if (resp.data[0]) {
          this.data = resp.data[0]['onix']
          this.name = this.data['DescriptiveDetail']['TitleDetail']['TitleElement']['TitleText']['content']
          if (this.data['CollateralDetail']['TextContent']) {
            this.description = this.data['CollateralDetail']['TextContent'][0]['Text']
          }
          if (this.data['CollateralDetail']['SupportingResource']) {
            this.img_url = this.data['CollateralDetail']['SupportingResource'][0]['ResourceVersion'][0]['ResourceLink']
          }
        } else {
          alert('本がみつかりませんでした。')
        }
      } else {
        alert('不正な値です。')
      }
    }
  }
}
</script>

<style scoped>
.contents {
  padding-bottom: 30px;
}
</style>

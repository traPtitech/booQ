<template>
  <div>
    <Dialog :dialog="dialog" target="alert">
      <template v-slot:headline>
        <h2 style="color:red;">エラー</h2>
      </template>
      <template v-slot:content>
        <h3>{{ errorMessage }}</h3>
      </template>
    </Dialog>
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
      <v-btn class="green green-text" @click.stop="setDialog('close','barcodeDialog')">
        バーコード読み取り
      </v-btn>
      <Dialog :dialog="dialog" target="barcodeDialog">
        <template v-slot:content>
          <BarCode  @search="getBookInformation" @changeCode="changeCode"/>
        </template>
      </Dialog>
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
          v-if="!loading && img_url.length"
          :src="img_url"
          aspect-ratio="1"
          position="left"
          contain
          max-height="400px"
        />
        <v-progress-circular
          v-else-if="loading"
          :size="70"
          :width="7"
        />
        <div>
          <p>{{ img_name }}</p>
        </div>
        <v-btn
          v-if="img_url.length"
          class="error"
          @click="remove"
        >
          画像削除
        </v-btn>
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
import BarCode from './BarCode'
import Dialog from './shared/Dialog'
export default {
  name: 'RegisterItemPage',
  components: {
    BarCode,
    Dialog
  },
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
      img_url: '',
      count: 1,
      dialog: {
        isOpen: false,
        closeText: 'close',
        target: ''
      },
      errorMessage: 'エラー',
      loading: false
    }
  },
  watch: {
    code (val) {
      this.code = this.code.replace('-', '')
    }
  },
  methods: {
    async register () {
      const res = await axios.post('/api/items', { name: this.name, code: this.code, type: Number(this.ownerID), description: this.description, img_url: this.img_url }).catch(e => { alert(e) })
      if (!res) {
        this.setAlert('close', 'エラーが発生したため物品の登録が行われませんでした。')
        return
      }
      const itemID = res.data.ID
      const userID = Number(this.ownerID) === 0 ? Number(this.$store.state.me.ID) : Number(this.ownerID)
      const res2 = await axios.post('/api/items/' + itemID + '/owners', { user_id: userID, rentalable: this.rentalable, count: Number(this.count) }).catch(e => { alert(e) })
      if (!res2) {
        this.setAlert('close', 'エラーが発生したため所有者の登録が行われませんでした。')
        return
      }
      await axios.post(`${traQBaseURL}/channels/` + process.env.VUE_APP_ACTIVITY_CHANNEL_ID + '/messages?embed=1', {
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
      this.loading = true
      const form = new FormData()
      form.enctype = 'multipart/form-data'
      form.append('file', file)
      const data = await axios.post('/api/files', form).catch(e => alert(e))
      if (!data) {
        this.setAlert('close', '画像の投稿に失敗しました')
      }
      this.img_url = data.data.url
      this.loading = false
    },
    remove () {
      this.img_url = ''
      this.img_name = ''
    },
    isbn13Toisbn10 (isbn13) {
      const isbn13Array = []
      for (let i = 0; i < isbn13.length; i++) {
        isbn13Array.push(parseInt(isbn13.charAt(i), 10))
      }
      let checkDegit = 11 - ((isbn13Array[3] * 10 + isbn13Array[4] * 9 + isbn13Array[5] * 8 + isbn13Array[6] * 7 + isbn13Array[7] * 6 + isbn13Array[8] * 5 + isbn13Array[9] * 4 + isbn13Array[10] * 3 + isbn13Array[11] * 2) % 11)
      if (checkDegit === 10) {
        checkDegit = 'X'
      } else if (checkDegit === 11) {
        checkDegit = 0
      }
      const res = isbn13Array.slice(3, 12)
      res.push(checkDegit)
      return res.join('')
    },
    async getBookInformation () {
      if (this.code.length === 10 || this.code.length === 13) {
        this.data = ''
        this.name = ''
        this.description = ''
        this.img_url = ''
        const openbd = axios
          .get(`https://api.openbd.jp/v1/get?isbn=${this.code}`)
          .then(async resp => {
            if (resp.data[0]) {
              this.data = resp.data[0].onix
              this.name = this.data.DescriptiveDetail.TitleDetail.TitleElement.TitleText.content
              if (this.data.CollateralDetail.TextContent) {
                this.description = this.data.CollateralDetail.TextContent[0].Text
              }
              if (this.data.CollateralDetail.SupportingResource) {
                this.img_url = this.data.CollateralDetail.SupportingResource[0].ResourceVersion[0].ResourceLink
              }
            }
            const index = runnings.findIndex(v => v === openbd)
            if (index === -1) {
              return false
            }
            runnings.splice(index, 1)
            return resp.data[0]
          })
        const googleBooksAPI = axios
          .get(
            `https://www.googleapis.com/books/v1/volumes?q=isbn:${this.code}&maxResults=1`
          )
          .then(async resp => {
            if (resp.data.totalItems !== 0) {
              this.data = resp.data.items[0].volumeInfo
              this.name = this.data.title
              if (this.data.description) {
                this.description = this.data.description
              }
              if (this.data.imageLinks.thumbnail) {
                this.img_url = this.data.imageLinks.thumbnail
              }
            }
            const index = runnings.findIndex(v => v === googleBooksAPI)
            if (index === -1) {
              return false
            }
            runnings.splice(index, 1)
            return resp.data.totalItems !== 0
          })
        const runnings = [openbd, googleBooksAPI]
        let result = await Promise.race(runnings)
        while (runnings.length > 0 && !result) {
          result = await Promise.race(runnings)
        }
        if (!result) {
          this.setAlert('close', '本が見つかりませんでした')
        }
      } else {
        this.setAlert('close', '不正な値です')
      }
    },
    changeCode (code) {
      this.code = code
      this.dialog = {
        isOpen: false,
        closeText: '',
        target: ''
      }
    },
    setDialog (closeText, target) {
      this.dialog = {
        isOpen: true,
        closeText: closeText,
        target: target
      }
    },
    setAlert (closeText, errmsg) {
      this.errorMessage = errmsg
      this.setDialog(closeText, 'alert')
    }
  }
}
</script>

<style scoped>
.contents {
  padding-bottom: 30px;
}
</style>

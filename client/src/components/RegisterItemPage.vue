<template>
  <v-container>
    <h1>物品登録ページ</h1>
    <!-- <h4>物品の種類を選択してください</h4>
    <v-container>
      <label v-for="(label,id) in typeOptions" v-bind:key="label">
        <input type="radio" name="type" :value="id" v-model="typeID">{{ label }}
      </label>
      <br>
      <span>登録する物品の種類: {{ typeOptions[typeID] }}</span>
    </v-container> -->
    <v-container v-if="$store.state.me.admin">
      <span>登録する物品の所有者: {{ ownerOptions[ownerID] }}</span>
      <br>
      <label v-for="(label,id) in ownerOptions" v-bind:key="label">
        <input type="radio" name="owner" :value="id" v-model="ownerID">{{ label }}
      </label>
    </v-container>
    <v-container>
      <input type="checkbox" id="checkbox" v-model="rentalable">
      <label for="checkbox">貸し出し可</label>
    </v-container>
    <v-container>
      <v-text-field solo v-model="code" placeholder="ISBN-10 or ASIN"/>
      <v-btn class="green green-text darken-2" v-on:click="img = 'http://images-jp.amazon.com/images/P/' + code + '.09.LZZZZZZZ.jpg';img_name = 'by amazon'">MakeImage</v-btn>
    </v-container>
    <v-container>
      <v-container>物品名</v-container>
      <v-text-field class="mt-0" solo required v-model="name" placeholder="Name"/>
    </v-container>
    <v-container>
      <v-container>物品詳細</v-container>
      <v-textarea class="mt-0" solo label="Solo textarea" rows="1" auto-grow v-model="description" placeholder="Description"></v-textarea>
    </v-container>
    <v-container>
      <v-container>物品イメージ</v-container>
      <label class="input-item__label">
        <input type="file" @change="onFileChange" />
      </label>
      <v-container class="preview-item">
        <v-img
          :src="img"
          aspect-ratio="1"
          position="left"
          :contain="true"
          max-height="500px"
        />
        <v-container>
          <p>{{ img_name }}</p>
        </v-container>
        <v-btn class="red" @click="remove">削除</v-btn>
      </v-container>
    </v-container>
    <v-container>
      <p>個数</p>
      <v-text-field class="mt-0" solo required v-model.number="count" type="number"/>
    </v-container>
    <v-btn class="blue" @click="register">登録</v-btn>
  </v-container>
</template>

<script>
import axios from 'axios'
import { getMe } from '@/utils/api'

export default {
  name: 'RegisterItemPage',
  data() {
    return{
      // typeID: 1,
      // typeOptions: {
      //   1: '本',
      //   0: '備品',
      // },
      ownerID: 0,
      ownerOptions: {
        0: '個人',
        1: 'traP',
        2: '支援課'
      },
      rentalable: true,
      code: '',
      img: 'https://q.trap.jp/api/1.0/files/3380fbc6-6141-4b60-99ae-a1d270842d60/thumbnail', //ここは適当に変えてください
      name: '',
      description: '',
      img_name: '',
      img_url: '',
      count: 1,
      res: ""
    };
  },
  methods: {
    register: function () {
      if (this.img_name != '') {
        //img_urlに画像のURLをセットしてください
      }
      axios.post(`/api/items`, { name: this.name, code: this.code, type: this.ownerID, description: this.description, img_url:this.img_url})
        .then(res => {
          if (this.ownerID==0) {
            axios.post(`/api/items/` + res.data.ID + `/owners`, {user_id: this.$store.state.me.ID, rentalable: this.rentalable})
              .then(res => {
                alert("Registered ”" + res.data.name + "”!所有者は" + this.$store.state.me.name + "です。")
              }).catch( e => {alert(e)})
          } else {
            axios.post(`/api/items/` + res.data.ID + `/owners`, {user_id: this.ownerID, rentalable: this.rentalable})
              .then(res => {
                alert("Registered ”" + res.data.name + "”!所有者は" + this.ownerOptions[this.ownerID] + "です。")
              }).catch( e => {alert(e)})
          }
        }).catch( e => {
          alert(e)
        })
    },
    onFileChange(e) {
      const files = e.target.files || e.dataTransfer.files;
      this.createImage(files[0]);
      this.img_name = files[0].name;
    },
    createImage(file) {
      const reader = new FileReader();
      reader.onload = e => {
        this.img = e.target.result;
      };
      reader.readAsDataURL(file);
    },
    remove() {
      this.img = 'https://q.trap.jp/api/1.0/files/3380fbc6-6141-4b60-99ae-a1d270842d60/thumbnail';
      this.img_name = ''
    }
  }
}
</script>

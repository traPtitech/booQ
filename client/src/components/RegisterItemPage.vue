<template>
  <v-container>
    <h1>物品登録ページ</h1>
    <h4>物品の種類を選択してください</h4>
    <v-container>
      <label v-for="(label,id) in typeOptions" v-bind:key="label">
        <input type="radio" name="type" :value="id" v-model="typeID">{{ label }}
      </label>
      <br>
      <span>登録する物品の種類: {{ typeOptions[typeID] }}</span>
    </v-container>
    <v-container>
      <label v-for="(label,id) in ownerOptions" v-bind:key="label">
        <input type="radio" name="owner" :value="id" v-model="ownerID">{{ label }}
      </label>
      <br>
      <span>登録する物品の所有者: {{ ownerOptions[ownerID] }}</span>
    </v-container>
    <v-container>
      <input type="checkbox" id="checkbox" v-model="checked">
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
        <img
          v-show="img"
          :src="img"
          alt=""
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
    <v-btn @click="register">登録</v-btn>
  </v-container>
</template>

<script>
export default {
  name: 'RegisterItemPage',
  data () {
    return {
      typeID: 1,
      typeOptions: {
        1: '本',
        0: '備品'
      },
      ownerID: 0,
      ownerOptions: {
        0: '個人',
        // 以下はAdminUserのみ表示されるように(おそらく後のissueのタスク)
        1: 'traP',
        2: '支援課'
      },
      checked: true,
      code: '',
      img: 'https://q.trap.jp/api/1.0/files/3380fbc6-6141-4b60-99ae-a1d270842d60/thumbnail', // ここは適当に変えてください
      name: '',
      description: '',
      img_name: '',
      count: 1
    }
  },
  methods: {
    register: function () {
      // axios.post~~~
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

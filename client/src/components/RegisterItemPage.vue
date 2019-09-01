<template>
  <div>
    <h1>物品登録ページ</h1>
    <h4>物品の種類を選択してください</h4>
    <div>
      <label v-for="(label,id) in typeOptions" v-bind:key="label">
        <input type="radio" name="type" :value="id" v-model="typeID">{{ label }}
      </label>
      <br>
      <span>登録する物品の種類: {{ typeOptions[typeID] }}</span>
    </div>
    <div>
      <label v-for="(label,id) in ownerOptions" v-bind:key="label">
        <input type="radio" name="owner" :value="id" v-model="ownerID">{{ label }}
      </label>
      <br>
      <span>登録する物品の所有者: {{ ownerOptions[ownerID] }}</span>
    </div>
    <div>
      <input type="checkbox" id="checkbox" v-model="checked">
      <label for="checkbox">貸し出し可</label>
    </div>
    <div>
      <input v-model="code" placeholder="ISBN-10 or ASIN">
      <button v-on:click="img = 'http://images-jp.amazon.com/images/P/' + code + '.09.LZZZZZZZ.jpg';img_name = 'by amazon'">MakeImage</button>
    </div>
    <div>
      <div>物品名</div>
      <input v-model="name" placeholder="Name">
    </div>
    <div>
      <div>物品詳細</div>
      <textarea v-model="description" placeholder="Description"></textarea>
    </div>
    <div>
      <div>物品イメージ</div>
      <label class="input-item__label">
        <input type="file" @change="onFileChange" />
      </label>
      <div class="preview-item">
        <img
          v-show="img"
          :src="img"
          alt=""
        />
        <div>
          <p>{{ img_name }}</p>
        </div>
        <button v-on:click="remove">削除</button>
      </div>
    </div>
    <div>
      <p>個数</p>
      <input v-model.number="count" type="number">
    </div>
    <button v-on:click="register">登録</button>
  </div>
</template>

<script>
export default {
  name: 'RegisterItemPage',
  data() {
    return{
      typeID: 1,
      typeOptions: {
        1: '本',
        0: '備品',
      },
      ownerID: 0,
      ownerOptions: {
        0: '個人',
        //以下はAdminUserのみ表示されるように(おそらく後のissueのタスク)
        1: 'traP',
        2: '支援課'
      },
      checked: true,
      code: '',
      img: 'https://q.trap.jp/api/1.0/files/3380fbc6-6141-4b60-99ae-a1d270842d60/thumbnail', //ここは適当に変えてください
      name: '',
      description: '',
      img_name: '',
      count: 1,
    };
  },
  methods: {
    register: function () {
        //axios.post~~~
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
    },
  }
}
</script>

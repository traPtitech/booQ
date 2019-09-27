<template>
  <v-container>
    <div class="wrapper" v-for="item in items" :key="item.id">
      <div>
        <v-col md="auto">
          <img :src="item.img_url" />
        </v-col>
      </div>
      <div>
        <v-col lg="2">
          <h3>{{item.name}}</h3>
          <div v-for="owner in item.owners" :key="owner.id">
            <p>{{owner.user.name}}</p>
          </div>
        </v-col>
      </div>
    </div>
  </v-container>
</template>

<script>
export default {
  name: 'ItemList',
  props: [
    'items'
  ]
}
</script>

<style>
.wrapper {
    display:flex;
    display:-ms-flexbox;/* --- IE10用 11はこの設定は不要 --- */
    display:-webkit-box;/*--- Android用 ---*/
    /*画面中央に表示されるように margin: auto;を設定している*/
    margin: auto;
    /* justify-content:stretch; */
}
</style>

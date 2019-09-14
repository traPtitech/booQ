<template>
  <div>
    <h4>所有者を追加する</h4>
    <label v-for="(label,id) in ownerOptions" v-bind:key="label">
      <div><input type="radio" name="owner" :value="id" v-model="ownerID">{{ label }}</div>
    </label>
    <button v-on:click="add(ownerID)">追加</button>
  </div>
</template>

<script>
export default {
  name: 'RegisterOwnerForm',
  data() {
    return{
      ownerID: 0,
      ownerOptions: {
        0: '自身',
        //以下はAdminUserのみ表示されるように(おそらく後のissueのタスク)
        1: 'traP',
        2: '支援課'
      },
    }
  },
  methods: {
    add(num) {
      // axios.postで
    }
  }
}
</script>

<template>
  <div>
    <h4>他のユーザーに権限を付与します</h4>
    <!-- @ryoha @trapみたいな感じで指定することを考えています -->
    <p>付与するユーザー</p><input v-model="target">
    <button v-on:click="submit">attach</button>
  </div>
</template>

<script>
export default {
  name: 'AdminPage',
  data () {
    return {
      target: ''
    }
  },
  methods: {
    attach () {
      // this.targetとかをaxios.postで
    }
  }
}
</script>

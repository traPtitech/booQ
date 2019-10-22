<template>
<v-btn dark icon :to="`/users/${user.name}`">
  <v-avatar :size="size">
    <v-img :src="checkSienka?'https://q.trap.jp/api/1.0/public/icon/' + user.name:sienka" width="100%" alt=""/>
  </v-avatar>
</v-btn>
</template>

<script>
export default {
  name: 'Icon',
  data () {
    return {
      sienka: './../../img/sienka-icon.jpg'
    }
  },
  computed: {
    checkSienka () {
      return this.user.name !== 'sienka'
    }
  },
  props: {
    user: Object,
    size: {
      type: Number,
      default: 40
    }
  }
}
</script>

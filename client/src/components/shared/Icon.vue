<template>
  <v-avatar :size="size">
    <v-img :src="'https://q.trap.jp/api/1.0/public/icon/' + user.name" width="100%" alt=""/>
  </v-avatar>
</template>

<script>
export default {
  name: 'Icon',
  props: {
    user: Object,
    size: {
      type: Number,
      default: 40
    }
  }
}
</script>

<template>
<v-btn
          dark
          icon
          :to="`/users/${user.name}`"
        >
  <v-avatar :size="size">
    <v-img :src="user.name!=='sienka'?'https://q.trap.jp/api/1.0/public/icon/' + user.name:'https://pbs.twimg.com/profile_images/877073041351622657/n31PKuuM.jpg'" width="100%" alt=""/>
  </v-avatar>
</v-btn>
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

<template>
  <!-- <div> -->
    <v-avatar size="40">
      <v-img :src="'https://q.trap.jp/api/1.0/public/icon/' + user.name" width="100%" alt=""/>
    </v-avatar>
  <!-- </div> -->
</template>

<script>
export default {
  name: 'Icon',
  props: [
    'user'
    // 'size'
  ]
}
</script>

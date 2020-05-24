<template>
<v-btn dark icon :to="`/users/${user.name}`">
  <v-avatar :size="size">
    <v-img :src="navigateImagePath"  width="100%" alt=""/>
  </v-avatar>
</v-btn>
</template>

<script>
export default {
  name: 'Icon',
  computed: {
    navigateImagePath () {
      return this.user.name !== 'sienka' ? `https://q.trap.jp/api/v3/public/icon/${this.user.name}` : './../../img/sienka-icon.jpg'
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

<template>
  <v-content>
    <div id="core-view">
      <v-fade-transition mode="out-in">
        <router-view />
      </v-fade-transition>
    </div>
  </v-content>
</template>

<script>
export default {
  metaInfo () {
    return {
      title: 'Vuetify Material Dashboard by CreativeTim'
    }
  }
}
</script>

<style>
#core-view {
  padding: 15px 0 100px 15px;
}
</style>

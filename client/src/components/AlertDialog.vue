<template>
  <div>
    <v-dialog v-model="alert.isAlert" persistent>
      <v-alert :type="alert.alertType" v-model="alert.isAlert" dismissible style="margin-bottom:0px;">
        <h3>{{ this.alert.title }}</h3>
        {{ this.alert.message }}
      </v-alert>
    </v-dialog>
  </div>
</template>
<script>
export default {
  props: {
    alert: {
      type: Object,
      default: () => ({
        isAlert: false,
        alertType: 'error',
        title: 'Error',
        message: 'error'
      })
    }
  }
}
</script>

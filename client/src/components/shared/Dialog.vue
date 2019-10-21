<template>
  <v-row justify="center">
    <v-dialog
      v-model="dialog"
      v-if="dialog"
      max-width="700"
      max-height="650"
    >
      <v-card id="barcodewrapper">
        <v-card-title class="headline">
            <slot name="headline"></slot>
        </v-card-title>
        <v-card-text>
          <slot name="content"></slot>
        </v-card-text>
        <v-card-actions>
          <v-layout justify-center>
            <v-btn
                color="green darken-1"
                text
                @click="dialog = false"
            >
                {{ this.closeText }}
            </v-btn>
          </v-layout>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
export default {
  props: {
    openText: {
      type: String,
      default: 'Open'
    },
    closeText: {
      type: String,
      default: 'close'
    }
  },
  data () {
    return {
      dialog: false
    }
  }
}
</script>

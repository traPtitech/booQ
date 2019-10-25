<template>
  <v-row justify="center" v-if="isTargeted">
    <v-dialog
      v-model="dialog.isOpen"
      v-if="dialog.isOpen"
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
                @click="dialog.isOpen = false"
            >
                {{ dialog.closeText }}
            </v-btn>
          </v-layout>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
export default {
  name: 'Dialog',
  props: {
    dialog: {
      type: Object,
      default: () => ({
        closeButton: 'close',
        isOpen: false,
        target: ''
      })
    },
    target: {
      type: String,
      default: ''
    }
  },
  computed: {
    isTargeted () {
      return this.$props.target === this.$props.dialog.target
    }
  }

}
</script>

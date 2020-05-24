<template>
  <v-icon>{{ svgPath }}</v-icon>
</template>

<script>
import { mdiChevronLeft, mdiChevronRight, mdiThumbUp, mdiCart, mdiMinusCircle, mdiViewDashboard, mdiHomeCity, mdiBookOpenPageVariant, mdiPlusBox, mdiHelp, mdiPlus, mdiPencil, mdiCartArrowDown } from '@mdi/js'

const mdiTable = {
  'mdi-chevron-left': mdiChevronLeft,
  'mdi-chevron-right': mdiChevronRight,
  'mdi-thumb-up': mdiThumbUp,
  'mdi-cart': mdiCart,
  'mdi-minus-circle': mdiMinusCircle,
  'mdi-view-dashboard': mdiViewDashboard,
  'mdi-home-city': mdiHomeCity,
  'mdi-book-open-page-variant': mdiBookOpenPageVariant,
  'mdi-plus-box': mdiPlusBox,
  'mdi-help': mdiHelp,
  'mdi-plus': mdiPlus,
  'mdi-pencil': mdiPencil,
  'mdi-cart-arrow-down': mdiCartArrowDown
}

export default {
  name: 'MdiIcon',
  props: {
    name: String
  },
  computed: {
    svgPath () {
      return mdiTable[this.name]
    }
  }
}
</script>

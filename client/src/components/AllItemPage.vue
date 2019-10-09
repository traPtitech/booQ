<template>
  <div v-if="data">
    <ItemList :items="data" />
  </div>
</template>

<script>
import ItemList from './shared/ItemList'
import axios from 'axios'

export default {
  name: 'AllItemPage',
  components: {
    ItemList
  },
  data () {
    return {
      data: null,
      error: null
    }
  },
  mounted () {
    axios
      .get(`/api/items`)
      .then(res => (this.data = res.data))
      .catch(e => { alert(e) })
  }
}
</script>

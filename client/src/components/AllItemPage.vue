<template>
  <ItemList :items="data" />
  <!-- <div>{{data}}</div> -->
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
  },
  methods: {
    checkRentalable (ownerID) {
      // いい感じにしてください。同じownerが複数いるときのロジックがわかりませんでした
      // 貸し出し可ならfalseを返し不可なら'ryohaが借りてます'みたいなのを返すと思ってます
      return false
    }
  }
}
</script>

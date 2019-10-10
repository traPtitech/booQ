<template>
  <div>
    <div v-if="data !== null">
      <div v-if="data.length !== 0">
        <ItemList :items="data" />
      </div>
      <div class="text-center" v-else>
        該当の物品はありません
      </div>
    </div>
    <div v-else>
      読み込み中...
    </div>
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
  watch: {
    '$route' (to, from) {
      this.mount()
    }
  },
  mounted () {
    this.mount()
  },
  methods: {
    async mount () {
      let searchParam = this.$route.query.search
      if (this.$route.query.search === undefined) {
        searchParam = ''
      }
      const res = await axios.get(`/api/items?search=` + searchParam)
        .catch(e => {
          alert(e)
          return false
        })
      this.data = res.data
    }
  }
}
</script>

<template>
  <div>
    <div v-if="data !== null">
      <div>
        <v-text-field
          class="search-input"
          label="Search..."
          color="success"
          v-model="searchString"
          v-on:keyup.enter="searchItem()"
        />
      </div>
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
  name: 'ItemListPage',
  components: {
    ItemList
  },
  data () {
    return {
      searchString: '',
      data: null,
      error: null
    }
  },
  watch: {
    '$route' (to, from) {
      switch (to.path) {
        case '/items/equipment':
          this.getItemsByType(1)
          break
        case '/items/property':
          this.getItemsByType(0)
          break
        default:
          this.search()
          break
      }
    }
  },
  mounted () {
    switch (this.$route.path) {
      case '/items/equipment':
        this.getItemsByType(1)
        break
      case '/items/property':
        this.getItemsByType(0)
        break
      default:
        this.search()
        break
    }
  },
  beforeDestroy () {
    this.$store.commit('resetNavBarTitle')
  },
  methods: {
    async search () {
      let searchParam = this.$route.query.search
      if (this.$route.query.search === undefined) {
        searchParam = ''
      }
      if (searchParam) {
        this.$store.commit('setNavBarTitle', `Search: ${searchParam}`)
      }
      const res = await axios.get('/api/items?search=' + searchParam)
        .catch(e => {
          alert(e)
        })
      this.data = res.data
      this.searchString = searchParam
    },
    async getItemsByType (type) {
      const res = await axios.get('/api/items')
        .catch(e => {
          alert(e)
        })
      const items = res.data.filter(item => {
        return item.type === type
      })
      this.data = items
    },
    searchItem () {
      this.$router.push({ path: '/items', query: { search: this.searchString } })
    }
  }
}
</script>

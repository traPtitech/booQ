<template>
  <div class="d-flex flex-wrap">
    <v-row no-gutters>
      <v-col md="auto">
        <div class="icon">
          <div>
            <h1>{{ $route.params.name }}</h1>
          </div>
          <div>
            <v-avatar size="200">
              <img :src="`https://q.trap.jp/api/1.0/public/icon/${$route.params.name}`" />
            </v-avatar>
          </div>
        </div>
      </v-col>
      <v-col cols="8">
        <div :style="`width: ${contentWidth}px;`">
          <div v-if="items !== null && comments !== null">
            <div class="content" v-if="items.length !== 0">
              <h3>所有物一覧</h3>
              <div>
                <ItemList :items="items" />
              </div>
            </div>
            <div class="content" v-else>
              <h3>このユーザーは物品を所有していません</h3>
            </div>
            <div v-if="comments.length !== 0">
              <div class="content">
                <h3>コメント一覧</h3>
                <div v-for="comment in comments" :key="comment.id">
                  <v-container class="pa-2" fluid>
                    <v-row>
                      <v-col>
                        <v-card :to="`/items/${comment.item_id}`">
                          <v-card-text>
                            <div class="headline mb-2">
                              {{comment.text}}
                            </div>
                          </v-card-text>
                        </v-card>
                      </v-col>
                    </v-row>
                  </v-container>
                </div>
              </div>
            </div>
            <div v-else>
              <h3>このユーザーはまだコメントを投稿していません</h3>
            </div>
          </div>
          <div v-else>
            読み込み中...
          </div>
        </div>
      </v-col>
    </v-row>
    <!-- <v-container>
      <ItemList :items="data" />
    </v-container> -->
  </div>
</template>

<script>
import ItemList from './shared/ItemList'
import axios from 'axios'

export default {
  name: 'UserPage',
  components: {
    ItemList
  },
  data () {
    return {
      items: null,
      comments: null,
      contentWidth: 600
    }
  },
  watch: {
    '$route' (to, from) {
      this.mount()
    }
  },
  mounted () {
    this.conputeWidth()
    window.addEventListener('resize', this.conputeWidth)
    this.mount()
  },
  beforeDestroy () {
    window.removeEventListener('resize', this.conputeWidth)
  },
  methods: {
    conputeWidth () {
      if (window.innerWidth > 991) {
        this.contentWidth = window.innerWidth - 600 // sideBar((window.innerWidth > 991で表示される)と物品のimgがともに260px
      } else if (window.innerWidth > 601) {
        this.contentWidth = window.innerWidth - 300
      } else {
        this.contentWidth = window.innerWidth - 30
      }
    },
    async mount () {
      const resItems = await axios.get('api/items?user=' + this.$route.params.name)
        .catch(e => {
          alert(e)
          return false
        })
      const resComments = await axios.get('api/comments?user=' + this.$route.params.name)
        .catch(e => {
          alert(e)
          return false
        })
      console.log(resComments)
      this.items = resItems.data
      this.comments = resComments.data
    }
  }
}
</script>

<style scoped>
  .icon {
    padding-right: 10px;
  }
  .content {
    margin-bottom: 30px;
  }
</style>

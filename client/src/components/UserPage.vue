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
          <div class="content" v-if="item.length != 0">
            <h3>所有物一覧</h3>
            <div>
              <ItemList :items="items" />
            </div>
          </div>
          <div class="content" v-else>
            <h3>あなたは物品を所有していません</h3>
          </div>
          <div v-if="comments.length == 0">
            <h3>あなたはまだコメントを投稿していません</h3>
          </div>
          <div class="content" v-else>
            <h3>コメント一覧</h3>
            <div v-for="comment in comments" :key="comment.id">
              <p>
                <v-container class="pa-2" fluid>
                  <v-row>
                    <v-col>
                      <v-card>
                        <v-card-text>
                          <div class="headline mb-2">{{comment.comment}}</div>
                        </v-card-text>
                      </v-card>
                    </v-col>
                  </v-row>
                </v-container>
              </p>
            </div>
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
      contentWidth: 600,
      sampleItems: [{
        id: 1,
        name: '小説　天気の子',
        code: 9784041026403,
        type: 1,
        owners: [
          {
            user: {
              id: 1,
              name: 'nagatech',
              displayName: 'ながてち',
              iconFileId: 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
              admin: true
            },
            rentalable: true
          }
        ],
        latest_logs: [
          {
            id: 1,
            item_id: 1,
            user: {
              id: 1,
              name: 'nagatech',
              displayName: 'ながてち',
              iconFileId: 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
              admin: true
            },
            owner: {
              id: 1,
              name: 'nagatech',
              displayName: 'ながてち',
              iconFileId: 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
              admin: true
            },
            type: 0,
            purpose: '読みたかったから。',
            due_date: '2019/07/30 23:30:00',
            created_at: '2019/07/28 22:00:00',
            updated_at: '2019/07/28 22:00:00'
          }
        ],
        like_counts: 1,
        img_url: 'https://cover.openbd.jp/9784041026403.jpg',
        created_at: '2019/07/28 22:00:00',
        updated_at: '2019/07/28 22:00:00'
      },
      {
        id: 2,
        name: '小説　天気の子',
        code: 9784041026403,
        type: 1,
        owners: [
          {
            user: {
              id: 1,
              name: 'nagatech',
              displayName: 'ながてち',
              iconFileId: 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
              admin: true
            },
            rentalable: true
          }
        ],
        latest_logs: [
          {
            id: 1,
            item_id: 1,
            user: {
              id: 1,
              name: 'nagatech',
              displayName: 'ながてち',
              iconFileId: 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
              admin: true
            },
            owner: {
              id: 1,
              name: 'nagatech',
              displayName: 'ながてち',
              iconFileId: 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
              admin: true
            },
            type: 0,
            purpose: '読みたかったから。',
            due_date: '2019/07/30 23:30:00',
            created_at: '2019/07/28 22:00:00',
            updated_at: '2019/07/28 22:00:00'
          }
        ],
        like_counts: 1,
        img_url: 'https://cover.openbd.jp/9784041026403.jpg',
        created_at: '2019/07/28 22:00:00',
        updated_at: '2019/07/28 22:00:00'
      }],
      sampleComments: [{
        id: 1,
        item_id: 1,
        user: {
          id: 1,
          name: 'nagatech',
          displayName: 'ながてち',
          iconFileId: 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
          admin: true
        },
        comment: '小説版は夏美の心理描写がよく描かれていて、映画版を補完するものになっている。あとがきと解説だけでも読む価値はあると思います。',
        created_at: '2019/07/28 22:00:00',
        updated_at: '2019/07/28 22:00:00'
      },
      {
        id: 2,
        item_id: 2,
        user: {
          id: 1,
          name: 'nagatech',
          displayName: 'ながてち',
          iconFileId: 'e0628393-8045-4c6c-b23c-6f5e6a2c252b',
          admin: true
        },
        comment: '小説版は夏美の心理描写がよく描かれていて、映画版を補完するものになっている。あとがきと解説だけでも読む価値はあると思います。',
        created_at: '2019/07/28 22:00:00',
        updated_at: '2019/07/28 22:00:00'
      }]
    }
  },
  mounted () {
    // this.items = this.sampleItems
    // this.comments = this.sampleComments
    this.conputeWidth()
    window.addEventListener('resize', this.conputeWidth)
    axios.get('api/items?user=' + this.$route.params.name).catch(e => { alert(e) }).then(res => { this.items = res.data })
    axios.get('api/comments?user=' + this.$route.params.name).catch(e => { alert(e) }).then(res => { this.comments = res.data })
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

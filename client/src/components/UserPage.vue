<template>
  <v-container>
    <v-row>
      <v-col cols="4">
        <div>
          <h1>{{ $route.params.name }}</h1>
        </div>
        <div>
          <v-avatar size="200">
            <img :src="`https://q.trap.jp/api/1.0/public/icon/${$route.params.name}`" />
          </v-avatar>
        </div>
      </v-col>
      <v-col cols="8">
        <div>
          <h3>所有物一覧</h3>
          <div>
            <ItemList :items="items" />
          </div>
        </div>
        <div>
          <h3>コメント一覧</h3>
          <div v-for="comment in comments" :key="comment.id">
            <p>
              <v-container class="pa-2" fluid>
                <v-row>
                  <v-col>
                    <v-card>
                      <v-card-text>
                        <div class="headline mb-2">{{}}</div>
                        {{comment.comment}}
                      </v-card-text>
                    </v-card>
                  </v-col>
                </v-row>
              </v-container>
            </p>
          </div>
        </div>
      </v-col>
    </v-row>
    <v-container>
      <ItemList :items="data" />
    </v-container>
  </v-container>
</template>
<script>
import ItemList from './shared/ItemList'
export default {
  name: 'UserPage',
  components: {
    ItemList
  },
  data () {
    return {
      items: null,
      comments: null,
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
    this.items = this.sampleItems
    this.comments = this.sampleComments
    // this.items = axios.get('api/items?user=$route.params.name')
    // this.comments = axios.get('api/comments?user=$route.params.name')
  }
}
</script>

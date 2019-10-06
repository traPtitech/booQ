<template>
  <v-container>
    <v-row>
      <v-col>
        <v-container>
          <h1>{{ $route.params.name }}</h1>
        </v-container>
        <v-container>
          <v-avatar size="200">
            <img :src="`https://q.trap.jp/api/1.0/public/icon/${$store.state.me.name}`" />
          </v-avatar>
        </v-container>
      </v-col>
      <v-col>
        <v-container>
          <h3>所有物一覧</h3>
          <div v-for="item in data1" :key="item.id">
            <p>
              <v-container class="pa-2" fluid>
                <v-row>
                  <v-col>
                    <v-card>
                      <v-card-text>
                        <div class="headline mb-2">{{item.name}}</div>
                      </v-card-text>
                    </v-card>
                  </v-col>
                </v-row>
              </v-container>
            </p>
          </div>
        </v-container>
        <v-container>
          <h3>コメント一覧</h3>
          <div v-for="comment in data2" :key="comment.id">
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
        </v-container>
      </v-col>
    </v-row>
    <v-container>
      <ItemList :items="data" />
    </v-container>
  </v-container>
</template>
<script>
export default {
  name: 'UserPage',
  /* data () {
    return {
      data1: null,
      data2: null
    }
  }, */
  data () {
    return {
      data1: null,
      data2: null,
      sampleData1: [{
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
      sampleData2: [{
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
    this.data1 = this.sampleData1
    this.data2 = this.sampleData2
    // this.data1 = axios.get('api/items?user=$route.params.name')
    // this.data2 = axios.get('api/comments?user=$route.params.name')
  }
}
</script>
<style>
.wrapper {
  display: flex;
  display: -ms-flexbox; /* --- IE10用 11はこの設定は不要 --- */
  display: -webkit-box; /*--- Android用 ---*/
  /*画面中央に表示されるように margin: auto;を設定している*/
  margin: auto;
  /* justify-content:stretch; */
}
/* .image {
    margin: 10px;
    -webkit-flex-basis: 30%;
    -ms-flex-basis: 30%;
    flex-basis: 30%;
    width: 30%;
    height: 50%;
  } */
.content {
  margin: 10px;
}
.comment {
  display: flex;
  display: -ms-flexbox; /*--- IE10用 11はこの設定は不要 ---*/
  display: -webkit-box; /*--- Android用 ---*/
  /*画面中央に表示されるように margin: auto;を設定している*/
  margin: auto;
}
</style>

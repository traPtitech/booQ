<template>
  <div  class="d-flex flex-wrap">
    <div>
      <v-row no-gutters>
        <v-col md="auto">
          <div class="image">
            <div>
              <img
                :src="data.img_url"
                style="width: 250px;"
              />
            </div>
            <div>
              <!-- TODO: 返す場合のボタンの処理を考える -->
              <v-btn block color="primary">借りる</v-btn>
            </div>
            <div>
              <v-btn v-if="isLiked" block @click="removeLike">
                <v-icon left color="indigo">mdi-thumb-up</v-icon>
                いいね {{ likeCount }}
              </v-btn>
              <v-btn v-else block @click="like">
                <v-icon left disabled>mdi-thumb-up</v-icon>
                いいね {{ likeCount }}
              </v-btn>
            </div>
            <div>
              <v-layout row wrap class="d-inline-flex">
                <v-flex v-for="like in data.likes" :key="like.id" >
                  <Icon :user="like" />
                </v-flex>
              </v-layout>
            </div>
          </div>
        </v-col>
      </v-row>
    </div>
    <div :style="`width: ${contentWidth}px;`">
      <h1>{{data.name}}</h1>
      <div class="content">
        {{ data.description }}
      </div>
      <div class="content">
        <h2>
          所有者
          <v-btn x-small outlined fab dark color="primary" @click="addOwner">
            <v-icon dark>mdi-plus</v-icon>
          </v-btn>
        </h2>
        <div v-for="owner in data.owners" :key="owner.id">
          <p v-if="checkRentalable(owner.user.id)">{{owner.user.name}}  {{checkRentalable(owner.user.id)}}</p>
          <p v-else v-on:click="rental">{{owner.user.name}}  貸し出し可</p>
        </div>
      </div>
      <div class="content">
        <h2>
          コメント
          <v-btn x-small outlined fab dark color="primary">
            <v-icon dark>mdi-plus</v-icon>
          </v-btn>
        </h2>
        <div>
          <div v-for="comment in data.comments" :key="comment.id">
            <v-flex>
              <Icon :user="comment.user" />
              {{ comment.comment }}
            </v-flex>
          </div>
        </div>
      </div>
      <div class="content">
        <div>
          <h2>ログ</h2>
          工事中
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Icon from './shared/Icon'
export default {
  name: 'ItemDetailPage',
  components: {
    Icon
  },
  data () {
    return {
      data: null,
      contentWidth: 600,
      sampleData: {
        id: 1,
        name: '小説　天気の子',
        code: '9784041026403',
        type: 1,
        owners: [
          {
            user: {
              id: 1,
              name: 'nagatech',
              displayName: 'ながてち',
              admin: true
            },
            rentalable: true
          }
        ],
        description: '高校1年の夏、帆高（ほだか）は離島から家出し、東京にやってきた。連日降り続ける雨の中、雑踏ひしめく都会の片隅で、帆高は不思議な能力を持つ少女・陽菜（ひな）に出会う。「ねぇ、今から晴れるよ」。それは祈るだけで、空を晴れに出来る力だった――。天候の調和が狂っていく時代に、運命に翻弄される少年と少女が自らの生き方を「選択」する物語。長編アニメーション映画『天気の子』の、新海誠監督自身が執筆した原作小説。',
        comments: [
          {
            id: 1,
            item_id: 1,
            user: {
              id: 1,
              name: 'nagatech',
              displayName: 'ながてち',
              admin: true
            },
            comment: '小説版は夏美の心理描写がよく描かれていて、映画版を補完するものになっている。あとがきと解説だけでも読む価値はあると思います。',
            created_at: '2019/07/28 22:00:00',
            updated_at: '2019/07/28 22:00:00'
          }
        ],
        logs: [
          {
            id: 1,
            item_id: 1,
            user: {
              id: 1,
              name: 'nagatech',
              displayName: 'ながてち',
              admin: true
            },
            owner: {
              id: 1,
              name: 'nagatech',
              displayName: 'ながてち',
              admin: true
            },
            type: 0,
            purpose: '読みたかったから。',
            due_date: '2019/07/30 23:30:00',
            created_at: '2019/07/28 22:00:00',
            updated_at: '2019/07/28 22:00:00'
          }
        ],
        tags: [
          {
            id: 1,
            name: '小説'
          }
        ],
        likes: [
          {
            id: 1,
            name: 'nagatech',
            displayName: 'ながてち',
            admin: true
          },
          {
            id: 2,
            name: 'ryoha',
            displayName: 'りょは',
            admin: true
          },
          {
            id: 3,
            name: 'Adwaver_4157',
            displayName: 'Waver',
            admin: true
          },
          {
            id: 4,
            name: 'series2',
            displayName: 'series2',
            admin: true
          }
        ],
        img_url: 'https://cover.openbd.jp/9784041026403.jpg',
        created_at: '2019/07/28 22:00:00',
        updated_at: '2019/07/28 22:00:00'
      }
    }
  },
  created () {
    // 本番ではaxios.getでマウントしてsampleDataを消してください
    this.data = this.sampleData
  },
  mounted () {
    this.conputeWidth()
    window.addEventListener('resize', this.conputeWidth)
  },
  beforeDestroy () {
    window.removeEventListener('resize', this.conputeWidth)
  },
  computed: {
    likeCount () {
      return this.data.likes.length
    },
    isLiked () {
      if (!this.$store.state.me) {
        return false
      }
      return this.data.likes.find(user => user.name === this.$store.state.me.name)
    }
  },
  methods: {
    conputeWidth () {
      if (window.innerWidth > 961) {
        this.contentWidth = window.innerWidth - 600
      } else if (window.innerWidth > 601) {
        this.contentWidth = window.innerWidth - 340
      } else {
        this.contentWidth = window.innerWidth
      }
    },
    checkRentalable (ownerID) {
      // いい感じにしてください。同じownerが複数いるときのロジックがわかりませんでした
      // 貸し出し可ならfalseを返し不可なら'ryohaが借りてます'みたいなのを返すと思ってます
      return false
    },
    like () {
      this.data.likes.push(this.$store.state.me)
      // TODO: axios.post(/likes)みたいな感じ
    },
    removeLike () {
      this.data.likes = this.data.likes.filter(user => user.name !== this.$store.state.me.name)
      // TODO: axios.delete(/likes)みたいな感じ
    },
    addOwner () {
      window.open('/register_owner_form', 'newwindow', 'width=400,height=800')
      // window.open('/items/' + this.data.id + '/owner/new', 'newwindow', 'width=400,height=300')
    },
    rental () {
      window.open('/rental_form', 'newwindow', 'width=400,height=800')
      // window.open('/items/' + this.data.id + '/rental', 'newwindow', 'width=400,height=300')
    }
  }
}
</script>

<style scoped>
  .image {
    padding-right: 10px;
  }
  .content {
    margin-bottom: 30px;
  }
</style>

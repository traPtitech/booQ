<template>
  <div>
    <div class="wrapper">
      <div>
        <v-row no-gutters class="mb-6">
          <v-col md="auto">
            <div class="image">
              <div>
                <img :src="data.img_url" />
              </div>
              <v-btn dark outlined round icon color="indigo" @click="like"> <v-icon dark>mdi-star</v-icon></v-btn>
              <div>
                <v-layout row wrap class="d-inline-flex">
                  <v-flex  xs4 v-for="like in data.likes" :key="like.id" >
                    <Icon :user="like" />
                  </v-flex>
                </v-layout>
              </div>
            </div>
          </v-col>
        </v-row>
      </div>
      <div>
        <div class="content">
          <h4>{{data.name}}</h4>
          <div v-for="owner in data.owners" :key="owner.user.id">
            <p v-if="checkRentalable(owner.user.id)">{{owner.user.name}}  {{checkRentalable(owner.user.id)}}</p>
            <p v-else v-on:click="clickRental">{{owner.user.name}}  貸し出し可</p>
          </div>
          <v-overlay :value="isOpenRentalForm">
            <v-btn @click="clickRental">OK</v-btn>
          </v-overlay>
          <v-btn outline round @click="clickAddOwner" color="indigo">所有者を追加</v-btn>
          <v-overlay :value="isOpenAddOwner">
            <v-btn @click="clickAddOwner">OK</v-btn>
          </v-overlay>
          <div v-for="comment in data.comments" :key="comment.id" class="comment">
            <Icon :user="comment.user" />
            <p>{{comment.comment}}</p>
          </div>
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
      isOpenAddOwner: false,
      isOpenRentalForm: false,
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
            id: 1,
            name: 'nagatech',
            displayName: 'ながてち',
            admin: true
          },
          {
            id: 1,
            name: 'nagatech',
            displayName: 'ながてち',
            admin: true
          },
          {
            id: 1,
            name: 'nagatech',
            displayName: 'ながてち',
            admin: true
          },
          {
            id: 1,
            name: 'nagatech',
            displayName: 'ながてち',
            admin: true
          },
          {
            id: 1,
            name: 'nagatech',
            displayName: 'ながてち',
            admin: true
          },
          {
            id: 1,
            name: 'nagatech',
            displayName: 'ながてち',
            admin: true
          },
          {
            id: 1,
            name: 'nagatech',
            displayName: 'ながてち',
            admin: true
          },
          {
            id: 1,
            name: 'nagatech',
            displayName: 'ながてち',
            admin: true
          },
          {
            id: 1,
            name: 'nagatech',
            displayName: 'ながてち',
            admin: true
          }
        ],
        img_url: 'https://cover.openbd.jp/9784041026403.jpg',
        created_at: '2019/07/28 22:00:00',
        updated_at: '2019/07/28 22:00:00'
      }
    }
  },
  mounted () {
    // 本番ではaxios.getでマウントしてsampleDataを消してください
    this.data = this.sampleData
  },
  methods: {
    checkRentalable (ownerID) {
      // いい感じにしてください。同じownerが複数いるときのロジックがわかりませんでした
      // 貸し出し可ならfalseを返し不可なら'ryohaが借りてます'みたいなのを返すと思ってます
      return false
    },
    like () {
      // axios.post(/likes)みたいな感じ？
    },
    clickAddOwner () {
      // window.open('/register_owner_form', 'newwindow', 'width=400,height=800')
      // window.open('/items/' + this.data.id + '/owner/new', 'newwindow', 'width=400,height=300')
      this.isOpenAddOwner = !this.isOpenAddOwner
    },
    clickRental () {
      // window.open('/rental_form', 'newwindow', 'width=400,height=800')
      // window.open('/items/' + this.data.id + '/rental', 'newwindow', 'width=400,height=300')
      this.isOpenRentalForm = !this.isOpenRentalForm
    }
  }
}
</script>

<style>
  .wrapper {
    display:flex;
    display:-ms-flexbox;/* --- IE10用 11はこの設定は不要 --- */
    display:-webkit-box;/*--- Android用 ---*/
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
    display:flex;
    display:-ms-flexbox;/*--- IE10用 11はこの設定は不要 ---*/
    display:-webkit-box;/*--- Android用 ---*/
    /*画面中央に表示されるように margin: auto;を設定している*/
    margin: auto;
  }
</style>

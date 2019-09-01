<template>
  <div>
    <div class="wrapper">
      <div class="image">
        <div>
          <img :src="data.img_url" />
        </div>
        <button v-on:click="like">いいね</button>
        <div v-for="like in data.likes" :key="like.id">
          <v-avatar size="40">
            <img :src="'https://q.trap.jp/api/1.0/files/' + like.iconFileId" />
          </v-avatar>
        </div>
      </div>
      <div class="content">
        <h4>{{data.name}}</h4>
        <div v-for="owner in data.owners" :key="owner.user.id">
          <p v-if="checkRentalable(owner.user.id)">{{owner.user.name}}  {{checkRentalable(owner.user.id)}}</p>
          <p v-else v-on:click="clickRental">{{owner.user.name}}  貸し出し可</p>
        </div>
        <button v-on:click="clickAddOwner">所有者を追加</button>
        <div v-for="comment in data.comments" :key="comment.id" class="comment">
          <v-avatar size="40">
            <img :src="'https://q.trap.jp/api/1.0/files/' + comment.user.iconFileId" />
          </v-avatar>
          <p>{{comment.comment}}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ItemDetailPage',
  data() {
    return{
      data: null,
      sampleData: {
        id: 1,
        name: "小説　天気の子",
        code: "9784041026403",
        type: 1,
        owners: [
            {
            user: {
                id: 1,
                name: "nagatech",
                displayName: "ながてち",
                iconFileId: "e0628393-8045-4c6c-b23c-6f5e6a2c252b",
                admin: true
            },
            rentalable: true
            }
        ],
        description: "高校1年の夏、帆高（ほだか）は離島から家出し、東京にやってきた。連日降り続ける雨の中、雑踏ひしめく都会の片隅で、帆高は不思議な能力を持つ少女・陽菜（ひな）に出会う。「ねぇ、今から晴れるよ」。それは祈るだけで、空を晴れに出来る力だった――。天候の調和が狂っていく時代に、運命に翻弄される少年と少女が自らの生き方を「選択」する物語。長編アニメーション映画『天気の子』の、新海誠監督自身が執筆した原作小説。",
        comments: [
            {
            id: 1,
            item_id: 1,
            user: {
                id: 1,
                name: "nagatech",
                displayName: "ながてち",
                iconFileId: "e0628393-8045-4c6c-b23c-6f5e6a2c252b",
                admin: true
            },
            comment: "小説版は夏美の心理描写がよく描かれていて、映画版を補完するものになっている。あとがきと解説だけでも読む価値はあると思います。",
            created_at: "2019/07/28 22:00:00",
            updated_at: "2019/07/28 22:00:00"
            }
        ],
        logs: [
            {
            id: 1,
            item_id: 1,
            user: {
                id: 1,
                name: "nagatech",
                displayName: "ながてち",
                iconFileId: "e0628393-8045-4c6c-b23c-6f5e6a2c252b",
                admin: true
            },
            owner: {
                id: 1,
                name: "nagatech",
                displayName: "ながてち",
                iconFileId: "e0628393-8045-4c6c-b23c-6f5e6a2c252b",
                admin: true
            },
            type: 0,
            purpose: "読みたかったから。",
            due_date: "2019/07/30 23:30:00",
            created_at: "2019/07/28 22:00:00",
            updated_at: "2019/07/28 22:00:00"
            }
        ],
        tags: [
            {
            id: 1,
            name: "小説"
            }
        ],
        likes: [
            {
            id: 1,
            name: "nagatech",
            displayName: "ながてち",
            iconFileId: "e0628393-8045-4c6c-b23c-6f5e6a2c252b",
            admin: true
            }
        ],
        img_url: "https://cover.openbd.jp/9784041026403.jpg",
        created_at: "2019/07/28 22:00:00",
        updated_at: "2019/07/28 22:00:00"
      }
    }
  },
  mounted() {
    // 本番ではaxios.getでマウントしてsampleDataを消してください
    this.data = this.sampleData
  },
  methods: {
    checkRentalable(ownerID) {
      // いい感じにしてください。同じownerが複数いるときのロジックがわかりませんでした
      // 貸し出し可ならfalseを返し不可なら'ryohaが借りてます'みたいなのを返すと思ってます
      return false
    },
    like() {
      // axios.post(/likes)みたいな感じ？
    },
    clickAddOwner() {
      window.open('/register_owner_form', 'newwindow', 'width=400,height=300')
    },
    clickRental() {
      window.open('/rental_form', 'newwindow', 'width=400,height=300')
    },
  }
}
</script>

<style>
    .wrapper {
        display:flex;
        display:-ms-flexbox;/*--- IE10用 11はこの設定は不要 ---*/
        display:-webkit-box;/*--- Android用 ---*/
        /*画面中央に表示されるように margin: auto;を設定している*/
        margin: auto;
    }
    .comment {
        display:flex;
        display:-ms-flexbox;/*--- IE10用 11はこの設定は不要 ---*/
        display:-webkit-box;/*--- Android用 ---*/
        /*画面中央に表示されるように margin: auto;を設定している*/
        margin: auto;
    }
</style>

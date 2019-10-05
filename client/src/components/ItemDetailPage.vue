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
          <div class="text-center">
            <v-dialog light v-model="isOpenRentalForm" max-width="320">
              <v-card width="320">
                <v-card-title class="headline">物品を借りる</v-card-title>
                <v-card-actions>
                  <v-menu bottom origin="center center" transition="scale-transition" open-on-hover>
                    <template v-slot:activator="{ on }">
                      <v-btn color="primary" dark v-on="on">所有者を選ぶ</v-btn>
                    </template>
                    <v-list>
                      <v-list-item 
                      v-for="(owner, i) in data.owners" 
                      :key="i" 
                      @click="rentOwnerID=owner.user.id" 
                      :disabled="!owner.rentalable">
                        <v-list-item-title>{{ owner.user.name }}</v-list-item-title>
                      </v-list-item>
                    </v-list>
                  </v-menu>
                </v-card-actions>
                <v-card-actions>
                  <div>
                    <v-form ref="form">
                      <v-textarea outlined v-model="purpose" label="目的"/>
                    </v-form>
                  </div>
                </v-card-actions>
                <v-card-actions>
                  <div>
                    <v-form ref="form">
                      <v-text-field outlined v-model.number="rentalCount" type="number" label="個数"/>
                    </v-form>
                  </div>
                </v-card-actions>
                <v-card-actions max-width="320">
                  <v-date-picker v-model="dueDate"></v-date-picker>
                </v-card-actions>
                <v-alert type="error" v-if="rentalCount<0">個数が負になっています</v-alert>
                <v-divider></v-divider>
                <v-card-actions>
                  <div class="flex-grow-1"></div>
                  <v-btn v-on:click="rental()">借りる</v-btn>
                </v-card-actions>
              </v-card>
            </v-dialog>
          </div>
          <v-btn outline round @click="clickAddOwner" color="indigo">所有者を追加</v-btn>
          <div class="text-center">
            <v-dialog light v-model="isOpenAddOwner" max-width="290">
              <v-card width="290">
                <v-card-title class="headline">所有者を追加する</v-card-title>
                <v-card-actions>
                  <div v-if="$store.state.me.admin" >
                    <label v-for="(label,id) in ownerOptions" v-bind:key="label">
                      <div><input type="radio" name="owner" :value="id" v-model="ownerID">{{ label }}</div>
                    </label>
                  </div>
                </v-card-actions>
                <v-card-actions>
                  <div>
                    <input type="checkbox" id="checkbox" v-model="rentalable">
                    <label for="checkbox">貸し出し可</label>
                  </div>
                </v-card-actions>
                <v-card-actions>
                  <div>
                    <v-form ref="form">
                      <v-text-field outlined v-model.number="count" type="number" label="個数"/>
                    </v-form>
                  </div>
                </v-card-actions>
                <v-alert type="error" v-if="count<0">個数が負になっています</v-alert>
                <v-divider></v-divider>
                <v-card-actions>
                  <div class="flex-grow-1"></div>
                  <v-btn v-on:click="add()">追加</v-btn>
                </v-card-actions>
              </v-card>
            </v-dialog>
          </div>
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
import axios from 'axios'

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
      ownerID: 0,
      ownerOptions: {
        0: '自身',
        1: 'traP',
        2: '支援課'
      },
      rentalable: true,
      count: 1,
      purpose: '',
      rentalCount: 1,
      dueDate: '',
      rentOwnerID: 0,
      // 以下はサンプルデータ
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
          },
          {
            user: {
              id: 1,
              name: 'nagatech',
              displayName: 'ながてち',
              admin: true
            },
            rentalable: true
          },
          {
            user: {
              id: 1,
              name: 'nagatech',
              displayName: 'ながてち',
              admin: true
            },
            rentalable: false
          },
          {
            user: {
              id: 1,
              name: 'nagatech',
              displayName: 'ながてち',
              admin: true
            },
            rentalable: false
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
      this.isOpenAddOwner = !this.isOpenAddOwner
    },
    clickRental () {
      this.isOpenRentalForm = !this.isOpenRentalForm
    },
    async add () {
      if (this.ownerID === 0) {
        const res = await axios.post(`/api/items/` + this.$route.params.id + `/owners`, { user_id: this.$store.state.me.ID, rentalable: this.rentalable, count: this.count }).catch(e => { alert(e) })
        alert('”' + this.data.name + '”の所有者に' + this.$store.state.me.name + 'を追加しました。')
      } else {
        const res = await axios.post(`/api/items/` + this.$route.params.id + `/owners`, { user_id: this.ownerID, rentalable: this.rentalable, count: this.count }).catch(e => { alert(e) })
        alert('”' + this.data.name + '”の所有者に' + this.ownerOptions[this.ownerID] + 'を追加しました。')
      }
      this.isOpenAddOwner = !this.isOpenAddOwner
    },
    async rental () { 
      const res = await axios.post(`/api/items/` + this.$route.params.id + `/logs`, { owner_id: this.rentOwnerID, type: 0, purpose: this.purpose, due_date: this.dueDate, count: this.rentalCount }).catch(e => { alert(e) })
      alert('あなたは”' + this.data.name + '”を' + this.rentalCount + '個借りました。' + this.dueDate)
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

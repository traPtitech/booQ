<template>
  <div  class="d-flex flex-wrap">
    <div>
      <v-row no-gutters>
        <v-col md="auto">
          <div class="image">
            <div>
              <img
                :src="data.img_url === '' ? 'https://q.trap.jp/api/1.0/files/3380fbc6-6141-4b60-99ae-a1d270842d60/thumbnail' : data.img_url"
                style="width: 250px;"
              />
            </div>
            <div>
              <RentalForm @add="reload" :data="data"/>
              <ReturnForm @returnItem="reload" :data="data"/>
              <v-btn block color="warning">返す</v-btn>
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
                  <Icon
                    :user="like"
                    :size="25"
                  />
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
          <RegisterOwnerForm/>
        </h2>
        <!-- FIXME: 他のタスクに手をつけたかったので表示が適当です -->
        <div v-for="owner in data.owners" :key="owner.id">
          <Icon
            :user="owner.user"
            :size="25"
          />
          {{ owner.user.name }} {{ checkRentalable(owner) }}
        </div>
      </div>
      <div class="content">
        <h2>
          コメント
          <CommentDialog/>
        </h2>
        <div>
          <div v-for="comment in data.comments" :key="comment.id">
            <v-flex>
              <Icon :user="comment.user" />
              {{ comment.text }}
            </v-flex>
          </div>
        </div>
      </div>
      <div class="content">
        <div>
          <h2>ログ</h2>
          <div v-for="log in reverseLogs" :key="log.id">
            <Icon
              :user="log.user"
              :size="25"
            />
            {{ createLogMessage(log) }} - {{ log.created_at }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Icon from './shared/Icon'
import axios from 'axios'
import RegisterOwnerForm from './shared/RegisterOwnerForm'
import RentalForm from './shared/RentalForm'
import CommentDialog from './shared/CommentDialog'
import ReturnForm from './shared/ReturnForm'

export default {
  name: 'ItemDetailPage',
  components: {
    Icon,
    RegisterOwnerForm,
    RentalForm,
    CommentDialog,
    ReturnForm
  },
  data () {
    return {
      data: null,
      contentWidth: 600
    }
  },
  created () {
    axios
      .get(`/api/items/` + this.$route.params.id)
      .then(res => (this.data = res.data))
      .catch(e => { alert(e) })
  },
  mounted () {
    this.conputeWidth()
    window.addEventListener('resize', this.conputeWidth)
  },
  beforeDestroy () {
    window.removeEventListener('resize', this.conputeWidth)
  },
  computed: {
    reverseLogs () {
      return this.data.logs.slice().reverse()
    },
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
      if (window.innerWidth > 991) {
        this.contentWidth = window.innerWidth - 600 // sideBar((window.innerWidth > 991で表示される)と物品のimgがともに260px
      } else if (window.innerWidth > 601) {
        this.contentWidth = window.innerWidth - 300
      } else {
        this.contentWidth = window.innerWidth - 30
      }
    },
    checkRentalable (owner) {
      // FIXME: ロジックがやばい
      if (!owner.rentalable) {
        return '貸し出しできません'
      }
      var latestLog = this.data.latest_logs.filter(function (log) {
        return (log.owner.ID = owner.owner_id)
      })
      var rentalableCount = 0
      if (latestLog === [] || !latestLog[0].count) {
        rentalableCount = owner.count
      } else {
        rentalableCount = latestLog[0].count
      }
      if (rentalableCount === 0) {
        return '貸し出しできません'
      } else if (rentalableCount === 1) {
        return '貸し出し可能'
      }
      return '貸し出し可能' + '×' + rentalableCount
      // return '貸し出し可能' + '×' + 1
    },
    createLogMessage (log) {
      const userName = log.user.name
      const ownerName = log.owner.name
      const ownerWord = ownerName === 'traP' ? '' : `${ownerName}さんの`
      let logComment = log.type === 0 ? '借りました' : '返しました'
      logComment = log.type === 2 ? '追加しました' : logComment
      return `${userName}さんが${ownerWord}物品を${logComment}`
    },
    checkLogType (log) {
      if (log.type === 0) {
        return '借りました'
      } else {
        return '返しました'
      }
    },
    async like () {
      var postLikeError = null
      await axios.post(`/api/items/` + this.$route.params.id + `/likes`, null)
        .catch(e => {
          alert(e)
          postLikeError = e
        })
      if (!postLikeError) {
        this.data.likes.push(this.$store.state.me)
      }
    },
    async removeLike () {
      var removeLikeError = null
      await axios.delete(`/api/items/` + this.$route.params.id + `/likes`, null)
        .catch(e => {
          alert(e)
          removeLikeError = e
        })
      if (!removeLikeError) {
        this.data.likes = this.data.likes.filter(user => user.name !== this.$store.state.me.name)
      }
    },
    async reload () {
      const res = await axios.get(`/api/items/` + this.$route.params.id).catch(e => { alert(e) })
      this.data = res.data
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

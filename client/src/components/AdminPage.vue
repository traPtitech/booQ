<template>
  <div>
    <h4>ユーザーに権限を付与します</h4>
    <div v-if="users.length">
      <div v-if="$store.state.me && $store.state.me.admin">
        <div>
          <v-btn @click="changeUserInfo">
            確定
          </v-btn>
        </div>
        <div v-for="user in users" :key="user.id">
          <v-checkbox
            v-model="user.admin"
            :label="user.name"
            hide-details
          />
        </div>
        <div>
          <v-btn @click="changeUserInfo">
            確定
          </v-btn>
        </div>
      </div>
      <div v-else>
        権限がありません
      </div>
    </div>
    <div v-else>
      読み込み中...
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'AdminPage',
  data () {
    return {
      users: [],
      error: ''
    }
  },
  async mounted () {
    const res = await axios.get('/api/users').catch(e => alert(e))
    this.users = res.data
  },
  methods: {
    async changeUserInfo () {
      // いまのUserの情報と比較して変更があるやつだけapiを飛ばす
      const res = await axios.get('/api/users').catch(e => alert(e))
      const latestUsers = res.data
      const changeUsers = []
      latestUsers.forEach(latestUser => {
        const targetUser = this.users.find(user => {
          return user.name === latestUser.name
        })
        if (targetUser.admin !== latestUser.admin) {
          changeUsers.push(targetUser)
        }
      })
      if (changeUsers.length === 0) {
        alert('変更がありません')
        return
      }
      await Promise.all(changeUsers.map(user => {
        return axios.put('/api/users', user)
      })).catch(e => {
        alert(e)
        this.error = e
      })
      if (this.error) {
        alert('問題が発生したため、処理が正常に行われませんでした。')
      } else {
        alert('正常に変更されました。')
      }
    }
  }
}
</script>

<template>
  <nobr>
    <v-btn x-small outlined fab dark color="primary" @click.stop="open">
      <v-icon dark>mdi-plus</v-icon>
    </v-btn>
    <div class="text-center">
      <v-dialog light v-model="isOpenCommentDialog" max-width="290">
        <v-card width="290">
          <v-card-title class="headline">コメントを追加する</v-card-title>
          <v-card-actions>
            <div>
              <v-form ref="form">
                <v-textarea outlined v-model="text" :rules="[() => !!text || 'This field is required']" label="コメント"/>
              </v-form>
            </div>
          </v-card-actions>
          <v-divider></v-divider>
          <v-card-actions>
            <div class="flex-grow-1"></div>
            <v-btn v-on:click="comment()">投稿</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
  </nobr>
</template>

<script>
import axios from 'axios'

export default {
  name: 'CommentDialog',
  data () {
    return {
      text: null,
      error: '',
      isOpenCommentDialog: false,
      message: ''
    }
  },
  methods: {
    async comment () {
      if (this.text === null) {
        alert('コメントを入力してください')
        return false
      }
      await axios.post(`/api/items/` + this.$route.params.id + `/comments`, { text: this.text })
        .catch(e => {
          alert(e)
          this.error = e
        })
      if (!this.error) {
        this.$parent.data.comments.push({ user: this.$store.state.me, text: this.text })
      }
      this.isOpenCommentDialog = !this.isOpenCommentDialog
    },
    open () {
      this.isOpenCommentDialog = !this.isOpenCommentDialog
    }
  }
}
</script>

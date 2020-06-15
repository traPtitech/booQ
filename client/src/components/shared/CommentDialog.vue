<template>
  <nobr>
    <v-btn x-small outlined fab dark color="primary" @click.stop="open">
      <mdi-icon dark name="mdi-plus" />
    </v-btn>
    <div class="text-center">
      <v-dialog light max-width=90% max-elevation=80% v-model="isOpenCommentDialog" >
        <v-card width=100%>
          <v-card-title class="headline">コメントを追加する</v-card-title>
          <v-card-actions width=100%>
              <v-form>
                <v-textarea
                  outlined v-model="text" :rules="[() => !!text || 'This field is required']"
                  label="コメント"
                  cols="150"
                  rows="10"
                  style="width:100%; height:60%;"
                ></v-textarea>
              </v-form>
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
import MdiIcon from '../shared/MdiIcon.vue'

export default {
  name: 'CommentDialog',
  components: {
    MdiIcon
  },
  data () {
    return {
      text: null,
      error: '',
      isOpenCommentDialog: false,
      message: ''
    }
  },
  props: {
    propItem: Object
  },
  methods: {
    async comment () {
      if (this.text === null) {
        alert('コメントを入力してください')
        return false
      }
      await axios.post(`/api/items/${this.$route.params.id}/comments`, { text: this.text })
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

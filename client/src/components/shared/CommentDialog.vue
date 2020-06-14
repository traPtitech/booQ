<script type="size">
export default {
  data: function () {
    return {
      width: window.innerWidth,
      height: window.innerHeight
    }
  },
  methods: {
    handleResize: function () {
      this.width = window.innerWidth
      this.height = window.innerHeight
    }
  },
  mounted: function () {
    window.addEventListener('resize', this.handleResize)
  },
  beforeDestroy: function () {
    window.removeEventListener('resize', this.handleResize)
  }
}
</script>

<template>
  <nobr>
    <v-btn x-small outlined fab dark color="primary" @click.stop="open">
      <mdi-icon dark name="mdi-plus" />
    </v-btn>
    <div class="text-center">
      <v-dialog light max-width=size.width*0.8 max-elevation=size.height*0.9 v-model="isOpenCommentDialog" >
        <v-card width=size.width*0.8 elevation=size.height*0.9>
          <v-card-title class="headline">コメントを追加する</v-card-title>
          <v-card-actions width=size.width*0.7 elevation=size.height*0.7>
            <div>
              <v-form>
                <v-textarea
                 outlined v-model="text" :rules="[() => !!text || 'This field is required']"
                 label="コメント"
                 width=size.width*0.7
                 height=size.height*0.7
                ></v-textarea>
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

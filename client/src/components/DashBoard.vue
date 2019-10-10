<template>
  <div>
    <h1>booQへようこそ！</h1>
    <div>
      booQは超すごい備品管理アプリケーションです。<br>
      ここのページは本来ダッシュボードでいい感じな表示になる予定ですが、工事中です。<br>
      かわりに説明文載せておくので読んでください。<br><br>
    </div>
    <About />
  </div>
</template>

<script>
import About from '@/components/About.vue'

export default {
  name: 'DashBoard',
  components: {
    About
  }
}
</script>

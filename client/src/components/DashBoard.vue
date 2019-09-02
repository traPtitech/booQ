<template>
  <div>
    <h1>DashBoard Page</h1>
    <!-- aタグでリンクしてるの頭悪そう(Vuexのいいところ？の読み込みが入らないところがなくなっているため) -->
    <div><a href="/items/new">物品の新規登録</a></div>
    <!-- あとで↓けす -->
    <div><a href="/items_test">物品の詳細(テスト)</a></div>
    <div><a href="/items">すべての物品</a></div>
    <div><a href="/admin">AdminPage</a></div>
  </div>
</template>

<script>
export default {
  name: 'DashBoard'
}
</script>

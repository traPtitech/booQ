<template>
  <div>
    <h1>User Page</h1>
    <h2>{{ $route.params.name }}</h2>
    {{ $store.state.me }}
  </div>
</template>

<script>
export default {
  name: 'UserPage'
}
</script>

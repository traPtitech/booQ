<template>
  <div>
    <h4>借りる</h4>
    <p>目的</p><input v-model="porpose">
    <!-- ↓はいい感じのpicker使ったらいいと思います -->
    <p>返却日</p><input v-model="dueDate">
    <button v-on:click="submit">submit</button>
  </div>
</template>

<script>
export default {
  name: 'RentalForm',
  data() {
    return{
      porpose: '',
      dueDate: ''
    }
  },
  methods: {
    submit() {
      // this.porposeとかをaxios.postで
    }
  }
}
</script>

<template>
  <v-container>
    <h4>借りる</h4>
    <p>目的</p><input v-model="porpose">
    <!-- ↓はいい感じのpicker使ったらいいと思います -->
    <p>返却日</p>
    <v-row justify="space-around">
      <v-select class="px2" solo type="date" v-model="dueDate"/>
      <!-- <v-date-picker color="indigo" width="290" :day-format="date => new Date(date).getDate()" show-current v-model="dueDate"></v-date-picker> -->
    </v-row>
    <div>{{dueDate}}</div>
    <v-btn round outlined color="indigo" @click="submit">submit</v-btn>
  </v-container>
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

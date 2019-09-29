<template>
  <div>
    <v-card
      class="mx-auto"
      width="1500"
      elevation="5"
      tile
    >
      <v-list
        two-line
        avatar
        nav
      >
        <v-list-item-group v-model="items" color="primary">
          <v-list-item
            v-for="item in items"
            :key="item.id"
            :to="`/items/${item.id}`"
            style="height: 100px;"
          >
            <img
              :src="item.img_url"
              class="item-list-image"
            />
            <v-list-item-content style="padding-left: 15px;">
              <v-list-item-title class="headline mb-1">{{ item.name }}</v-list-item-title>
              <v-list-item-subtitle>{{ item.owners.map(i => i.user.name).join(', ') }}</v-list-item-subtitle>
            </v-list-item-content>
             <v-list-item-icon v-if="item.like_counts">
              <v-icon>thumb_up_alt</v-icon>
              {{ item.like_counts }}
            </v-list-item-icon>
          </v-list-item>
        </v-list-item-group>
      </v-list>
    </v-card>
  </div>
</template>

<script>
export default {
  name: 'ItemList',
  props: [
    'items'
  ]
}
</script>

<style>
.item-list-image {
  max-height: 100px;
  max-width: 120px;
}
</style>

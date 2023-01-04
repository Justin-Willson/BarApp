<script lang="ts">
import { defineComponent } from 'vue'
import type { Ingredient } from '@/domain/ingredient';
import type { PropType } from 'vue'

export default defineComponent({
  props: {
    ingredient: {
      type: Object as PropType<Ingredient>,
        required: true
      }
  },
  mounted() {
    this.ingredient
  }
})
</script>

<template>  
  <b-row class="justify-content-md-center" :class="ingredient.is_in_stock ? 'inStock' : 'outOfStock' ">
    <b-col col lg="4">{{ ingredient.name }}</b-col>
    <b-col cols="12" md="auto">{{ ingredient.notes }}</b-col>
    <b-col col lg="2">{{ ingredient.is_in_stock }}</b-col>
  </b-row>
</template>

<style scoped>
.inStock {
  color: green
}
.outOfStock {
  color: red
}
</style>
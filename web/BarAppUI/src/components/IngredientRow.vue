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
  <tr class="ingredient_row" :class="ingredient.is_in_stock ? 'inStock' : 'outOfStock' ">
    <td col lg="4">{{ ingredient.name }}</td>
    <td cols="12" md="auto">{{ ingredient.notes }}</td>
    <td col lg="2">{{ ingredient.is_in_stock }}</td>
  </tr>
</template>

<style scoped>
.ingredient_row {
  outline: black;
}
.inStock {
  color: green
}
.outOfStock {
  color: red
}
</style>
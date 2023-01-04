<script lang="ts">
import { defineComponent } from 'vue'
import type { Ingredient } from '@/domain/ingredient';
import type { PropType } from 'vue'

export default defineComponent({
  props: {
    ingredients: {
      type: Object as PropType<Ingredient[]>,
        required: true
      }
  },
  mounted() {
    this.ingredients
  }
})

</script>

<template>
    <b-container class="bv-example-row">
        <IngredientRow v-for="ingredient in ingredients" :ingredient=ingredient  />
    </b-container>
</template>

<style scoped>

</style>
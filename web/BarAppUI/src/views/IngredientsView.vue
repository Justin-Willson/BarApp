<script lang="ts">
import IngredientTable from '../components/IngredientTable.vue';
import { Ingredient } from '@/domain/ingredient';

export default {
  data () {
    return {
      ingredients: [] as Ingredient[],
    }
  },
  async mounted (): Promise<void> {
    const ingredients = await Ingredient.getIngredients()
    this.ingredients = ingredients
  },
}
</script>

<template>
  <main>
   <IngredientTable :ingredients=ingredients></IngredientTable>
  </main>
</template>
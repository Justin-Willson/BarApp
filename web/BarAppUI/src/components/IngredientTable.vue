<script lang="ts">
import { defineComponent } from 'vue'
import type { Ingredient } from '@/domain/ingredient';
import type { PropType } from 'vue'
import AddIngredientRow from './AddIngredientRow.vue';

export default defineComponent({
    props: {
        ingredients: {
            type: Object as PropType<Ingredient[]>,
            required: true
        }
    },
    mounted() {
        this.ingredients;
    },
    components: { AddIngredientRow }
})

</script>

<template>
  
    <table class="table">
      <thead class="thead-dark">
        <th >Name</th>
        <th >Notes</th>
        <th >In Stock?</th>
      </thead>
      <tbody>
        <IngredientRow v-for="ingredient in ingredients" :ingredient=ingredient  />        
      </tbody>
    </table>
    <AddIngredientRow />
</template>

<style scoped>

</style>
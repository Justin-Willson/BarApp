<script setup lang="ts">
import IngredientRow from '../components/IngredientRow.vue';
import { Ingredient } from '@/domain/ingredient';

var ingredient = new Ingredient("Vodka", true, "A very plain thing")
</script>

<template>
  <main>
    <IngredientRow :ingredient=ingredient  />
  </main>
</template>

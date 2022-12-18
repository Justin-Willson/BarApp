<script setup lang="ts">
import type { Ingredient } from '@/domain/ingredient';

defineProps<{
  ingredient: Ingredient
}>()
</script>

<template>
  <b-container class="bv-example-row">
  <b-row class="justify-content-md-center">
    <b-col col lg="2">{{ ingredient.name }}</b-col>
    <b-col cols="12" md="auto">{{ ingredient.notes }}</b-col>
    <b-col col lg="2">{{ ingredient.is_in_stock }}</b-col>
  </b-row>
</b-container>
</template>

<style scoped>

</style>
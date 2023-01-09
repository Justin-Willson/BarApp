<script lang="ts">
import { defineComponent } from 'vue'
import { Ingredient } from '@/domain/ingredient';
import type { PropType } from 'vue'
import { createStructuralDirectiveTransform } from '@vue/compiler-core';

export default defineComponent({
  data () {
    return {
      name: "",
      notes: "",
      is_in_stock: false,
    }
  },
  methods: {
    async add_ingredient(e: Event) {
      e.preventDefault()
      let ingredient = new Ingredient(this.name, this.is_in_stock, this.notes)
      let response = await ingredient.postIngredient()
      alert(response.name)
      this.clearForm()

    },
    clearForm() {
      this.name =""
      this.notes = ""
      this.is_in_stock = false
    }
  }
})
</script>

<template>
  <form v-on:submit="add_ingredient" >
  <div class="row">
    <div class="col-lg-4">
      <input type="text" class="form-control" placeholder="Ingredient" v-model="name">
    </div>
    <div class="col-lg-12">
      <input type="text" class="form-control" placeholder="Notes" v-model="notes">
    </div>
    <div class="col-lg-1">
      <input type="checkbox" class="form-check-input" id="checkbox" v-model="is_in_stock" />
    </div>
    <div class="col-lg-1">
      <button type="submit" class="btn btn-primary">Add</button>
    </div>
  </div>
</form>

  <span>{{is_in_stock}} {{name}} {{notes}}</span>
</template>

<style scoped>
</style>
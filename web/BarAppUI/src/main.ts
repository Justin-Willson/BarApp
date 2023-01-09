import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import './assets/main.css'
import "bootstrap"
import 'bootstrap/dist/css/bootstrap.css'

import IngredientTableVue from './components/IngredientTable.vue'
import IngredientRowVue from './components/IngredientRow.vue'

const app = createApp(App)

app.use(router)

app.component('IngredientTable', IngredientTableVue)
app.component('IngredientRow', IngredientRowVue)

app.mount('#app')

import { createApp, h } from 'vue'
import ElementPlus from 'element-plus'

import { router } from './router'
import { RouterView } from 'vue-router'

import 'minireset.css'
import 'element-plus/lib/theme-chalk/index.css'
import './assets/styles.css'

const app = createApp({
  setup() {
    return () => h(RouterView)
  }
})

app.use(ElementPlus)
app.use(router)
app.mount('#app')

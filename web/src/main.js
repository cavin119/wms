import Vue from 'vue'
import App from '@/App.vue'

import ElementUI  from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
Vue.use(ElementUI)

import router from '@/router/index'
// time line css
import '../node_modules/timeline-vuejs/dist/timeline-vuejs.css'

import '@/permission'

import { store } from '@/store/index'
Vue.config.productionTip = false

// 路由守卫
import VueBus from '@/utils/bus'
Vue.use(VueBus)

new Vue({
  render: h => h(App),
  router,
  store
}).$mount('#app')


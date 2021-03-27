import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import { Button, Form, FormItem, Select, Option, Input } from 'element-ui'
Vue.use(Button)
Vue.use(Form)
Vue.use(FormItem)
Vue.use(Select)
Vue.use(Option)
Vue.use(Input)

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')

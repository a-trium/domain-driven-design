import Vue from 'vue'

import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

import App from './App.vue'
import Router from 'vue-router'
import Routes from './router'
import store from './store'
import './registerServiceWorker'


Vue.use(ElementUI)

const router = new Router({
    routes: Routes,
})

Vue.config.productionTip = false

new Vue({
    router,
    store,
    render: (h) => h(App),
}).$mount('#app')

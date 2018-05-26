import Vue from 'vue'
import Router from 'vue-router'

import Home from './views/Home.vue'
import About from './views/About.vue'
import Login from './views/Login.vue'
import Register from './views/Register.vue'

Vue.use(Router)

export default new Router({
    routes: [
        {
            path: '/',
            name: 'home',
            component: Home,
        },
        {
            path: '/',
            name: 'register',
            component: Register,
        },
        {
            path: '/',
            name: 'login',
            component: Login,
        },
        {
            path: '/about',
            name: 'about',
            component: About,
        },
    ],
})

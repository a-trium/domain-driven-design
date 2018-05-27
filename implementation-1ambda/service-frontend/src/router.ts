import Vue from 'vue'
import VueRouter from 'vue-router'

import Home from './views/Home.vue'
import About from './views/About.vue'
import Login from './views/Login.vue'
import Register from './views/Register.vue'

Vue.use(VueRouter)

export default [
    {
        path: '/',
        name: 'home',
        displayName: '7 Street',
        component: Home,
    },
    {
        path: '/register',
        name: 'register',
        displayName: 'Register',
        component: Register,
    },
    {
        path: '/login',
        name: 'login',
        displayName: 'Login',
        component: Login,
    },
    {
        path: '/about',
        name: 'about',
        displayName: 'About',
        component: About,
    },
]


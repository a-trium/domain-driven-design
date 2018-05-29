import Vue from 'vue'
import Router from 'vue-router'

import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

import App from '@/App.vue'
import routes from "@/router.ts"
import store from '@/store'
import '@/registerServiceWorker'
import { AuthAPI } from "@/common/auth.service.ts"
import { Exception } from "@/generated/swagger"

Vue.use(ElementUI)

const router = new Router({
    routes: routes,
})

AuthAPI.whoami({ credentials: 'include' })
    .then((response) => {

        router.beforeEach((to: any, from: any, next: any) => {
            // if the page doesn't require authentication, move to the page
            if (!to.matched.some((record: any) => record.meta.requiresAuth)) {
                return next()
            }

            // check user is authenticated
            const uid = store.state.uid
            if (!uid) {
                if (to.name == 'login') {
                    return next('login')
                }

                // if not, redirect to the login page with flash message

                store.commit('setFlashMessage', `Please Login for '${to.path}'`)
                return next('login')
            }

            next()
        })

        if (!response.uid || response.uid.trim() == '') {
            store.commit('logout')
            router.push('/login')
            return
        }

        store.commit('login', response.uid)
    })
    .catch((response) => {
        response.json().then((parsed: Exception) => {
            router.push('/login')
            console.log(parsed)
        })
    })


Vue.config.productionTip = false

new Vue({
    router,
    store,
    render: (h) => h(App),
}).$mount('#app')

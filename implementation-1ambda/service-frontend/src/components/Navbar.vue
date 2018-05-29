<template>
    <el-menu :router="true"
             :default-active="$route.path"
             mode="horizontal"
             background-color="#545c64"
             text-color="#fff"
             active-text-color="#ffd04b">
        <template v-for='item in routes'>
            <el-menu-item :index="item.path" :route='item'
                          v-if="shouldDisplay(item.meta.requiresAuth, authenticated, item.meta.common)">
                {{item.displayName}}
            </el-menu-item>
        </template>
        <el-submenu :index="'nav-dropdown'" v-if="authenticated" style="float: right;">
            <template slot="title"><i class="el-icon-menu"></i><span>{{uid}}</span></template>
            <el-menu-item index="'nav-dropdown-setting'">
                <i class="el-icon-setting"></i>
                <span>Settings</span>
            </el-menu-item>
            <el-menu-item index="'nav-dropdown-logout'" v-on:click="logout">
                <i class="el-icon-circle-close"></i>
                <span>Logout</span>
            </el-menu-item>
        </el-submenu>
        <el-menu-item index="github" v-else style="float: right;">
            <a href="https://github.com/1ambda" target="_blank"style="display: block; text-decoration: none;">
                <i class="el-icon-location-outline"></i>
                <span>Github</span>
            </a>
        </el-menu-item>
    </el-menu>
</template>

<script lang="ts">
    import { Component, Vue } from 'vue-property-decorator'
    import { mapActions, mapGetters, mapMutations, mapState } from 'vuex'

    import Router from '@/router.ts'
    import { AuthAPI } from '@/common/auth.service.ts'
    import { Exception } from '@/generated/swagger'

    @Component({
        components: {},
        computed: {
            ...mapState([ 'uid', 'flashMessage', ]),
            ...mapGetters([ 'authenticated', ]),
        },
        watch: {
            flashMessage(newMessage, oldMessage) {
                if (!newMessage) {
                    return
                }

                this.$notify.error({
                    title: `Error`,
                    message: newMessage,
                })

                this.$store.commit('cleanFlashMessage')
            }
        },
    })
    export default class Navbar extends Vue {
        routes = Router
        $notify: any
        $route: any
        $router: any
        $store: any

        shouldDisplay(requiresAuth: boolean, authenticated: boolean, common: boolean) {
            if (common) {
                return true
            } else if (requiresAuth && authenticated) {
                return true
            } else if (!requiresAuth && !authenticated) {
                return true
            }

            return false
        }

        logout() {
            AuthAPI.logout({}, { credentials: 'include', })
                .then((response) => {
                    this.$notify({
                        title: 'Success',
                        message: 'Logged out',
                        type: 'success',
                    })

                    this.$store.commit('logout')
                    this.$router.push('/login')
                })
                .catch((response) => {
                    response.json().then((parsed: Exception) => {
                        this.$notify.error({
                            title: `Error (${parsed.type})`,
                            message: parsed.message,
                        })
                    })
                })
        }
    }
</script>
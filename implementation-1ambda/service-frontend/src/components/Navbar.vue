<template>
    <el-menu :default-active="$route.path"
             :router="true"
             mode="horizontal"
             background-color="#545c64"
             text-color="#fff"
             active-text-color="#ffd04b">
        <template v-for='item in routes'>
            <el-menu-item :index="item.path" :route='item'>
                {{item.displayName}}
            </el-menu-item>
        </template>
    </el-menu>
</template>

<script>
    import {Component, Vue} from 'vue-property-decorator'
    import Router from "../router"

    @Component({
        components: {},
    })
    export default class Navbar extends Vue {
        data() {
            return {
                routes: Router,
            }
        }
    }
</script>
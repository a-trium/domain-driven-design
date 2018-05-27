<template>
    <div>
        <div style="margin: 20px;"></div>
        <el-form :label-position="'right'" :rules="rules" label-width="100px" :model="ruleForm" ref="ruleForm">
            <el-form-item label="ID" prop="uid">
                <el-input v-model="ruleForm.uid"></el-input>
            </el-form-item>
            <el-form-item label="Password">
                <el-input type="password" v-model="ruleForm.password" auto-complete="off"></el-input>
            </el-form-item>

            <el-form-item>
                <el-button type="primary" @click="submitForm('ruleForm')">Login</el-button>
            </el-form-item>
        </el-form>
    </div>

</template>

<script lang="ts">
    import { Component, Vue } from 'vue-property-decorator'
    import { Exception, LoginRequest } from "../generated/swagger"
    import { AuthAPI } from "../common/api"

    @Component({
        components: {},
    })
    export default class Login extends Vue {
        $refs: any
        $notify: any

        private ruleForm = {
            uid: '',
            password: '',
        }

        private rules = {
            uid: [
                { required: true, message: 'Please input id', trigger: 'blur' },
                { min: 4, max: 30, message: 'Length should be 4 to 30', trigger: 'blur' },
                { pattern: /^([a-zA-Z0-9]+)$/, message: 'Please use alpha numeric only', trigger: 'blur' }, ],
            password: [
                { required: true, message: 'Please input password', trigger: 'blur' }, ]
        }

        submitForm(formName: string) {
            this.$refs[ formName ].validate((valid: any) => {
                if (!valid) {
                    return
                }

                const request: LoginRequest = {
                    uid: this.ruleForm.uid,
                    password: this.ruleForm.password,
                }

                AuthAPI.login(request, { credentials: 'include' })
                    .then((response) => {
                        console.log(response)
                    })
                    .catch((response) => {
                        response.json().then((parsed: Exception) => {
                            this.$notify.error({
                                title: `Error (${parsed.type})`,
                                message: parsed.message,
                            })
                        })
                    })
            })
        }
    }
</script>

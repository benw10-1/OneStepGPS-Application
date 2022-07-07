<template>
    <div class="login-page">
        <div class="login-text">
            <h1>Track your fleets</h1>
            <p>Use the One Step GPS public API to display your data</p>
        </div>
        <div class="login-body">
            <div class="login-form">
                <div class="login-form-header">
                    <h1 v-if="switch_">Sign up</h1>
                    <h1 v-else>Sign in</h1>
                </div>
                <div class="login-form-body">
                    <div v-if="keySwitch" class="api-form-input">
                        <input type="text" placeholder="API Key" v-model="APIKey" v-on:keyup="(e) => {setAPIKey(e.currentTarget.value)}" />
                    </div>
                    <div v-else class="login-form-input">
                        <input type="text" placeholder="John Doe" v-model="Name" v-on:keyup="(e) => {setName(e.currentTarget.value)}" />
                        <input type="password" placeholder="Password" v-model="Password" v-on:keyup="(e) => {setPass(e.currentTarget.value)}" />
                    </div>
                    <div class="login-form-tip" v-if="!keySwitch">
                        <span v-if="switch_" v-on:click="other">Already have an account?</span>
                        <span v-else v-on:click="other">Don't have an account?</span>
                    </div>
                    <div class="login-form-button">
                        <button @click="login">
                            <span v-if="keySwitch">
                                Submit
                            </span>
                            <span v-else-if="switch_">
                                Sign up
                            </span>
                            <span v-else>
                                Sign in
                            </span>
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { Requests } from '../helpers'

export default {
    name: 'LogIn',
    props: {
    },
    data() {
        return {
            switch_: false,
            keySwitch: false,
            Name: '',
            Password: '',
            APIKey: '',
            Error: '',
        }
    },
    methods: {
        async login() {
            if (this.Name && this.Password) {
                let data = this.switch_ ? await Requests.signup(this.Name, this.Password) : await Requests.login(this.Name, this.Password)
                if (!data?.Token) {
                    this.Error = "Invalid username or password"
                    return
                }

                if (!data.APIKey) this.keySwitch = true
            }
        },
        other() {
            this.switch_ = !this.switch_
            this.Password = ''
            this.keySwitch = false
        },
        setName(name) {
            const cleaned = name.replace(/[^a-zA-Z0-9]/g, '')
            this.Name = cleaned
        },
        setPass(pass) {
            const cleaned = pass.replace(/[^a-zA-Z0-9!@#$%&]/g, '')
            this.Password = cleaned
        },
        setAPIKey(key) {
            const cleaned = key.trim()
            this.APIKey = cleaned
        }
    },
}
</script>

<style scoped>
.login-page {
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
}

.login-text {

}
</style>

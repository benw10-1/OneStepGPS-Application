<template>
    <div class="login-page">
        <div class="login-text">
            <h1>Track your fleets</h1>
            <p>Use the One Step GPS public API to display your data</p>
        </div>
        <div class="login-body">
            <div class="login-inner">
                <div class="login-form-header">
                    <h1>{{ switch_ ? "Sign up" : "Log in" }}</h1>
                </div>
                <div class="login-form-body" v-on:keydown="(e) => { if (e.key === 'Enter') login() }">
                    <div v-if="keySwitch" class="login-form-input">
                        <input type="text" placeholder="API Key" v-model="APIKey"
                            v-on:input="(e) => { setAPIKey(e.currentTarget.value) }" />
                    </div>
                    <div v-else class="login-form-input">
                        <input type="text" placeholder="JohnDoe77" v-model="Name"
                            v-on:input="(e) => { setName(e.currentTarget.value) }" />
                        <input type="password" placeholder="Password" v-model="Password"
                            v-on:input="(e) => { setPass(e.currentTarget.value) }" />
                    </div>
                    <div class="login-form-submit" v-if="!keySwitch">
                        <span v-on:click="other" class="login-form-switch">
                            {{ switch_ ? "Already have an account?" : "Don't have an account?" }}
                        </span>
                        <button @click="login">
                            {{ keySwitch ? "Submit" : (switch_ ? "Sign up" : "Log in") }}
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
            collapsed: window.innerWidth < 975,
            small: window.innerWidth < 450
        }
    },
    mounted() {
        window.addEventListener('resize', this.resizer)
    },
    unmounted() {
        window.removeEventListener('resize', this.resizer)
    },
    methods: {
        async login() {
            if (this.Name && this.Password) {
                this.Error = ''
                let data = this.keySwitch ?
                    await Requests.setAPIKey(this.APIKey) :
                    (this.switch_ ?
                        await Requests.signup(this.Name, this.Password) :
                        await Requests.login(this.Name, this.Password))
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
        },
        resizer() {
            this.collapsed = window.innerWidth < 975
            this.small = window.innerWidth < 450
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
    flex-direction: v-bind("collapsed ? 'column' : 'row'");
}

.login-text {
    margin-right: v-bind("collapsed ? '0' : '102px'");
    margin-bottom: v-bind("collapsed ? '102px' : '0'");
}

.login-text h1,
.login-form-header h1 {
    margin: 0;
    font-size: 2.125rem;
    font-weight: 400;
    line-height: 1.235;
    letter-spacing: .0125em;
    color: rgb(25, 118, 210);
}

.login-text p {
    margin: 0;
    font-size: 1rem;
    font-weight: 400;
    line-height: 1.75;
    letter-spacing: .0125em;
    color: rgba(0, 0, 0, 0.87)
}

.login-body {
    box-shadow: rgb(0 0 0 / 20%) 0px 2px 1px -1px, rgb(0 0 0 / 14%) 0px 1px 1px 0px, rgb(0 0 0 / 12%) 0px 1px 3px 0px;
    width: v-bind("small ? '100%' : '400px'");
    background-color: white;
    border-radius: 4px;
    position: relative;
    height: 370px;
    padding: 24px;
    display: flex;
    box-sizing: border-box;
}
.login-inner {
    flex: 1;
    display: flex;
    flex-direction: column;
}
.login-form-body {
    display: flex;
    flex-direction: column;
    flex: 1
}
.login-form-input {
    display: flex;
    flex-direction: column;
    flex: 1;
    justify-content: center;
}
</style>

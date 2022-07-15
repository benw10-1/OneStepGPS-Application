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
                        <CustomInput placeholder="376529419055430ab4f3056b01e61fc6" label="API Key" :value="APIKey"
                            :onChange="setAPIKey" :required="true" flexItem :error="Error" />
                    </div>
                    <div v-else class="login-form-input">
                        <CustomInput placeholder="JohnDoe77" label="Username" :value="Name" :onChange="setName"
                            :required="true" flexItem :error="Error" />
                        <CustomInput placeholder="Password" label="Password" :value="Password" :onChange="setPass"
                            :required="true" type="password" flexItem :error="switch_ ? '' : Error" />
                    </div>
                    <div class="login-form-submit">
                        <span v-on:click="other" class="login-form-switch" v-if="!keySwitch">
                            {{ switch_ ? "Already have an account?" : "Don't have an account?" }}
                        </span>
                        <CustomButton padding="10px 14px" width="100px"
                            :value="keySwitch ? 'SUBMIT' : (switch_ ? 'SIGNUP' : 'LOGIN')" :onClick="login" />
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
document.title = "Login"
</script>

<script>
import { Requests } from '../helpers'
import CustomInput from "../components/CustomInput.vue"
import CustomButton from "../components/CustomButton.vue"

export default {
    name: 'LogIn',
    components: {
        CustomInput,
        CustomButton
    },
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
                // reset error
                this.Error = ''
                let data = this.keySwitch ?
                    await Requests.setAPIKey(this.APIKey) :
                    (this.switch_ ?
                        await Requests.signup(this.Name, this.Password) :
                        await Requests.login(this.Name, this.Password))
                // rerender with error if there is one
                if (!data?.Token) {
                    this.Error = this.keySwitch ? "API Key invalid" : this.switch_ ? "Username already taken" : "Invalid username or password"
                    return
                }
                // need to set API key
                if (!data.APIKey) this.keySwitch = true
                if (data.error) {
                    this.Error = data.error
                    return
                }
            }
        },
        other() {
            // switch between login and signup
            this.switch_ = !this.switch_
            this.Password = ''
            this.Error = ''
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
        // trigger rerender on resize
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
    margin: v-bind("collapsed ? '102px 16px' : '0 102px 0 0'");
    text-align: v-bind("collapsed ? 'center' : 'left'");
}

.login-text h1,
.login-form-header h1 {
    margin: 0;
    font-size: 2.125rem;
    font-weight: 400;
    line-height: 1.235;
    letter-spacing: .0125em;
    color: rgb(25, 118, 210);
    text-align: v-bind("collapsed ? 'center' : 'left'");
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
    height: 315px;
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

.login-form-submit {
    display: flex;
    flex-direction: row;
    justify-content: v-bind("collapsed ? 'space-around' : 'flex-end'");
    align-items: center;
}

.login-form-switch {
    margin-right: v-bind("collapsed ? '0' : '16px'");
    font-size: 1rem;
    font-weight: 400;
    /* line-height: 1.75; */
    letter-spacing: .0125em;
    color: #3366BB;
    text-decoration: underline;
    cursor: pointer;
    user-select: none;
}
</style>

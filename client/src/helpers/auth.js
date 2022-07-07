import decode from "jwt-decode"

export default {
    getProfile() {
        const tok = this.getToken()
        if (!tok) return null

        return decode(tok).data
    },
    loggedIn() {
        const token = this.getToken()
        // If there is a token and it's not expired, return `true`
        return token && !this.isTokenExpired(token)
    },
    isTokenExpired(token) {
        // Decode the token to get its expiration time that was set by the server
        const decoded = decode(token)
        // If the expiration time is less than the current time (in seconds), the token is expired and we return `true`
        if (decoded.exp < Date.now() / 1000) {
            this.logout()
            return true
        }
        // If token hasn't passed its expiration time, return `false`
        return false
    },
    getToken() {
        let item = localStorage.getItem('id_token')
        return item === "undefined" ? undefined : item
    },
    login(idToken) {
        localStorage.setItem('id_token', idToken)
        if (this.onLogin) this.onLogin()
    },
    logout() {
        localStorage.removeItem('id_token')
        if (this.onLogout) this.onLogout()
    },
    onLogin(callback) {
        this.onLogin = callback
    },
    onLogout(callback) {
        this.onLogout = callback
    }
}

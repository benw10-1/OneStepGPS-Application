import { reactive } from "vue";

export default reactive({
    data: [],
    async setData(data) {
        this.data = data
        console.log(this.data)
    },
    log() {
        console.log(this.data)
    }
});
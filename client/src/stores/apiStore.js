import { reactive } from "vue";

export default reactive({
    data: [],
    async setData(data) {
        this.data = data
    }
});
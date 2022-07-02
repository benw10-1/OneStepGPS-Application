import { reactive } from "vue";
import { Requests } from "@/helpers";

export default reactive({
    data: {},
    async refresh() {
        this.data = await Requests.getInfo();
    }
});
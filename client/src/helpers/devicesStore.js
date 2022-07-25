import { defineStore } from "pinia";
import Requests from "./requests";

export default defineStore("devicesStore", {
    state: () => ({
        devices: [],
        interval: null
    }),
    actions: {
        async refreshDevices() {
            try {
                const data = await Requests.getDevices();
                this.devices = data?.result_list ?? data ?? [];
            }
            catch (err) {
                console.log(err);
                return
            }
        },
        startInterval() {
            this.interval = setInterval(this.refreshDevices, 10000);
        },
        stopInterval() {
            clearInterval(this.interval);
        }
    },
});
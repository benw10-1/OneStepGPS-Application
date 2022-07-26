import { defineStore } from "pinia";
import Requests from "./requests";

export default defineStore("settingsStore", {
    state: () => ({
        settings: {},
    }),
    getters: {
        deviceSettings(state) {
            return (id) => {
                return state.settings.deviceSettings[id] ?? {
                    icon: 'car',
                    displayName: null,
                    visible: true,
                };
            };
        },
        allDeviceSettings(state) {
            return state.settings.deviceSettings;
        },
        filterSettings(state) {
            return state.settings.filter;
        },
        sortSettings(state) {
            return state.settings.sort;
        },
        mapSettings(state) {
            return state.settings.mapSettings ?? {
                showLabels: true,
                cluster: true,
                mapDisplay: 'STREET',
            };
        }
    },
    actions: {
        async refreshSettings() {
            try {
                const data = await Requests.getPreferences();
                this.settings = data ?? {};
            }
            catch (err) {
                console.log(err);
                return
            }
        },
        async saveSettings() {
            try {
                await Requests.setPreferences(this.settings);
            }
            catch (err) {
                console.log(err);
                return
            }
        },
        updateDeviceSettings(device_id, data) {
            const toUpdate = this.settings.deviceSettings ?? {}

            toUpdate[device_id] = {
                ...toUpdate[device_id],
                ...data
            }

            this.settings.deviceSettings = toUpdate;
            this.saveSettings();
        },
        updateFilter(data) {
            this.settings.filter = {
                ...this.settings.filter ?? {},
                ...data
            }
            this.saveSettings();
        },
        updateSort(data) {
            this.settings.sort = data
            this.saveSettings();
        },
        updateMapSettings(data) {
            this.settings.mapSettings = {
                ...this.settings.mapSettings ?? {},
                ...data
            }
            this.saveSettings();
        }
    },
});
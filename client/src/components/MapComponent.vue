<template>
    <div class="map-container">
        <Transition>
            <div class="settings-panel" v-if="showSettings">
                <div class="setting-header">
                    <span>Map Settings</span>
                    <span class="close material-icons-outlined" @click="hideSettingsPanel">
                        close
                    </span>
                </div>
                <div class="settings-container">
                    <SwitchSetting :initial="settings.mapDisplay === 'STREET' ? 1 : 0" text="Map Display"
                        :switch-array="['SATELITE', 'STREET']" :onChange="(change) => {
                            updateSettings({
                                mapDisplay: change
                            })
                        }" />
                </div>
                <div class="settings-container">
                    <SwitchSetting :initial="settings.showLabels ? 0 : 1" text="Labels" :switch-array="['SHOW', 'HIDE']"
                        :onChange="(change) => { updateSettings({ showLabels: change.toLowerCase() === 'show' }) }" />
                </div>
                <div class="settings-container">
                    <SwitchSetting :initial="settings.cluster ? 0 : 1" text="Cluster"
                        :switch-array="['CLUSTER', 'NONE']"
                        :onChange="(change) => { updateSettings({ cluster: change?.toLowerCase() === 'cluster' }) }" />
                </div>
                <div class="setting-header">
                    <span>User Settings</span>
                </div>
                <div class="settings-container">
                    <SwitchSetting text="Logout" :switch-array="['logout']" :onChange="(change) => {
                        logout()
                    }" :is-icon="true" />
                </div>
            </div>
        </Transition>
        <div id="map-target"></div>
    </div>
</template>

<script>
import { PreferenceHolder, Auth, deviceStore } from '../helpers'
import { mapState } from 'pinia';
import "leaflet/dist/leaflet.css";
import L from "leaflet"
import SwitchSetting from './SwitchSetting.vue';
import "leaflet.markercluster/dist/MarkerCluster.css";
import "leaflet.markercluster/dist/MarkerCluster.Default.css";
import "leaflet.markercluster/dist/leaflet.markercluster-src.js";
// import "leaflet-rotatedmarker"

export default {
    name: 'MapComponent',
    components: {
        SwitchSetting
    },
    // add our rerender triggers to our data updater and initalialize map and resize trigger
    mounted() {
        PreferenceHolder.onUpdate(this.updateSettings)
        this.initMap()
        window.addEventListener('resize', this.resizer)
    },
    // remove our rerender triggers from our data updater and remove map and resize trigger
    unmounted() {
        PreferenceHolder.removeUpdate(this.updateSettings)
        window.removeEventListener('resize', this.resizer)
        if (this.map) {
            this.map.remove()
        }
    },
    computed: {
        ...mapState(deviceStore, ['devices']),
    },
    data() {
        return {
            search: '',
            // filtered devices to display
            display: [],
            map: null,
            fitDevices: false,
            featureLayer: null,
            loaded: false,
            tileLayer: null,
            showSettings: false,
            geoJSON: null,
            // colors based on device status
            colors: {
                off: "rgba(255, 0, 0, .67)",
                driving: "rgba(0, 255, 0, .67)",
                idle: "rgba(135, 206, 250, .67)",
                offline: "rgba(218, 223, 225, 1)",
            },
            settings: PreferenceHolder.get().mapSettings ?? {
                showLabels: true,
                cluster: true,
                mapDisplay: 'STREET',
            },
            selected: null,
        }
    },
    methods: {
        select(device) {
            this.selected = device
            if (device) this.moveCenter([device.lat, device.lng])
            this.updateDisplay(this.devices)
        },
        deselect() {
            this.selected = null
        },
        showSettingsPanel() {
            this.showSettings = true
            window.addEventListener('mousedown', this.hideSettingsPanel)
        },
        hideSettingsPanel(e) {
            if (!e.currentTarget.classList?.contains('close')) {
                for (const x of e.composedPath()) {
                    if (x.classList?.contains("settings-panel")) {
                        return
                    }
                }
            }
            window.removeEventListener('mousedown', this.hideSettingsPanel)
            this.showSettings = false
        },
        logout() {
            Auth.logout()
        },
        changeCluster(cluster, map = this.map) {
            if (this.featureLayer) map.removeLayer(this.featureLayer)

            this.featureLayer = L.markerClusterGroup({
                showCoverageOnHover: false,
                spiderfyOnMaxZoom: false,
                zoomToBoundsOnClick: true,
                maxClusterRadius: 52,
                disableClusteringAtZoom: cluster ? 19 : false,
            })
            this.featureLayer.on('clusterclick', (e) => {
                // on click tell parent we have clicked on a cluster
                this.$emit('cluster-click', e.layer.getAllChildMarkers())
            })
            this.featureLayer.addTo(map)
        },
        // main update method
        updateDisplay(data) {
            if (!data || !data.length) {
                this.fitDevices = true
                return
            }
            const prefs = PreferenceHolder.get()?.deviceSettings ?? {}
            this.display = data.filter((d) => {
                if (prefs[d.device_id]) {
                    return prefs[d.device_id].visible === undefined || prefs[d.device_id].visible
                } else {
                    return true
                }
            })
            // if map is loaded, update map
            if (this.map && this.loaded) {
                if (this.fitDevices) {
                    this.fitDevices = false
                    // resize map after loading
                    this.map.invalidateSize()
                    // center map on a box that contains all devices
                    if (data.length) this.map.fitBounds(this.getDevicesView())
                }

                if (this.display.length && this.featureLayer) {
                    this.featureLayer.clearLayers()
                    // remap all of our features to new devices
                    this.geoJSON.addData(this.display.map((device, i) => {
                        // console.log(device)
                        return {
                            type: 'Feature',
                            // "Point" for geojson will be converted to marker
                            geometry: {
                                type: 'Point',
                                coordinates: [device.lng, device.lat],
                            },
                            // props passed to marker so that we can access them later
                            properties: {
                                device_id: device.device_id ?? i,
                                display_name: prefs[device.device_id]?.displayName ?? device.display_name,
                                drive_status: device.drive_status,
                                drive_status_begin_time: device.drive_status_begin_time,
                                heading: device.heading,
                                online: device.online,
                                lat: device.lat,
                                lng: device.lng,
                                color: device.online ? this.colors[device.drive_status] : this.colors['offline'],
                                speed: device.speed,
                                selected: this.selected?.device_id === device.device_id,
                            }
                        }
                    }))
                }
            }
        },
        updateSettings(settings) {
            let safe
            if (settings.isPrefs__) {
                settings = settings.mapSettings ?? {}
            }
            else {
                safe = true
            }
            if (settings.cluster !== undefined) this.changeCluster(settings.cluster)
            if (settings.mapDisplay) {
                this.changeMapType(settings.mapDisplay?.toLowerCase())
            }
            this.settings = {
                ...this.settings,
                ...settings,
            }
            if (safe) PreferenceHolder.set({ mapSettings: this.settings }, this.updateSettings)
            this.updateDisplay(this.devices)
        },
        changeMapType(type, map = this.map) {
            type = type?.toLowerCase()

            if (!this.tileLayer) {
                this.tileLayer = L.tileLayer(null, {
                    attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                }).addTo(map)
            }
            if (type === "satelite") {
                this.tileLayer.setUrl("https://server.arcgisonline.com/ArcGIS/rest/services/World_Imagery/MapServer/tile/{z}/{y}/{x}")
            }
            else {
                this.tileLayer.setUrl("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png")
            }
        },
        initMap() {
            // view controls custom
            const map = L.map('map-target', {
                // attributionControl: false,
                zoomControl: false,
                // zoomAnimation:false,
                fadeAnimation: true,
                markerZoomAnimation: true
                // renderer: L.canvas(),
            });

            this.changeMapType(this.settings.mapDisplay ?? "street", map)
            this.changeCluster(this.settings.cluster, map)

            this.geoJSON = L.geoJSON(null, {
                // on point create. Returns marker to be created with geoJSON data
                pointToLayer: (feature, latlng) => {
                    // the icon for our marker
                    const angle = feature.properties.heading
                    const svgSt = `transform-box: fill-box; transform-origin: center; transform: rotate(${angle}deg); height: 35px; position: relative;`
                    const [short, long] = ["8px", "14px"]
                    const transformLabel = `transform: translate(${(() => {
                        if (angle < 45) {
                            return short
                        }
                        if (angle < 135) {
                            return long
                        }
                        return short
                    })()
                        }, -50%);`

                    const svgIcon = L.divIcon({
                        className: 'device-icon',
                        // the physical HTML element being disaplyed at point [lat, lon]
                        html: `<div class="device-icon-inner" id=${'id' + feature.properties.device_id.replace(/[^\w\d]/gi, "")}>
                            <div style="${svgSt}">
                                ${feature.properties.selected ? `
                                <div class="device-icon-selected"></div>
                                ` : ``}
                                <svg
                                    width="24"
                                    height="35"
                                    viewBox="-26.5 -29 53 58"
                                    preserveAspectRatio="none"
                                    xmlns="http://www.w3.org/2000/svg"
                                >
                                    <path 
                                        shape-rendering="optimizeQuality" 
                                        d="M 25 27.5 L 0 17.5 L -25 27.5 L 0 -27.5 Z" 
                                        stroke="#000000" 
                                        stroke-width="1" 
                                        fill="${feature.properties.color}"
                                    />
                                </svg>
                            </div>
                            ${this.settings.showLabels ? `
                            <div class="device-label-anchor" style="${transformLabel}">
                                <div class="device-label">
                                    <div class="label-content">
                                        ${feature.properties.display_name + (feature.properties.speed ? ` (${feature.properties.speed} mph)` : ``)}
                                    </div>
                                    <div class="tail-shadow"></div>
                                    <div class="tail-layer"></div>
                                    <div class="tail-background"></div>
                                </div>
                            </div>` : ''}
                        </div>`,
                        iconSize: [40, 40],
                        iconAnchor: [20, 20],
                        tooltipAnchor: [8, 0],
                    })

                    const marker = L.marker(latlng, {
                        icon: svgIcon,
                        // imported with leaflet-rotatedmarker
                        // rotationAngle: feature.properties.heading,
                    })

                    marker.on("click", () => {
                        this.select(feature.properties)
                    })

                    return marker
                },
                // on feature create
                onEachFeature: (feature, layer) => {
                    // geoJSON data is passed to the display layer
                    this.featureLayer.addLayer(layer)

                    layer.on('click', () => {
                        // on click tell parent we have clicked on a device
                        this.$emit('device-click', feature.properties)
                    })
                }
            })

            map.setMaxZoom(18);
            map.setMinZoom(3);
            // load trigger
            map.on("load", () => {
                setTimeout(() => {
                    // update map size
                    map.invalidateSize()
                    this.loaded = true
                    this.$emit("map-loaded", {
                        moveCenter: this.moveCenter,
                        select: this.select
                    })
                    this.updateDisplay(this.devices)
                }, 50)
                this.map = map
            })
            map.on("click", () => {
                this.select(null)
            })
            const preventActions = function (el) {
                L.DomEvent.on(el, 'click', L.DomEvent.stopPropagation);
                L.DomEvent.on(el, 'click', L.DomEvent.preventDefault);
                L.DomEvent.on(el, 'dblclick', L.DomEvent.stopPropagation);
                L.DomEvent.on(el, 'dblclick', L.DomEvent.preventDefault);
                L.DomEvent.on(el, 'mouseover', () => {
                    map.dragging.disable()
                });
                L.DomEvent.on(el, 'mouseout', () => {
                    map.dragging.enable()
                });
            }
            // add view controls
            const parent = this
            L.Control.ViewControls = L.Control.extend({
                onAdd: function () {
                    const container = L.DomUtil.create('div', 'view-controller');

                    const zoomInButton = L.DomUtil.create('button', 'view-controller-button', container);
                    const zoomIn = L.DomUtil.create('span', 'material-icons-outlined', zoomInButton);
                    zoomIn.innerHTML = 'add';

                    const zoomOutButton = L.DomUtil.create('button', 'view-controller-button', container);
                    const zoomOut = L.DomUtil.create('span', 'material-icons-outlined', zoomOutButton);
                    zoomOut.innerHTML = 'remove';

                    const zoomOutMapButton = L.DomUtil.create('button', 'view-controller-button', container);
                    const zoomOutMap = L.DomUtil.create('span', 'material-icons-outlined', zoomOutMapButton);
                    zoomOutMap.innerHTML = 'zoom_out_map';

                    preventActions(container)

                    preventActions(zoomInButton)
                    L.DomEvent.on(zoomInButton, 'click', (e) => {
                        e.currentTarget.blur();
                        map.zoomIn();
                    });

                    preventActions(zoomOutButton)
                    L.DomEvent.on(zoomOutButton, 'click', (e) => {
                        e.currentTarget.blur();
                        map.zoomOut();
                    });
                    preventActions(zoomOutMapButton)
                    L.DomEvent.on(zoomOutMapButton, 'click', (e) => {
                        e.currentTarget.blur();
                        map.fitBounds(parent.getDevicesView());
                    });

                    return container;
                }
            })
            // factory for view controls
            L.control.viewControls = function (options) {
                return new L.Control.ViewControls(options);
            }

            L.Control.Settings = L.Control.extend({
                onAdd: function () {
                    const container = L.DomUtil.create('div', 'view-controller');

                    const settingsButton = L.DomUtil.create('button', 'view-controller-button', container);
                    const settings = L.DomUtil.create('span', 'material-icons', settingsButton);
                    settings.innerHTML = 'settings';

                    preventActions(container)
                    preventActions(settingsButton)
                    L.DomEvent.on(settingsButton, 'click', () => {
                        parent.showSettingsPanel()
                    })

                    return container;
                }
            })
            // factory function settings button
            L.control.settings = function (options) {
                return new L.Control.Settings(options);
            }

            L.control.settings({ position: 'topright' }).addTo(map);
            L.control.viewControls({ position: 'topleft' }).addTo(map);
            // arbitrary default zoom level
            map.setView([0, 0], 13)
        },
        moveCenter(lat, lng) {
            this.map.panTo(L.latLng(lat, lng), 12)
        },
        // return bounding box of all viewable devices
        getDevicesView() {
            if (!this.devices.length) return null

            const { left, top, bottom, right } = this.display.reduce((acc, device) => {
                const { lat, lng } = device
                // convert to pixel coordinates for actual comparison
                const point = this.map.latLngToContainerPoint([lat, lng])
                const { x, y } = point
                // update bounding box
                acc.left = Math.min(acc.left, x)
                acc.top = Math.min(acc.top, y)
                acc.right = Math.max(acc.right, x)
                acc.bottom = Math.max(acc.bottom, y)

                return acc
            }, { left: Infinity, top: Infinity, right: -Infinity, bottom: -Infinity })
            // contruct latlng bounding box from pixel bounding box
            const bounds = L.latLngBounds([
                this.map.containerPointToLatLng([left, top]),
                this.map.containerPointToLatLng([right, bottom])
            ])

            const viewPadding = 0.06
            bounds.pad(viewPadding)

            return bounds
        },
        resizer() {
            this.map.invalidateSize()
            this.$forceUpdate()
        },
    },
    watch: {
        devices: {
            handler(newDevices) {
                this.updateDisplay(newDevices)
            },
        },
    },
}
</script>

<style>
.map-container {
    flex: 1;
    height: 100%;
    box-sizing: border-box;
    position: relative;
}

#map-target {
    width: 100%;
    height: 100%;
}

#map-target:focus {
    outline: none;
}

.view-controller {
    z-index: 1;
    display: flex;
    justify-content: center;
    flex-direction: column;
    align-items: center;
    padding: 0;
    background-color: white;
    box-shadow: 0 3px 1px -2px rgba(0, 0, 0, .2), 0 2px 2px 0 rgba(0, 0, 0, .14), 0 1px 5px 0 rgba(0, 0, 0, .12);
}

.view-controller-button {
    padding: 0.34rem;
    width: 35px;
    cursor: pointer;
    outline: none;
    border: none;
    /* width: 30px; */
    background-color: transparent;
    color: #808080;
    display: grid;
    place-items: center;
    transition: background-color 0.2s ease-in-out;
    user-select: none;
}

.view-controller-button:hover {
    background-color: #f5f5f5;
}

.view-controller-button:active {
    background-color: #e0e0e0;
}

.device-icon-inner {
    position: relative;
    display: grid;
    place-items: center;
    width: 100%;
    height: 100%;
}

.device-label-anchor {
    position: absolute;
    top: 50%;
    left: 100%;
    z-index: 11;
}

.device-label {
    background-color: #ffffff;
    position: relative;
    border-radius: 4px;
    -moz-border-radius: 4px;
    -webkit-border-radius: 4px;
    box-shadow: 0px 0px 8px -1px black;
    -moz-box-shadow: 0px 0px 8px -1px black;
    -webkit-box-shadow: 0px 0px 8px -1px black;
}

.label-content {
    padding: 0.2rem;
    font-size: 0.8rem;
    font-weight: bold;
    text-align: center;
    color: #000000;
    white-space: nowrap;
}

.popup-content {
    padding: .5rem;
}

.tail-shadow {
    background-color: transparent;
    width: 4px;
    height: 4px;
    position: absolute;
    top: 16px;
    left: -8px;
    z-index: -10;
    box-shadow: 0px 0px 8px 1px black;
    -moz-box-shadow: 0px 0px 8px 1px black;
    -webkit-box-shadow: 0px 0px 8px 1px black;
}

.tail-shadow-popup {
    background-color: transparent;
    width: 4px;
    height: 4px;
    position: absolute;
    bottom: -16px;
    left: 50%;
    transform: translateX(-50%);
    z-index: -10;
    box-shadow: 0px 0px 8px 1px black;
    -moz-box-shadow: 0px 0px 8px 1px black;
    -webkit-box-shadow: 0px 0px 8px 1px black;
}

.tail-layer {
    width: 0px;
    height: 0px;
    border: 10px solid;
    border-color: transparent #e0e0e0 transparent transparent;
    position: absolute;
    top: 50%;
    transform: translateY(-50%);
    left: -20px;
}

.tail-layer-popup {
    width: 0px;
    height: 0px;
    border: 10px solid;
    border-color: transparent transparent transparent #e0e0e0;
    position: absolute;
    bottom: -20px;
    left: 50%;
    transform: translateX(-50%);
}

.tail-background {
    width: 0px;
    height: 0px;
    border: 10px solid;
    border-color: transparent #ffffff transparent transparent;
    position: absolute;
    left: -18px;
    top: 50%;
    transform: translateY(-50%);
}

.tail-background-popup {
    width: 0px;
    height: 0px;
    border: 10px solid;
    border-color: transparent #ffffff transparent transparent;
    position: absolute;
    bottom: -18px;
    left: 50%;
    transform: translateX(-50%);
}

.leaflet-marker-icon {
    display: grid;
    place-items: center;
}

.settings-panel {
    position: absolute;
    top: 10px;
    right: 10px;
    width: 300px;
    box-shadow: 0 3px 1px -2px rgba(0, 0, 0, .2), 0 2px 2px 0 rgba(0, 0, 0, .14), 0 1px 5px 0 rgba(0, 0, 0, .12);
    background: white;
    padding: 12px 10px;
    z-index: 10000;
    display: flex;
    border-radius: 4px;
    flex-direction: column;
}

.setting-header {
    position: relative;
    padding: 0.2rem 0.2rem 0.6rem;
    font-size: 1.35rem;
    font-weight: bold;
    text-align: left;
    border-bottom: 1px solid #e0e0e0;
}

.close {
    position: absolute;
    top: 0;
    right: 0;
    padding: 0.2rem;
    font-size: 0.8rem;
    font-weight: bold;
    text-align: right;
    cursor: pointer;
    transition: all 0.2s ease-in;
}

.close:hover {
    color: rgb(25, 118, 210);
}

.device-icon-selected {
    position: absolute;
    z-index: -1;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    background-color: rgba(114, 114, 114, .55);
    border-radius: 100%;
    width: 45px;
    height: 45px;
}

.icon-selected {
    position: relative;
}

.icon-popup {
    position: absolute;
    top: 50%;
}
</style>
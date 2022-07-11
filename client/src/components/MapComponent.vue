<template>
    <div class="map-container">
        <div id="map-target"></div>
    </div>
</template>

<script>
import { Holder } from '../helpers'
import "leaflet/dist/leaflet.css";
import L from "leaflet"
import "leaflet-rotatedmarker"

export default {
    name: 'MapComponent',
    // add our rerender triggers to our data updater and initalialize map and resize trigger
    mounted() {
        Holder.onUpdate(this.updateDevices)
        this.initMap()
        window.addEventListener('resize', this.resizer)
    },
    // remove our rerender triggers from our data updater and remove map and resize trigger
    unmounted() {
        Holder.removeUpdate(this.updateDevices)
        window.removeEventListener('resize', this.resizer)
    },
    data() {
        return {
            // all devices
            devices: [],
            search: '',
            // filtered devices to display
            display: [],
            map: null,
            center: [0, 0],
            first: true,
            featureLayer: null,
            loaded: false,
            tileLayer: null,
            // colors based on device status
            colors: {
                off: "rgba(255, 0, 0, .67)",
                driving: "rgba(0, 255, 0, .67)",
                idle: "rgba(135, 206, 250, .67)",
                offline: "rgba(218, 223, 225, 1)",
            },
            settings: {
                showLabels: true,
                grouping: true
            }
        }
    },
    methods: {
        // main update method
        updateDevices(data) {
            if (!data) return
            // create copy of array to always trigger rerender
            this.devices = data instanceof Array ? [...data] : { ...data }
            this.display = this.devices
            // if map is loaded, update map
            if (this.map && this.loaded) {
                if (this.first) {
                    this.first = false
                    // center map on a box that contains all devices
                    if (this.devices.length) this.map.fitBounds(this.getDevicesView())
                    // resize map after loading
                    this.map.invalidateSize()
                }

                if (this.display.length && this.featureLayer) {
                    this.featureLayer.clearLayers()
                    // remap all of our features to new devices
                    this.featureLayer.addData(this.display.map(device => ({
                        type: 'Feature',
                        // "Point" for geojson will be converted to marker
                        geometry: {
                            type: 'Point',
                            coordinates: [device.lng, device.lat]
                        },
                        // props passed to marker so that we can access them later
                        properties: {
                            device_id: device.device_id,
                            display_name: device.display_name,
                            drive_status: device.drive_status,
                            drive_status_begin_time: device.drive_status_begin_time,
                            heading: device.heading,
                            online: device.online,
                            lat: device.lat,
                            lng: device.lng,
                            color: device.online ? this.colors[device.drive_status] : this.colors['offline'],
                        }
                    })))
                }
            }
        },
        updateSettings(settings) {
            this.settings = settings
            this.updateDevices(this.devices)
        },
        initMap() {
            // view controls custom
            const map = L.map('map-target', {
                // attributionControl: false,
                zoomControl: false,
                // renderer: L.canvas(),
            });

            // Open Street Map Tile Layer
            this.tileLayer = L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
                attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
            }).addTo(map);

            map.setMaxZoom(18);
            map.setMinZoom(3);
            // geoJSON function wrapper
            this.featureLayer = L.geoJSON(null, {
                // on point create. Returns marker to be created with geoJSON data
                pointToLayer: (feature, latlng) => {
                    // the icon for our marker
                    const svgIcon = L.divIcon({
                        className: 'device-icon',
                        // the physical HTML element being disaplyed at point [lat, lon]
                        html: `<div class="device-icon-inner">
                            <svg
                              width="24"
                              height="40"
                              viewBox="-26.5 -71.5 53 73"
                              preserveAspectRatio="none"
                              xmlns="http://www.w3.org/2000/svg"
                            >
                              <path 
                                shape-rendering="optimizeQuality" 
                                d="M 25 0 L 0 -10 L -25 0 L 0 -55 Z" 
                                stroke="#000000" 
                                stroke-width="1" 
                                fill="${feature.properties.color}"
                              />
                            </svg>
                        </div>`,
                        iconSize: [40, 40],
                        iconAnchor: [20, 20],
                        popupAnchor: [0, -20],
                    })

                    const marker = L.marker(latlng, {
                        icon: svgIcon,
                        // imported with leaflet-rotatedmarker
                        rotationAngle: feature.properties.heading,
                    })

                    if (settings.showLabels) marker.bindTooltip(feature.properties.display_name).openTooltip()

                    return marker
                },
                // on feature create
                onEachFeature: (feature, layer) => {
                    layer.on('click', () => {
                        // on click tell parent we have clicked on a device
                        this.$emit('device-click', feature.properties)
                    })
                }
            }).addTo(map);
            // laod trigger
            map.on("load", () => {
                setTimeout(() => {
                    map.invalidateSize()
                    this.loaded = true
                    this.$emit("map-loaded", {
                        moveCenter: this.moveCenter,

                    })
                }, 10)
                this.map = map
            })
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

                    L.DomEvent.on(container, 'click', L.DomEvent.stopPropagation);
                    L.DomEvent.on(container, 'click', L.DomEvent.preventDefault);

                    L.DomEvent.on(zoomInButton, 'click', L.DomEvent.stopPropagation);
                    L.DomEvent.on(zoomInButton, 'click', L.DomEvent.preventDefault);
                    L.DomEvent.on(zoomInButton, 'click', () => {
                        map.zoomIn();
                    });

                    L.DomEvent.on(zoomOutButton, 'click', L.DomEvent.stopPropagation);
                    L.DomEvent.on(zoomOutButton, 'click', L.DomEvent.preventDefault);
                    L.DomEvent.on(zoomOutButton, 'click', () => {
                        map.zoomOut();
                    });
                    L.DomEvent.on(zoomOutMapButton, 'click', L.DomEvent.stopPropagation);
                    L.DomEvent.on(zoomOutMapButton, 'click', L.DomEvent.preventDefault);
                    L.DomEvent.on(zoomOutMapButton, 'click', () => map.fitBounds(parent.getDevicesView()));

                    return container;
                }
            })

            L.control.viewControls = function (options) {
                return new L.Control.ViewControls(options);
            }

            L.Control.Settings = L.Control.extend({
                onAdd: function () {
                    const container = L.DomUtil.create('div', 'view-controller');

                    const settingsButton = L.DomUtil.create('button', 'view-controller-button', container);
                    const settings = L.DomUtil.create('span', 'material-icons', settingsButton);
                    settings.innerHTML = 'settings';

                    L.DomEvent.on(container, 'click', L.DomEvent.stopPropagation);
                    L.DomEvent.on(container, 'click', L.DomEvent.preventDefault);

                    L.DomEvent.on(settingsButton, 'click', L.DomEvent.stopPropagation);
                    L.DomEvent.on(settingsButton, 'click', L.DomEvent.preventDefault);
                    L.DomEvent.on(settingsButton, 'click', () => {
                        // map.zoomIn();
                    });

                    return container;
                }
            })

            L.control.settings = function (options) {
                return new L.Control.Settings(options);
            }

            L.control.settings({ position: 'topright' }).addTo(map);
            L.control.viewControls({ position: 'topleft' }).addTo(map);
            // arbitrary default zoom level
            map.setView(this.center, 13)
        },
        moveCenter(lat, lng) {
            this.map.setView([lat, lng], 13)
            this.center = [lat, lng]
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
}
</script>

<style>
.map-container {
    flex: 1;
    height: 100%;
    box-sizing: border-box;
}

#map-target {
    width: 100%;
    height: 100%;
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
}

.view-controller-button:hover {
    background-color: #f5f5f5;
}
</style>
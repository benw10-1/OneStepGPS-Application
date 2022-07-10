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
                idle: "rgba(255, 255, 0, .67)",
            },
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
                        }
                    })))
                }
            }
        },
        initMap() {
            const map = L.map('map-target', {
                // attributionControl: false,
                // zoomControl: false,
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
                                stroke="#808080" 
                                stroke-width="1" 
                                fill="${this.colors[feature.properties.drive_status]}"
                              />
                            </svg>
                        </div>`,
                        iconSize: [40, 40],
                        iconAnchor: [20, 20],
                        popupAnchor: [0, -20],
                    })

                    return L.marker(latlng, {
                        icon: svgIcon,
                        // imported with leaflet-rotatedmarker
                        rotationAngle: feature.properties.heading,
                    })
                },
                // on feature create
                onEachFeature: (feature, layer) => {
                    layer.on('click', () => {
                        // on click tell parent we have clicked on a device
                        this.$emit('device-click', feature.properties)
                    })
                }
            }).addTo(map);

            map.on("load", () => {
                setTimeout(() => {
                    map.invalidateSize()
                    this.loaded = true
                }, 10)
                this.map = map
            })

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

<style scoped>
.map-container {
    flex: 1;
    height: 100%;
    box-sizing: border-box;
}

#map-target {
    width: 100%;
    height: 100%;
}
</style>
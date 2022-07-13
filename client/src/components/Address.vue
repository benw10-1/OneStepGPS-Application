<template>
    <div class="address-container">
        <div v-if="loading" class="loader"></div>
        <div v-else class="address-text">
            {{ address }}
        </div>
    </div>
</template>

<script>
import { Requests } from '@/helpers'

export default {
    name: "AddressComponent",
    props: {
        latlng: {
            type: Object,
            required: true,
        },
    },
    data() {
        return {
            loading: true,
            address: '',
        }
    },
    watch: {
        latlng: {
            immediate: true,
            handler(latlng) {
                this.getAddress(latlng)
            },
        },
    },
    methods: {
        getAddress(latlng=this.latlng) {
            this.loading = true
            Requests.reverseGeocode(latlng.lat, latlng.lng).then(response => {
                this.address = response.address
                this.loading = false
            })
        },
    },
    mounted() {
        this.getAddress()
    },
}
</script>

<style>
.address-container {
    display: flex;
    justify-content: v-bind("loading ? 'center' : 'unset'");
    align-items: v-bind("loading ? 'center' : 'unset'");
    width: 100%;
    height: 100%;
}
.address-text {
    font-size: inherit;
    font-weight: inherit;
    color: inherit;
}
</style>
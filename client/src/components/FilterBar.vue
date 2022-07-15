<template>

</template>

<script>

export default {
    name: "FilterBar",
    props: {
        tags: {
            type: Array,
            required: true,
        },
        onChange: {
            type: Function,
            required: true,
        },
        initial: {
            type: Array,
            required: true,
        },
    },
    data() {
        return {
            checked: this.initial,
            
        }
    }
}
</script>

<style scoped>

</style>
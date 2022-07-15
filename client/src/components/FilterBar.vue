<template>
    <ExpandableVue />
</template>

<script>
import ExpandableVue from './Expandable.vue'
import preferenceHolder from '@/helpers/preferenceHolder'

export default {
    name: "FilterBar",
    components: {
        ExpandableVue
    },
    props: {
        disabled: {
            type: Boolean,
            default: false,
        },
        tags: {
            type: Array,
            default: ["driving", "idle", "off", "offline"],
        }
    },
    data() {
        return {
            checked: this.initial,
            expanded: false,
            driveStates: {
                off: "Stopped",
                driving: "Driving",
                idle: "Idle",
                offline: "No signal",
            },
        }
    },
    methods: {
        onChange(tag) {
            if (this.disabled) {
                return
            }
            if (this.checked.includes(tag)) {
                this.checked = this.checked.filter(t => t !== tag)
            } else {
                this.checked = [...this.checked, tag]
            }
            
        },
        onExpand() {
            this.expanded = true
        },
        onCollapse() {
            this.expanded = false
        },
    },
}
</script>

<style scoped>

</style>
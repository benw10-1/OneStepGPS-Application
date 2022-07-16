<template>
    <ExpandableVue :onChange="onChangeExpand" :expanded="expanded" :snap="disabled">
        <template v-slot:header>
            <div class="filter-header">
                <div class="filter-header-left">
                    <span class="material-icons-outlined">{{ expanded ? 'expand_less' : 'expand_more' }}</span>
                    <span class="header-text">Filter Devices</span>
                </div>
                <IconButton :icon="disabled ? 'filter_list_off' : 'filter_list'" :size="24" :onClick="toggleDisabled" />
            </div>
        </template>
        <template v-slot:default>
            <div class="filter-content">
                <div class="filter-group">
                    <div class="filter-group-header">
                        <span class="filter-group-header-text">Status</span>
                    </div>
                    <div class="filter-group-content">
                        <FilterButton v-for="(tag, i) of tags" :checked="checked.includes(tag)" :onClick="onChange" :tag="tag" :text="driveStates[tag]" :key="i" />
                    </div>
                </div>
            </div>
        </template>
    </ExpandableVue>
</template>

<script>
import ExpandableVue from './Expandable.vue'
import preferenceHolder from '@/helpers/preferenceHolder'
import IconButton from './IconButton.vue'
import FilterButton from './FilterButton.vue'

export default {
    name: "FilterBar",
    components: {
        ExpandableVue,
        IconButton,
        FilterButton
    },
    props: {
        tags: {
            type: Array,
            default: () => ["driving", "idle", "off", "offline"],
        }
    },
    data() {
        return {
            checked: preferenceHolder.get()?.filter?.tags ?? this.tags,
            expanded: false,
            disabled: preferenceHolder.get()?.filter?.disabled,
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
            preferenceHolder.set({
                filter: {
                    tags: this.checked,
                    disabled: this.disabled,
                }
            })
        },
        onChangeExpand(expanded) {
            this.expanded = expanded
        },
        collapse() {
            this.expanded = false
        },
        toggleDisabled() {
            this.disabled = !this.disabled
            preferenceHolder.set({
                filter: {
                    tags: this.checked,
                    disabled: this.disabled,
                }
            })
        },
    },
    watch: {
        disabled(newValue) {
            if (newValue) {
                this.collapse()
            }
        },
    },
}
</script>

<style scoped>
.filter-header {
    display: flex;
    width: 100%;
    align-items: center;
    justify-content: space-between;
    padding: 0.5rem .66rem;
    box-sizing: border-box;
}
.filter-header-left {
    display: flex;
    align-items: center;
    justify-content: flex-start;
}

.header-text {
    margin-left: 0.5rem;
    font-size: 1.2rem;
    font-weight: 500;
}

.filter-content {
    padding: 0.5rem;
    width: 100%;
    height: fit-content;
    box-sizing: border-box;
}
.filter-group-content {
    display: flex;
    width: 100%;
    height: fit-content;
    box-sizing: border-box;
    flex-wrap: wrap;
}
</style>
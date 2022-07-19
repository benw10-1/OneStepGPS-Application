<template>
    <ExpandableVue :onChange="nextSort">
        <template v-slot:header>
            <div class="sort-header">
                <span class="header-text">Sort Devices</span>
                <div class="icon-container">
                    <span class="up-arrow material-icons-outlined">arrow_right_alt</span>
                    <span class="down-arrow material-icons-outlined">arrow_right_alt</span>
                </div>
            </div>
        </template>
    </ExpandableVue>
</template>

<script>
import ExpandableVue from './Expandable.vue'
import { PreferenceHolder } from '@/helpers'

export default {
    name: "SortBar",
    components: {
        ExpandableVue,
    },
    mounted() {
        PreferenceHolder.onUpdate(this.prefUpdate)
    },
    unmounted() {
        PreferenceHolder.removeUpdate(this.prefUpdate)
    },
    data() {
        return {
            sort: "none",
            selectedColor: "rgb(25, 118, 210)",
            unselectedColor: "rgba(0, 0, 0, .33)",
        }
    },
    methods: {
        nextSort() {
            switch (this.sort) {
                case "none":
                    this.sort = "asc";
                    break;
                case "asc":
                    this.sort = "desc";
                    break;
                case "desc":
                    this.sort = "none";
                    break;
                default:
                    this.sort = "none";
                    break;
            }
        },
        prefUpdate(prefs) {
            this.sort = prefs.sort ?? "none"
        },
    },
    watch: {
        sort(newSort) {
            PreferenceHolder.set({ sort: newSort }, this.prefUpdate)
        },
    },
}
</script>

<style scoped>
.header-text {
    margin-left: 0.5rem;
    font-size: 1.2rem;
    font-weight: 500;
}
.sort-header {
    display: flex;
    width: 100%;
    align-items: center;
    justify-content: space-between;
    padding: 0.5rem;
    box-sizing: border-box;
}
.icon-container {
    width: 30px;
    height: 30px;
    position: relative;
}
.up-arrow {
    position: absolute;
    top: 0;
    left: -1px;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 1.5rem;
    color: v-bind("(sort === 'asc') ? selectedColor : unselectedColor");
    transform: rotate(-90deg);
    transition: all .2s ease-in;
}
.down-arrow {
    position: absolute;
    bottom: 0;
    right: -1px;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 1.5rem;
    color: v-bind("(sort === 'desc') ? selectedColor : unselectedColor");
    transform: rotate(90deg);
    transition: all .2s ease-in;
}
</style>
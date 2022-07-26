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
import { settingsStore } from '@/helpers'
import { mapActions, mapState } from 'pinia'

export default {
    name: "SortBar",
    components: {
        ExpandableVue,
    },
    computed: {
        ...mapState(settingsStore, ['sortSettings']),
    },
    data() {
        return {
            selectedColor: "rgb(25, 118, 210)",
            unselectedColor: "rgba(0, 0, 0, .33)",
        }
    },
    methods: {
        ...mapActions(settingsStore, ['updateSort']),
        nextSort() {
            switch (this.sortSettings) {
                case "none":
                    this.updateSort("asc");
                    break;
                case "asc":
                    this.updateSort("desc");
                    break;
                case "desc":
                    this.updateSort("none");
                    break;
                default:
                    this.updateSort("none");
                    break;
            }
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
    color: v-bind("(sortSettings === 'asc') ? selectedColor : unselectedColor");
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
    color: v-bind("(sortSettings === 'desc') ? selectedColor : unselectedColor");
    transform: rotate(90deg);
    transition: all .2s ease-in;
}
</style>
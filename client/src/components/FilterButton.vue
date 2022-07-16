<template>
    <div class="filter-button" @click="onClick_">
        <div class="filter-text">
            <span>{{ text }}</span>
        </div>
        <div class="filter-button-icon-container">
            <span class="material-icons-outlined">check</span>
        </div>
    </div>
</template>

<script>

export default {
    name: "FilterButton",
    props: {
        checked: {
            type: Boolean,
            default: true,
        },
        onClick: {
            type: Function,
            required: true,
        },
        tag: {
            type: String,
            required: true,
        },
        text: {
            type: String,
            required: true,
        },
    },
    methods: {
        onClick_() {
            this.onClick(this.tag)
        },
    },
}
</script>

<style scope>
.filter-button {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 0.5rem;
    border-radius: 0.25rem;
    background-color: rgba(111, 111, 111, .40);
    cursor: pointer;
    box-sizing: border-box;
    margin: 8px 8px 8px 0;
}

.filter-button-icon-container {
    display: grid;
    place-items: center;
    width: v-bind("checked ? '25px' : '0'");
    height: 25px;
    overflow: hidden;
    transition: all 0.2s ease-in;
    user-select: none;
    margin-left: 4px;
}

.filter-text {
    user-select: none;
}
</style>
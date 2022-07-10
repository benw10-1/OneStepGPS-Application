<template>
    <button class="custom-button" v-bind:disabled="disabled" v-on:click="onClick">
        {{ value }}
    </button>
</template>

<script>
export default {
    name: 'CustomButton',
    props: {
        value: {
            type: String,
            default: '',
        },
        type: {
            type: String,
            default: 'filled',
        },
        disabled: {
            type: Boolean,
            default: false,
        },
        onClick: {
            type: Function,
            default: () => {}
        },
        padding: {
            type: String,
            default: '0.5rem 1rem',
        },
        margin: {
            type: String,
            default: '0',
        },
        width: {
            type: String,
            default: 'auto',
        },
        color: {
            type: String,
            default: 'white',
        },
        background: {
            type: String,
            default: 'rgb(25, 118, 210)',
        },
    },
}
</script>

<style scoped>
.custom-button {
    background-color: v-bind(background);
    border-radius: 3px;
    color: v-bind(color);
    cursor: pointer;
    font-size: 14px;
    font-stretch: 100%;
    font-weight: 600;
    margin: v-bind(margin);
    padding: v-bind(padding);
    letter-spacing: 0.02857em;
    user-select: none;
    text-align: center; 
    /* text-transform: uppercase; */
    box-sizing: border-box;
    border: none;
    transition: background-color 250ms cubic-bezier(0.4, 0, 0.2, 1) 0ms, box-shadow 250ms cubic-bezier(0.4, 0, 0.2, 1) 0ms, border-color 250ms cubic-bezier(0.4, 0, 0.2, 1) 0ms, color 250ms cubic-bezier(0.4, 0, 0.2, 1) 0ms;
    width: v-bind(width);
    box-shadow: rgba(0, 0, 0, 0.2) 0px 3px 1px -2px, rgba(0, 0, 0, 0.14) 0px 2px 2px 0px, rgba(0, 0, 0, 0.12) 0px 1px 5px 0px
}
</style>

<template>
    <div class="custom-input">
        <div class="label-container" v-if="label">
            <span>{{ label }}</span>
        </div>
        <div class="input-container">
            <input type="text" v-model="onChange ? value : undefined" @input="(e) => { onChange ? onChange(e.currentTarget.value) : null }" />
        </div>
    </div>
</template>

<script>
export default {
    name: 'CustomInput',
    props: {
        value: {
            type: String,
            default: '',
        },
        placeholder: {
            type: String,
            default: '',
        },
        type: {
            type: String,
            default: 'text',
        },
        disabled: {
            type: Boolean,
            default: false,
        },
        autofocus: {
            type: Boolean,
            default: false,
        },
        autocomplete: {
            type: String,
            default: 'off',
        },
        error: {
            type: String,
            default: '',
        },
        label: {
            type: String,
            default: '',
        },
        onChange: {
            type: Function,
            default: undefined,
        },
    },
    data() {
        return {
            focused: false,
        }
    },
}
</script>

<style>

</style>

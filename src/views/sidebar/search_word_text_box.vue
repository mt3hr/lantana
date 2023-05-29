<template>
    <v-text-field class="search_word_text_field" v-model="search_word" :label="'キーワード'" @keypress.enter="emit_updated_search_word" />
</template>

<script setup lang="ts">
import { Ref, ref, nextTick, watch } from 'vue';

const emits = defineEmits<{
    (e: 'errors', errors: Array<string>): void
    (e: 'updated_search_word', word: string): void
}>()

let search_word: Ref<string> = ref("")

defineExpose({
    set_search_word_by_application,
    get_search_word
})

function get_search_word(): string {
    return search_word.value
}
function set_search_word_by_application(new_search_word: string): void {
    search_word.value = new_search_word
}

function emit_errors(errors: Array<string>) {
    emits("errors", errors)
}
function emit_updated_search_word() {
    emits("updated_search_word", search_word.value)
}
</script>

<style>
.search_word_text_field {
    width: 300px;
    min-width: 300px;
    max-width: 300px;
    height: 56px;
    max-height: 56px;
    min-height: 56px;
}
</style>
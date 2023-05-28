<template>
    <v-dialog class="add_lantana_dialog" v-model="is_show" :width="500">
        <add_lantana_view class="add_lantana_view" :is_dialog="true" :option="option" @errors="emit_errors"
            @added_lantana="emit_added_lantana" @added_kmemo="emit_added_kmemo" @added_tag="emit_added_tag"
            @added_text="emit_added_text" @request_close_dialog="hide" />
    </v-dialog>
</template>
<script lang="ts" setup>
import { Ref, ref, watch, nextTick, defineExpose } from 'vue';
import add_lantana_view from '../lantana/add_lantana_view.vue';
import { ApplicationConfig } from '@/lantana_data/application-config';
import { Text } from '@/lantana_data/text';
import { Tag } from '@/lantana_data/tag';
import { Kmemo } from '@/lantana_data/kmemo';
import { Lantana } from '@/lantana_data/lantana';

interface Props {
    option: ApplicationConfig
}

const props = defineProps<Props>()
const emits = defineEmits<{
    (e: 'errors', errors: Array<string>): void
    (e: 'added_lantana', lantana: Lantana): void
    (e: 'added_kmemo', kmemo: Kmemo): void
    (e: 'added_tag', tag: Tag): void
    (e: 'added_kmemo', text: Text): void
}>()

let is_show: Ref<boolean> = ref(false)

defineExpose({ show })

watch(() => is_show.value, () => {
    is_show.value = is_show.value
})

function show() {
    is_show.value = true
}
function hide() {
    is_show.value = false
}
function emit_errors(errors: Array<string>) {
    emits("errors", errors)
}
function emit_added_lantana(lantana: Lantana) {
    emits("added_lantana", lantana)
}
function emit_added_kmemo(kmemo: Kmemo) {
    emits("added_kmemo", kmemo)
}
function emit_added_tag(tag: Tag) {
    emits("added_tag", tag)
}
function emit_added_text(text: Text) {
    emits("added_kmemo", text)
}
</script>
<style>
.add_lantana_view {
    overflow-y: scroll !important;
}

.v-dialog>.v-overlay__content>.v-card {
    display: block !important;
}
</style>
<template>
    <v-card>
        <div>
            <lantana_summary_view v-for="lantana in lantanas" :lantana="lantana" :key="lantana.lantana_id"
                @errors="emit_errors" @added_tag="emit_added_tag" @added_text="emit_added_text"
                @copied_lantana_id="emit_copied_lantana_id" @deleted_lantana="emit_deleted_lantana"
                @clicked_lantana="emit_clicked_lantana" @added_kmemo="emit_added_kmemo" />
            <v-overlay v-model="is_loading" contained :persistent="true">
                <div class="progress_overlay">
                    <v-progress-circular class="progress" indeterminate :color="'indigo'" />
                </div>
            </v-overlay>
        </div>
    </v-card>
</template>

<script setup lang="ts">
import lantana_summary_view from './lantana_summary_view.vue';
import { Ref, ref, watch, nextTick } from 'vue';
import { Lantana } from '@/lantana_data/lantana';
import { Tag } from '@/lantana_data/tag';
import { Text } from '@/lantana_data/text';
import { Kmemo } from '@/lantana_data/kmemo';

interface Props {
    lantanas: Array<Lantana>
    loading: boolean
}

const props = defineProps<Props>()
const emits = defineEmits<{
    (e: 'errors', errors: Array<string>): void
    (e: 'copied_lantana_id', lantana: Lantana): void
    (e: 'added_kmemo', kmemo: Kmemo): void
    (e: 'added_tag', tag: Tag): void
    (e: 'added_text', text: Text): void
    (e: 'deleted_lantana', lantana: Lantana): void
    (e: 'clicked_lantana', lantana: Lantana): void
}>()

const is_loading: Ref<boolean> = ref(true)

watch(() => props.loading, () => {
    is_loading.value = props.loading
})

function emit_errors(errors: Array<string>) {
    emits("errors", errors)
}
function emit_copied_lantana_id(lantana: Lantana) {
    emits("copied_lantana_id", lantana)
}
function emit_added_tag(tag: Tag) {
    emits("added_tag", tag)
}
function emit_added_text(text: Text) {
    emits("added_text", text)
}
function emit_deleted_lantana(deleted_lantana: Lantana) {
    emits("deleted_lantana", deleted_lantana)
}
function emit_clicked_lantana(lantana: Lantana) {
    emits("clicked_lantana", lantana)
}
function emit_added_kmemo(kmemo: Kmemo) {
    emits("added_kmemo", kmemo)
}
</script>

<style>
.progress_overlay {
    position: absolute;
    width: -webkit-fill-available;
    height: -webkit-fill-available;
}

.progress {
    top: calc(50% - (32px/2));
    left: calc(50% - (32px/2));
}

.v-overlay__content:has(.progress_overlay),
.progress_overlay {
    width: -webkit-fill-available;
    height: -webkit-fill-available;
}
</style>
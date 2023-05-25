<template>
    <v-menu :style="style" v-model="is_show">
        <v-list>
            <v-list-item @click="show_add_kmemo_dialog">
                <v-list-item-title>Kmemo追加</v-list-item-title>
            </v-list-item>
            <v-list-item @click="show_add_tag_dialog">
                <v-list-item-title>タグ追加</v-list-item-title>
            </v-list-item>
            <v-list-item @click="show_add_text_dialog">
                <v-list-item-title>テキスト追加</v-list-item-title>
            </v-list-item>
            <v-list-item @click="copy_lantana_id_to_clipboard">
                <v-list-item-title>IDをコピー</v-list-item-title>
            </v-list-item>
            <v-list-item @click="show_delete_lantana_dialog">
                <v-list-item-title>削除</v-list-item-title>
            </v-list-item>
        </v-list>
    </v-menu>
    <add_kmemo_to_lantana_dialog :lantana="lantana_info" ref="add_kmemo_dialog_ref" @errors="emit_errors"
        @added_kmemo="emit_added_kmemo" />
    <add_tag_to_lantana_dialog :lantana="lantana_info" ref="add_tag_dialog_ref" @errors="emit_errors"
        @added_tag="emit_added_tag" />
    <add_text_to_lantana_dialog :lantana="lantana_info" ref="add_text_dialog_ref" @errors="emit_errors"
        @added_text="emit_added_text" />
    <delete_lantana_dialog :lantana="lantana_info" ref="delete_task_dialog_ref" @errors="emit_errors"
        @deleted_task="emit_deleted_task" />
</template>

<script setup lang="ts">
import { LantanaInfo } from '@/lantana_data/lantana-info';
import { Ref, ref, watch } from 'vue';
import add_tag_to_lantana_dialog from '../dialog/add_tag_dialog.vue';
import add_kmemo_to_lantana_dialog from '../dialog/add_kmemo_to_lantana_dialog.vue';
import add_text_to_lantana_dialog from '../dialog/add_text_dialog.vue';
import delete_lantana_dialog from '../dialog/delete_task_dialog.vue';
import { Lantana } from '@/lantana_data/lantana';
import { Kmemo } from '@/lantana_data/kmemo';

interface Props {
    lantana_info: LantanaInfo
    x: number
    y: number
}

const props = defineProps<Props>()
const emits = defineEmits<{
    (e: 'errors', errors: Array<string>): void
    (e: 'copied_lantana_id', lantana: Lantana): void
    (e: 'added_tag'): void
    (e: 'added_text'): void
    (e: 'added_kmemo', kmemo: Kmemo): void
    (e: 'deleted_lantana', lantana: Lantana): void
    (e: 'deleted_tag'): void
    (e: 'deleted_text'): void
}>()

let style: Ref<string> = ref(generate_style())
let is_show: Ref<boolean> = ref(false)
const add_tag_dialog_ref = ref<InstanceType<typeof add_tag_to_lantana_dialog> | null>(null);
const add_text_dialog_ref = ref<InstanceType<typeof add_text_to_lantana_dialog> | null>(null);
const delete_lantana_dialog_ref = ref<InstanceType<typeof delete_lantana_dialog> | null>(null);
const add_kmemo_dialog_ref = ref<InstanceType<typeof add_kmemo_to_lantana_dialog> | null>(null);

defineExpose({ show })

watch(() => props.x, () => {
    style.value = generate_style()
})
watch(() => props.y, () => {
    style.value = generate_style()
})

function show() {
    is_show.value = true
}
function generate_style(): string {
    return `{ position: absolute; left: ${props.x}px; top: ${props.y}px; }`
}
function show_add_kmemo_dialog() {
    add_kmemo_dialog_ref.value!.show()
}
function show_add_tag_dialog() {
    add_tag_dialog_ref.value!.show()
}
function show_add_text_dialog() {
    add_text_dialog_ref.value!.show()
}
function copy_lantana_id_to_clipboard() {
    navigator.clipboard.writeText(props.lantana_info.lantana.lantana_id)
    emit_copied_task_id()
}
function show_delete_lantana_dialog() {
    delete_lantana_dialog_ref.value!.show()
}
function emit_errors(errors: Array<string>) {
    emits("errors", errors)
}
function emit_copied_task_id() {
    emits("copied_lantana_id", props.lantana_info.lantana)
}
function emit_added_kmemo(kmemo: Kmemo) {
    emits("added_kmemo", kmemo)
}
function emit_added_tag() {
    emits("added_tag")
}
function emit_added_text() {
    emits("added_text")
}
function emit_deleted_task(deleted_lantana_info: Lantana) {
    emits("deleted_lantana", deleted_lantana_info)
}

</script>

<style scoped></style>
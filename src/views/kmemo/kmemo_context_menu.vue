<template>
    <v-menu :style="style" v-model="is_show">
        <v-list>
            <v-list-item @click="show_add_tag_dialog">
                <v-list-item-title>タグ追加</v-list-item-title>
            </v-list-item>
            <v-list-item @click="show_add_text_dialog">
                <v-list-item-title>テキスト追加</v-list-item-title>
            </v-list-item>
            <v-list-item @click="copy_kmemo_id_to_clipboard">
                <v-list-item-title>IDをコピー</v-list-item-title>
            </v-list-item>
            <v-list-item @click="show_delete_kmemo_dialog">
                <v-list-item-title>削除</v-list-item-title>
            </v-list-item>
        </v-list>
    </v-menu>
    <add_tag_to_kmemo_dialog :kmemo="kmemo" ref="add_tag_dialog_ref" @errors="emit_errors" @added_tag="emit_added_tag" />
    <add_text_to_kmemo_dialog :kmemo="kmemo" ref="add_text_dialog_ref" @errors="emit_errors"
        @added_text="emit_added_text" />
    <delete_kmemo_dialog :kmemo="kmemo" ref="delete_kmemo_dialog_ref" @errors="emit_errors"
        @deleted_kmemo="emit_deleted_kmemo" />
</template>

<script setup lang="ts">
import { Ref, ref, watch } from 'vue';
import add_tag_to_kmemo_dialog from '../dialog/add_tag_to_kmemo_dialog.vue';
import add_text_to_kmemo_dialog from '../dialog/add_text_to_kmemo_dialog.vue';
import delete_kmemo_dialog from '../dialog/delete_kmemo_dialog.vue';
import { Kmemo } from '@/lantana_data/kmemo';
import { Tag } from '@/lantana_data/tag';
import { Text } from '@/lantana_data/text';

interface Props {
    kmemo: Kmemo
    x: number
    y: number
}

const props = defineProps<Props>()
const emits = defineEmits<{
    (e: 'errors', errors: Array<string>): void
    (e: 'copied_kmemo_id', kmemo: Kmemo): void
    (e: 'added_tag', tag: Tag): void
    (e: 'added_text', text: Text): void
    (e: 'deleted_kmemo', kmemo: Kmemo): void
}>()

let style: Ref<string> = ref(generate_style())
let is_show: Ref<boolean> = ref(false)
const add_tag_dialog_ref = ref<InstanceType<typeof add_tag_to_kmemo_dialog> | null>(null);
const add_text_dialog_ref = ref<InstanceType<typeof add_text_to_kmemo_dialog> | null>(null);
const delete_kmemo_dialog_ref = ref<InstanceType<typeof delete_kmemo_dialog> | null>(null);

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
function show_add_tag_dialog() {
    add_tag_dialog_ref.value!.show()
}
function show_add_text_dialog() {
    add_text_dialog_ref.value!.show()
}
function copy_kmemo_id_to_clipboard() {
    navigator.clipboard.writeText(props.kmemo.id)
    emit_copied_kmemo_id()
}
function show_delete_kmemo_dialog() {
    delete_kmemo_dialog_ref.value!.show()
}
function emit_errors(errors: Array<string>) {
    emits("errors", errors)
}
function emit_copied_kmemo_id() {
    emits("copied_kmemo_id", props.kmemo)
}
function emit_added_tag(tag: Tag) {
    emits("added_tag", tag)
}
function emit_added_text(text: Text) {
    emits("added_text", text)
}
function emit_deleted_kmemo(deleted_kmemo_info: Kmemo) {
    emits("deleted_kmemo", deleted_kmemo_info)
}

</script>

<style scoped></style>
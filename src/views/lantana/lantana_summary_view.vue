<template>
    <div class="lantana" @contextmenu.prevent.stop="show_contextmenu" @click="emit_clicked_lantana">
        <p class="time">{{ format_time(lantana.time) }}</p>
        <p class="lantana_content">
            <lantana_flowers_view :editable="false" :mood="lantana.mood" />
        </p>
        <lantana_context_menu :lantana="lantana" :x="x_contextmenu" :y="y_contextmenu" @errors="emit_errors"
            @added_tag="emit_added_tag" @added_text="emit_added_text" @copied_lantana_id="emit_copied_lantana_id"
            @deleted_lantana="emit_deleted_lantana" ref="contextmenu" @added_kmemo="emit_added_kmemo" />
    </div>
</template>

<script setup lang="ts">
import lantana_flowers_view from './lantana_flowers_view.vue';
import lantana_context_menu from './lantana_context_menu.vue';
import { Ref, ref } from 'vue';
import { Kmemo } from '@/lantana_data/kmemo';
import { Lantana } from '@/lantana_data/lantana';
import { Tag } from '@/lantana_data/tag';
import { Text } from '@/lantana_data/text';

interface Props {
    lantana: Lantana
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
const contextmenu = ref<InstanceType<typeof lantana_context_menu> | null>(null);

let x_contextmenu: Ref<number> = ref(0)
let y_contextmenu: Ref<number> = ref(0)

function show_contextmenu(e: MouseEvent) {
    x_contextmenu.value = e.x
    y_contextmenu.value = e.y
    contextmenu.value!.show()
}

function emit_errors(errors: Array<string>) {
    emits("errors", errors)
}
function emit_added_tag(tag: Tag) {
    emits("added_tag", tag)
}
function emit_added_text(text: Text) {
    emits("added_text", text)
}
function emit_copied_lantana_id(lantana: Lantana) {
    emits("copied_lantana_id", lantana)
}
function emit_deleted_lantana(lantana: Lantana) {
    emits("deleted_lantana", lantana)
}
function format_time(time: Date): string {
    let year = time.getFullYear()
    let month = time.getMonth() + 1
    let date = time.getDate()
    let hour = time.getHours()
    let minute = time.getMinutes()
    let second = time.getSeconds()
    const day_of_week = ['日', '月', '火', '水', '木', '金', '土'][time.getDay()]
    const month_str = ('0' + month).slice(-2)
    const date_str = ('0' + date).slice(-2)
    const hour_str = ('0' + hour).slice(-2)
    const minute_str = ('0' + minute).slice(-2)
    const second_str = ('0' + second).slice(-2)
    return year + '/' + month_str + '/' + date_str + '(' + day_of_week + ')' + ' ' + hour_str + ':' + minute_str + ':' + second_str
}
function emit_clicked_lantana() {
    emits("clicked_lantana", props.lantana)
}
function emit_added_kmemo(kmemo: Kmemo) {
    emits("added_kmemo", kmemo)
}
</script>

<style scoped>
.lantana_content {
    white-space: pre-line;
}

time {
    color: gray;
    font-size: small;
}
</style>
<template>
    <div class="kmemo" @contextmenu.prevent.stop="show_contextmenu">
        <!-- <p class="time">{{ format_time(kmemo_info.kmemo.time) }}</p> -->
        <tag_view v-for="tag in kmemo_info?.tags" :tag="tag" @errors="emit_errors" @deleted_tag="emit_deleted_tag" />
        <p class="kmemo_content">
            {{ kmemo_info?.kmemo.content }}
        </p>
        <text_view v-for="text in kmemo_info?.texts" :text="text" @errors="emit_errors" @deleted_text="emit_deleted_text" />
        <kmemo_context_menu :kmemo="kmemo_info.kmemo" :x="x_contextmenu" :y="y_contextmenu" @errors="emit_errors"
            @deleted_text="emit_deleted_text" ref="contextmenu" @added_tag="emit_added_tag" @added_text="emit_added_text"
            @deleted_kmemo="emit_deleted_kmemo" @copied_kmemo_id="emit_copied_kmemo_id" />
    </div>
</template>

<script setup lang="ts">
import { KmemoInfo } from '@/lantana_data/kmemo-info';
import tag_view from '../tag/tag_view.vue';
import text_view from '../text/text_view.vue';
import kmemo_context_menu from './kmemo_context_menu.vue';
import { Ref, ref } from 'vue';
import { Kmemo } from '@/lantana_data/kmemo';
import { Tag } from '@/lantana_data/tag';
import { Text } from '@/lantana_data/text';

interface Props {
    kmemo_info: KmemoInfo
}

const props = defineProps<Props>()
const emits = defineEmits<{
    (e: 'errors', errors: Array<string>): void
    (e: 'copied_kmemo_id', kmemo: Kmemo): void
    (e: 'added_tag', tag: Tag): void
    (e: 'added_text', text: Text): void
    (e: 'deleted_kmemo', kmemo: Kmemo): void
    (e: 'deleted_tag', tag: Tag): void
    (e: 'deleted_text', text: Text): void
}>()
const contextmenu = ref<InstanceType<typeof kmemo_context_menu> | null>(null);

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
function emit_deleted_tag(tag: Tag) {
    emits("deleted_tag", tag)
}
function emit_deleted_text(text: Text) {
    emits("deleted_text", text)
}
function emit_deleted_kmemo(kmemo: Kmemo) {
    emits("deleted_kmemo", kmemo)
}
function emit_added_tag(tag: Tag) {
    emits("added_tag", tag)
}
function emit_added_text(text: Text) {
    emits("added_text", text)
}
function emit_copied_kmemo_id(kmemo: Kmemo) {
    emits("copied_kmemo_id", kmemo)
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
</script>

<style scoped>
.kmemo {
    background-color: #eee;
    border: dashed 1px;
    margin: 8px;
    padding: 8px;
}

.kmemo_content {
    white-space: pre-line;
}

time {
    color: gray;
    font-size: small;
}
</style>
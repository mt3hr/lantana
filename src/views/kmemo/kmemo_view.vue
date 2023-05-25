<template>
    <div class="kmemo" @contextmenu.prevent.stop="show_contextmenu">
        <!-- //TODO タグとテキスト -->
        <p class="kmemo_content">{{ kmemo_info?.kmemo.content }}</p>
        <kmemo_context_menu :kmemo_info="kmemo_info" :x="x_contextmenu" :y="y_contextmenu" @errors="emit_errors"
            @deleted_text="emit_deleted_text" ref="contextmenu" />
    </div>
</template>

<script setup lang="ts">
//TODO メソッド全然足りない
import { KmemoInfo } from '@/lantana_data/kmemo-info';
import kmemo_context_menu from './kmemo_context_menu.vue';
import { Ref, ref } from 'vue';

interface Props {
    kmemo_info: KmemoInfo
}

const props = defineProps<Props>()
const emits = defineEmits<{
    (e: 'errors', errors: Array<string>): void
    (e: 'deleted_text'): void
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
function emit_deleted_text() {
    emits("deleted_text")
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
</style>
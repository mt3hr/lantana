<template>
    <v-menu :style="style" v-model="is_show">
        <v-list>
            <v-list-item @click="delete_text">
                <v-list-item-title>削除</v-list-item-title>
            </v-list-item>
        </v-list>
    </v-menu>
</template>

<script setup lang="ts">
import { DeleteTextRequest } from '@/api_request_response/delete-text-request';
import { LantanaServerAPI } from '@/api_request_response/lantana-server-api';
import { Text } from '@/lantana_data/text';
import { Ref, ref, watch } from 'vue';

interface Props {
    text: Text
    x: number
    y: number
}

const props = defineProps<Props>()
const emits = defineEmits<{
    (e: 'errors', errors: Array<string>): void
    (e: 'deleted_text', text: Text): void
}>()

let style: Ref<string> = ref(generate_style())
let is_show: Ref<boolean> = ref(false)
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
function delete_text() {
    const api = new LantanaServerAPI()
    const request = new DeleteTextRequest()
    request.text_id = props.text.id
    api.delete_text(request)
        .then(res => {
            if (res.errors && res.errors.length != 0) {
                emit_errors(res.errors)
                return
            }
            emits("deleted_text", props.text)
        })
}

function emit_errors(errors: Array<string>) {
    emits("errors", errors)
}
</script>

<style scoped></style>
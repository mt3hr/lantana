<template>
    <v-dialog v-model="is_show" :width="500">
        <v-card class="pa-5">
            <v-card-title>
                テキスト追加
            </v-card-title>
            <v-textarea v-model="text_content" :autofocus="true" />
            <v-card-actions>
                <v-row>
                    <v-col cols="auto">
                        <v-btn @click="submit">
                            追加
                        </v-btn>
                    </v-col>
                    <v-spacer />
                    <v-col cols="auto">
                        <v-btn @click="close_dialog">
                            キャンセル
                        </v-btn>
                    </v-col>
                </v-row>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script setup lang="ts">
import { AddTextRequest } from '@/api_request_response/add-text-request';
import { LantanaServerAPI } from '@/api_request_response/lantana-server-api';
import generate_uuid from '@/generate_uuid';
import { Kmemo } from '@/lantana_data/kmemo';
import { Text } from '@/lantana_data/text';
import { Ref, ref, watch } from 'vue';

interface Props {
    kmemo: Kmemo
}

const props = defineProps<Props>()
const emits = defineEmits<{
    (e: 'errors', errors: Array<string>): void
    (e: 'added_kmemo', text: Text): void
}>()

let text_content: Ref<string> = ref("")
let is_show: Ref<boolean> = ref(false)

defineExpose({ show })

watch(() => is_show.value, () => {
    is_show.value = is_show.value
})

function close_dialog() {
    is_show.value = false
}
function submit() {
    if (text_content.value == "") {
        return
    }
    const api = new LantanaServerAPI()
    const request = new AddTextRequest()
    const text = new Text()
    text.id = generate_uuid()
    text.text = text_content.value
    text.target = props.kmemo.id
    text.time = new Date(Date.now())
    request.text = text
    api.add_text(request)
        .then(res => {
            if (res.errors && res.errors.length != 0) {
                emit_errors(res.errors)
                return
            }
            emit_added_text(text)
            clear_fields()
            close_dialog()
        })
}
function show() {
    is_show.value = true
}
function clear_fields() {
    text_content.value = ""
}

function emit_errors(errors: Array<string>) {
    emits("errors", errors)
}
function emit_added_text(text: Text) {
    emits("added_kmemo", text)
}
</script>

<style></style>
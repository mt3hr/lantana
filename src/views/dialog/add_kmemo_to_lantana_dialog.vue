<template>
    <v-dialog v-model="is_show" :width="500">
        <v-card class="pa-5">
            <v-card-title>
                Kmemo
            </v-card-title>
            <v-textarea v-model="kmemo_content" :autofocus="true" />
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
import { AddKmemoRequest } from '@/api_request_response/add-kmemo-request';
import { LantanaServerAPI } from '@/api_request_response/lantana-server-api';
import generate_uuid from '@/generate_uuid';
import { Lantana } from '@/lantana_data/lantana';
import { Kmemo } from '@/lantana_data/kmemo';
import { Ref, ref, watch } from 'vue';

interface Props {
    lantana: Lantana
}

const props = defineProps<Props>()
const emits = defineEmits<{
    (e: 'errors', errors: Array<string>): void
    (e: 'added_kmemo', kmemo: Kmemo): void
}>()

let kmemo_content: Ref<string> = ref("")
let is_show: Ref<boolean> = ref(false)

defineExpose({ show })

watch(() => is_show.value, () => {
    is_show.value = is_show.value
})

function close_dialog() {
    is_show.value = false
}
function submit() {
    if (kmemo_content.value == "") {
        return
    }
    const api = new LantanaServerAPI()
    const request = new AddKmemoRequest()
    const kmemo = new Kmemo()
    kmemo.id = generate_uuid()
    kmemo.content = kmemo_content.value
    kmemo.time = props.lantana.time
    request.kmemo = kmemo
    api.add_kmemo(request)
        .then(res => {
            if (res.errors && res.errors.length != 0) {
                emit_errors(res.errors)
                return
            }
            emit_added_kmemo(kmemo)
            clear_fields()
            close_dialog()
        })
}
function show() {
    is_show.value = true
}
function clear_fields() {
    kmemo_content.value = ""
}

function emit_errors(errors: Array<string>) {
    emits("errors", errors)
}
function emit_added_kmemo(kmemo: Kmemo) {
    emits("added_kmemo", kmemo)
}
</script>

<style></style>
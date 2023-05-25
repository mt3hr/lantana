<template>
    <v-dialog v-model="is_show" :width="500">
        <v-card class="pa-5">
            <v-card-title>
                Kmemo削除
            </v-card-title>
            <v-card-actions>
                <v-row>
                    <v-col cols="auto">
                        <v-btn @click="delete_kmemo" :autofocus="true">
                            削除
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
import { DeleteKmemoRequest } from '@/api_request_response/delete-kmemo-request';
import { LantanaServerAPI } from '@/api_request_response/lantana-server-api';
import { Kmemo } from '@/lantana_data/kmemo';
import { Ref, ref, watch } from 'vue';

interface Props {
    kmemo: Kmemo
}

const props = defineProps<Props>()
const emits = defineEmits<{
    (e: 'errors', errors: Array<string>): void
    (e: 'deleted_kmemo', kmemo: Kmemo): void
}>()

let is_show: Ref<boolean> = ref(false)

defineExpose({ show })

watch(() => is_show.value, () => {
    is_show.value = is_show.value
})

function close_dialog() {
    is_show.value = false
}
function delete_kmemo() {
    const api = new LantanaServerAPI()
    const request = new DeleteKmemoRequest()
    request.kmemo_id = props.kmemo.id
    api.delete_kmemo(request)
        .then(res => {
            if (res.errors && res.errors.length != 0) {
                emit_errors(res.errors)
                return
            }
            emit_deleted_kmemo(props.kmemo)
            close_dialog()
        })
}
function show() {
    is_show.value = true
}

function emit_errors(errors: Array<string>) {
    emits("errors", errors)
}
function emit_deleted_kmemo(deleted_kmemo: Kmemo) {
    emits("deleted_kmemo", deleted_kmemo)
}
</script>

<style></style>
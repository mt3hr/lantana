<template>
    <v-dialog v-model="is_show" :width="500">
        <v-card class="pa-5">
            <v-card-title>
                Lantana削除
            </v-card-title>
            <v-card-actions>
                <v-row>
                    <v-col cols="auto">
                        <v-btn @click="delete_lantana" :autofocus="true">
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
import { DeleteLantanaRequest } from '@/api_request_response/delete-lantana-request';
import { LantanaServerAPI } from '@/api_request_response/lantana-server-api';
import { Lantana } from '@/lantana_data/lantana';
import { Ref, ref, watch } from 'vue';

interface Props {
    lantana: Lantana
}

const props = defineProps<Props>()
const emits = defineEmits<{
    (e: 'errors', errors: Array<string>): void
    (e: 'deleted_lantana', lantana: Lantana): void
}>()

let is_show: Ref<boolean> = ref(false)

defineExpose({ show })

watch(() => is_show.value, () => {
    is_show.value = is_show.value
})

function close_dialog() {
    is_show.value = false
}
function delete_lantana() {
    const api = new LantanaServerAPI()
    const request = new DeleteLantanaRequest()
    request.lantana_id = props.lantana.lantana_id
    api.delete_lantana(request)
        .then(res => {
            if (res.errors && res.errors.length != 0) {
                emit_errors(res.errors)
                return
            }
            emit_deleted_lantana(props.lantana)
            close_dialog()
        })
}
function show() {
    is_show.value = true
}

function emit_errors(errors: Array<string>) {
    emits("errors", errors)
}
function emit_deleted_lantana(deleted_lantana: Lantana) {
    emits("deleted_lantana", deleted_lantana)
}
</script>

<style></style>
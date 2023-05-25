<template>
    <v-dialog v-model="is_show" :width="500">
        <v-card class="pa-5">
            <v-card-title>
                タグ追加
            </v-card-title>
            <v-text-field v-model="tag_name" @keypress.enter="submit" :autofocus="true" />
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
import { AddTagRequest } from '@/api_request_response/add-tag-request';
import { LantanaServerAPI } from '@/api_request_response/lantana-server-api';
import generate_uuid from '@/generate_uuid';
import { Lantana } from '@/lantana_data/lantana';
import { Tag } from '@/lantana_data/tag';
import { Ref, ref, watch } from 'vue';

interface Props {
    lantana: Lantana
}

const props = defineProps<Props>()
const emits = defineEmits<{
    (e: 'errors', errors: Array<string>): void
    (e: 'added_tag'): void
}>()

let tag_name: Ref<string> = ref("")
let is_show: Ref<boolean> = ref(false)

defineExpose({ show })

watch(() => is_show.value, () => {
    is_show.value = is_show.value
})

function close_dialog() {
    is_show.value = false
}
function submit() {
    if (tag_name.value == "") {
        return
    }
    const api = new LantanaServerAPI()
    const request = new AddTagRequest()
    const tag = new Tag()
    tag.id = generate_uuid()
    tag.tag = tag_name.value
    tag.target = props.lantana.lantana_id
    tag.time = new Date(Date.now())
    request.tag = tag
    api.add_tag(request)
        .then(res => {
            if (res.errors && res.errors.length != 0) {
                emit_errors(res.errors)
                return
            }
            emit_added_tag()
            clear_fields()
            close_dialog()
        })
}
function show() {
    is_show.value = true
}
function clear_fields() {
    tag_name.value = ""
}

function emit_errors(errors: Array<string>) {
    emits("errors", errors)
}
function emit_added_tag() {
    emits("added_tag")
}
</script>

<style></style>
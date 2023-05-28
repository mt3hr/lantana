<template>
    <v-card class="pa-5">
        <v-card-title>
            Lantana追加
        </v-card-title>
        <lantana_flowers_view :mood="mood" :editable="true" @updated_mood="update_mood" ref="lantana_flowers_view_ref"/>
        <add_text_to_lantana_view v-for="(text_data, index) in text_datas_dummy" :key="text_data.text_id"
            @errors="emit_errors" @updated_text_data="(text_data) => update_text_data(index, text_data)"
            @delete_text_request="delete_text_data(text_data.text_id)" />
        <v-row>
            <v-spacer />
            <v-col cols="auto">
                <v-btn @click="add_text_data">テキスト追加</v-btn>
            </v-col>
        </v-row>
        <v-card-actions>
            <v-row>
                <v-col cols="auto">
                    <v-btn @click="submit">
                        追加
                    </v-btn>
                </v-col>
                <v-spacer />
                <v-col cols="auto">
                    <v-btn v-if="is_dialog" @click="emit_request_close_dialog">
                        キャンセル
                    </v-btn>
                    <v-btn v-if="!is_dialog" @click="close_page">
                        キャンセル
                    </v-btn>
                </v-col>
            </v-row>
        </v-card-actions>
    </v-card>
</template>
<script lang="ts" setup>
import { Ref, ref, watch, nextTick, defineExpose } from 'vue';
import add_text_to_lantana_view from './add_text/add_text_to_lantana_view.vue';
import lantana_flowers_view from '../lantana/lantana_flowers_view.vue';
import { LantanaServerAPI } from '@/api_request_response/lantana-server-api';
import { ApplicationConfig } from '@/lantana_data/application-config';
import { Text } from '@/lantana_data/text';
import { Tag } from '@/lantana_data/tag';
import { Kmemo } from '@/lantana_data/kmemo';
import { Lantana } from '@/lantana_data/lantana';
import AddToTextLantanaType from './add_text/add_text_to_lantana_type';
import { AddTextToLantanaData } from './add_text/add_text_to_lantana_data';
import { AddLantanaRequest } from '@/api_request_response/add-lantana-request';
import generate_uuid from '@/generate_uuid';
import { AddKmemoRequest } from '@/api_request_response/add-kmemo-request';
import { AddTextRequest } from '@/api_request_response/add-text-request';

interface Props {
    option: ApplicationConfig
    is_dialog: boolean
}

const props = defineProps<Props>()
const emits = defineEmits<{
    (e: 'errors', errors: Array<string>): void
    (e: 'added_lantana', lantana: Lantana): void
    (e: 'added_kmemo', kmemo: Kmemo): void
    (e: 'added_tag', tag: Tag): void
    (e: 'added_text', text: Text): void
    (e: 'request_close_dialog'): void
}>()
const lantana_flowers_view_ref = ref<InstanceType<typeof lantana_flowers_view> | null>(null);
const text_datas: Ref<Array<AddTextToLantanaData>> = ref(new Array<AddTextToLantanaData>())
const text_datas_dummy: Ref<Array<AddTextToLantanaData>> = ref(new Array<AddTextToLantanaData>())
const mood: Ref<number> = ref(0)

function add_text_data() {
    const text_data = new AddTextToLantanaData()
    const text_data_dummy = new AddTextToLantanaData()
    text_data_dummy.text_id = text_data.text_id
    text_datas.value.push(text_data)
    text_datas_dummy.value.push(text_data) // ダミーを入れておく。submit時に実データをいれる
}
function delete_text_data(text_id: string) {
    let match_index = -1;
    for (let i = 0; i < text_datas.value.length; i++) {
        if (text_datas?.value[i].text_id == text_id) {
            match_index = i
            break
        }
    }
    if (match_index != -1) {
        text_datas.value.splice(match_index, 1)
    }
    for (let i = 0; i < text_datas_dummy.value.length; i++) {
        if (text_datas_dummy?.value[i].text_id == text_id) {
            match_index = i
            break
        }
    }
    if (match_index != -1) {
        text_datas_dummy.value.splice(match_index, 1)
    }
}
function update_text_data(index: number, text_data: AddTextToLantanaData) {
    text_datas.value[index] = text_data
}
function emit_errors(errors: Array<string>) {
    emits("errors", errors)
}
async function submit(): Promise<null> {
    const api = new LantanaServerAPI()
    const now = new Date(Date.now())

    const add_lantana_request = new AddLantanaRequest()
    const lantana = new Lantana()
    lantana.lantana_id = generate_uuid()
    lantana.mood = mood.value
    lantana.time = now
    add_lantana_request.lantana = lantana
    await api.add_lantana(add_lantana_request)
        .then((res) => {
            if (res.errors && res.errors.length != 0) {
                emit_errors(res.errors)
            }
        })
    emits("added_lantana", lantana)

    for (let i = 0; i < text_datas.value.length; i++) {
        if (text_datas.value[i].type == AddToTextLantanaType.kmemo) {
            const add_kmemo_request = new AddKmemoRequest()
            const kmemo = new Kmemo()
            kmemo.id = text_datas.value[i].text_id
            kmemo.content = text_datas.value[i].content
            kmemo.time = now
            add_kmemo_request.kmemo = kmemo
            await api.add_kmemo(add_kmemo_request)
                .then((res) => {
                    if (res.errors && res.errors.length != 0) {
                        emit_errors(res.errors)
                    }
                })
            emits("added_kmemo", kmemo)
        } else if (text_datas.value[i].type == AddToTextLantanaType.text) {
            const add_text_request = new AddTextRequest()
            const text = new Text()
            text.id = text_datas.value[i].text_id
            text.text = text_datas.value[i].content
            text.target = lantana.lantana_id
            text.time = now
            add_text_request.text = text
            await api.add_text(add_text_request)
                .then((res) => {
                    if (res.errors && res.errors.length != 0) {
                        emit_errors(res.errors)
                    }
                })
            emits("added_text", text)
        }
    }
    clear_fields()
    finish()
    return new Promise<null>((resolve) => { })
}
function finish() {
    if (props.is_dialog) {
        emit_request_close_dialog()
    } else {
        close_page()
    }
}
function clear_fields() {
    text_datas.value = new Array<AddTextToLantanaData>()
    text_datas_dummy.value = new Array<AddTextToLantanaData>()
    mood.value = 0
    lantana_flowers_view_ref.value?.set_mood(0)
}
function update_mood(mood_: number) {
    mood.value = mood_
}
function emit_request_close_dialog() {
    emits("request_close_dialog")
}
function close_page() {
    window.close()
}
</script>
<style scoped></style>
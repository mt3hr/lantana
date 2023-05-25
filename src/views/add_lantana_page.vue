<template>
    <sidebar_view :option="option" />
</template>
<script lang="ts" setup>
import { Ref, ref, watch, nextTick, defineExpose } from 'vue';
import sidebar_view from './sidebar/sidebar_view.vue';
import { LantanaServerAPI } from '@/api_request_response/lantana-server-api';
import { ApplicationConfig } from '@/lantana_data/application-config';
import { GetApplicationConfigRequest } from '@/api_request_response/get-application-config-request';

const option: Ref<ApplicationConfig> = ref(new ApplicationConfig())
const message: Ref<string> = ref("")
const is_show_message_snackbar: Ref<boolean> = ref(false)
const api = new LantanaServerAPI()
api.get_application_config(new GetApplicationConfigRequest())
    .then(res => {
        if (res.errors && res.errors.length != 0) {
            write_messages(res.errors)
            return
        }
        option.value = res.application_config
    })

function write_message(message_: string) {
    message.value = message_
    is_show_message_snackbar.value = true
}
const sleep = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))
async function write_messages(messages: Array<string>) {
    let is_first = true
    for (let i = 0; i < messages.length; i++) {
        const message_ = messages[i]
        await sleep(is_first ? 0 : 5000)
        write_message(message_)
        is_first = false
    }
}
</script>
<style scoped></style>
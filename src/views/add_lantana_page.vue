<template>
  <div id="control-height"></div>
    <v-app-bar class="app_bar" app color="indigo" flat dark height="50px">
        <v-toolbar-title>lantana</v-toolbar-title>
        <v-btn icon="mdi-format-align-justify" @click="$router.push(LantanaAPIAddress.lantana_log_viewer_page_address)" />
    </v-app-bar>

    <v-main class="main">
        <add_lantana_view class="add_lantana_view" :option="option" :is_dialog="false" @added_lantana="added_lantana"/>
    </v-main>
    <v-snackbar v-model="is_show_message_snackbar">
        <v-container class="ma-0 pa-0">
            <v-row class="ma-0 pa-0">
                <v-col cols="11" class="ma-0 pa-0">
                    <p>{{ message }}</p>
                </v-col>
                <v-col cols="1" class="ma-0 pa-0">
                    <v-btn icon="mdi-close" @click="is_show_message_snackbar = false" width="20px" height="20px"
                        class="ma-0 pa-0" />
                </v-col>
            </v-row>
        </v-container>
    </v-snackbar>
</template>
<script lang="ts" setup>
import { Ref, ref, watch, nextTick, defineExpose } from 'vue';
import add_lantana_view from './lantana/add_lantana_view.vue';
import { ApplicationConfig } from '@/lantana_data/application-config';
import { LantanaServerAPI } from '@/api_request_response/lantana-server-api';
import { GetApplicationConfigRequest } from '@/api_request_response/get-application-config-request';
import { LantanaAPIAddress } from '@/api_request_response/lantana-api-address';
import { useRoute, useRouter } from 'vue-router';
import { Lantana } from '@/lantana_data/lantana';

const route = useRoute()
const router = useRouter()
if (route.query["page"] == "lantana_log_viewer") {
    router.push(LantanaAPIAddress.lantana_log_viewer_page_address)
}

const is_show_message_snackbar: Ref<boolean> = ref(false)
const message: Ref<string> = ref("")
let option: Ref<ApplicationConfig> = ref(new ApplicationConfig())


const actual_height = window.innerHeight
const element_height = document!.querySelector('#control-height') ? document!.querySelector('#control-height')!.clientHeight : actual_height
const bar_height = (actual_height - element_height) + "px"

const api = new LantanaServerAPI()
const sleep = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))
api.get_application_config(new GetApplicationConfigRequest())
    .then((res) => {
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
async function write_messages(messages: Array<string>) {
    let is_first = true
    for (let i = 0; i < messages.length; i++) {
        const message_ = messages[i]
        await sleep(is_first ? 0 : 5000)
        write_message(message_)
        is_first = false
    }
}
function added_lantana(lantana: Lantana) {
    write_message("lantanaを追加しました")
}
</script>

<style>
.main {
    padding-top: 50px !important;
}

#app,
body,
.html,
.v-application {
    height: 100vh;
    min-height: 100vh;
    max-height: 100vh;
    overflow-y: hidden;
}

.app_bar {
    height: 50px;
    max-height: 50px;
    min-height: 50px;
}

#control-height {
    height: 100vh;
    width: 0;
    position: absolute;
}

.add_lantana_view {
    overflow-y: scroll;
    height: calc((100vh - 50px + v-bind(bar_height)));
    max-height: calc(100vh - 50px + v-bind(bar_height));
    min-height: calc(100vh - 50px + v-bind(bar_height));
}
</style>


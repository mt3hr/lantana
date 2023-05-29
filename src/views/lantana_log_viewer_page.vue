<template>
    <div id="control-height"></div>
    <v-navigation-drawer :width="340" class="navigation_drawer" v-model="is_show_drawer" app>
        <sidebar_view :option="option" ref="sidebar_ref" @errors="write_messages"
            @updated_checked_tags="updated_checked_tags" @updated_mood_filter_query="updated_mood_filter_query"
            @updated_search_word="updated_search_word" @updated_tags_by_user="updated_tags_by_user" />
    </v-navigation-drawer>

    <v-app-bar class="app_bar" app color="indigo" flat dark height="50px">
        <v-app-bar-nav-icon @click.stop="is_show_drawer = !is_show_drawer" />
        <v-toolbar-title>lantana</v-toolbar-title>
    </v-app-bar>

    <v-main class="main">
        <table class="lantana_view">
            <tr>
                <td>
                    <lantana_list_view class="lantana_list_view" :loading="loading" :lantanas="lantanas"
                        @errors="write_messages" @added_tag="added_tag" @added_text="added_text"
                        @copied_lantana_id="copied_lantana_id" @deleted_lantana="deleted_lantana"
                        @clicked_lantana="clicked_lantana" @added_kmemo="added_kmemo" />
                </td>
                <td>
                    <lantana_detail_view class="lantana_detail_view" v-if="watching_lantana_info"
                        :lantana_info="watching_lantana_info" @errors="write_messages" @added_tag="added_tag"
                        @added_text="added_text" @copied_kmemo_id="copied_kmemo_id" @copied_lantana_id="copied_lantana_id"
                        @deleted_kmemo="deleted_kmemo" @deleted_lantana="deleted_lantana" @deleted_tag="deleted_tag"
                        @deleted_text="deleted_text" @added_kmemo="added_kmemo" />
                </td>
            </tr>
        </table>
    </v-main>
    <v-avatar :style="floatingActionButtonStyle()" color="indigo" class="position-fixed">
        <v-btn color="white" icon="mdi-plus" variant="text" @click="show_add_lantana_dialog" />
    </v-avatar>

    <add_lantana_dialog :option="option" @errors="write_messages" @added_lantana="added_lantana" @added_tag="added_tag"
        @added_text="added_text" ref="add_lantana_dialog_ref" />

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
import sidebar_view from './sidebar/sidebar_view.vue';
import lantana_list_view from './lantana/lantana_list_view.vue';
import lantana_detail_view from './lantana/lantana_detail_view.vue';
import add_lantana_dialog from './dialog/add_lantana_dialog.vue';
import { LantanaServerAPI } from '@/api_request_response/lantana-server-api';
import { ApplicationConfig } from '@/lantana_data/application-config';
import { GetApplicationConfigRequest } from '@/api_request_response/get-application-config-request';
import { Lantana } from '@/lantana_data/lantana';
import { Tag } from '@/lantana_data/tag';
import { Text } from '@/lantana_data/text';
import { LantanaInfo } from '@/lantana_data/lantana-info';
import { Kmemo } from '@/lantana_data/kmemo';
import { SearchLantanaRequest } from '@/api_request_response/search-lantana-request';
import { LantanaSearchQuery } from '@/lantana_data/lantana-search-query';
import MoodFilterQuery from './sidebar/mood_filter_query';
import { abort } from 'process';

const sidebar_ref = ref<InstanceType<typeof sidebar_view> | null>(null);
const add_lantana_dialog_ref = ref<InstanceType<typeof add_lantana_dialog> | null>(null);
const option: Ref<ApplicationConfig> = ref(new ApplicationConfig())
const is_show_drawer: Ref<boolean | null> = ref(null)
const loading: Ref<boolean> = ref(false)
const message: Ref<string> = ref("")
const is_show_message_snackbar: Ref<boolean> = ref(false)
const lantanas: Ref<Array<Lantana>> = ref(new Array<Lantana>())
const watching_lantana_info: Ref<LantanaInfo | null> = ref(null)
const watching_lantana: Ref<Lantana | null> = ref(null)
const api = new LantanaServerAPI()

const actual_height = window.innerHeight
const element_height = document!.querySelector('#control-height') ? document!.querySelector('#control-height')!.clientHeight : actual_height
const bar_height = (actual_height - element_height) + "px"

let abort_controller = new AbortController()

api.get_application_config(new GetApplicationConfigRequest())
    .then(res => {
        if (res.errors && res.errors.length != 0) {
            write_messages(res.errors)
            return
        }
        option.value = res.application_config
    })
    .then(() => {
        update_lantana_summary_datas()
    })

async function update_lantana_summary_datas(): Promise<null> {
    abort_controller.abort()
    abort_controller = new AbortController()
    loading.value = true
    const query: LantanaSearchQuery = sidebar_ref.value?.construct_lantana_search_query()!
    const request = new SearchLantanaRequest()
    request.query = query
    await api.search_lantana(request, abort_controller)
        .then(res => {
            if (res.errors && res.errors.length != 0) {
                write_messages(res.errors)
                return
            }
            lantanas.value = res.lantanas
            loading.value = false
        })
        .catch((err) => {
            return // DOMException: The user aborted a request.が飛んで邪魔なので握りつぶす
        })
    return new Promise<null>((resolve) => { return })
}

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
function floatingActionButtonStyle() {
    return {
        'bottom': '10px',
        'right': '10px',
        'height': '50px',
        'width': '50px'
    }
}
async function update_lantana_detail_view(): Promise<null> {
    if (watching_lantana.value == null) {
        return new Promise<null>((resolve) => { return null })
    }
    return api.get_lantana_info(watching_lantana.value)
        .then(lantana_info => {
            watching_lantana_info.value = lantana_info
        })
        .then(() => { return null })
}
function added_tag(tag: Tag) {
    sidebar_ref.value?.update_tag_struct_promise()
    update_lantana_detail_view()
    write_message("タグを追加しました")
    update_lantana_summary_datas()
}
function added_text(text: Text) {
    update_lantana_detail_view()
    write_message("テキストを追加しました")
    update_lantana_summary_datas()
}
function copied_kmemo_id(kmemo: Kmemo) {
    write_message("lantanaのIDをコピーしました")
}
function copied_lantana_id(lantana: Lantana) {
    write_message("lantanaのIDをコピーしました")
}
function deleted_lantana(lantana: Lantana) {
    write_message("lantanaを削除しました")
    update_lantana_summary_datas()
    if (watching_lantana.value && lantana.lantana_id == watching_lantana.value.lantana_id) {
        watching_lantana.value = null
        watching_lantana_info.value = null
    }
}
function show_add_lantana_dialog() {
    add_lantana_dialog_ref.value?.show()
}
function added_lantana(lantana: Lantana) {
    write_message("lantanaを追加しました")
    update_lantana_summary_datas()
}
function added_kmemo(kmemo: Kmemo) {
    write_message("kmemoを追加しました")
    update_lantana_detail_view()
}
function deleted_tag(tag: Tag) {
    write_message("タグを削除しました")
    update_lantana_detail_view()
}
function deleted_text(text: Text) {
    write_message("テキストを削除しました")
    update_lantana_detail_view()
}
function deleted_kmemo(kmemo: Kmemo) {
    write_message("Kmemoを削除しました")
    update_lantana_detail_view()
}
function updated_checked_tags(tags: Array<string>) {
    update_lantana_summary_datas()
}
function updated_mood_filter_query(mood_filter_query: MoodFilterQuery) {
    update_lantana_summary_datas()
}
function updated_search_word(word: string) {
    update_lantana_summary_datas()
}
function updated_tags_by_user() {
    update_lantana_summary_datas()
}
function clicked_lantana(lantana: Lantana) {
    watching_lantana.value = lantana
    update_lantana_detail_view()
}
</script>

<style scoped>
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

.lantana_view_wrap,
.lantana_detail_view,
.lantana_list_view {
    overflow-y: scroll;
    height: calc((100vh - 55px + v-bind(bar_height)));
    max-height: calc(100vh - 55px + v-bind(bar_height));
    min-height: calc(100vh - 55px + v-bind(bar_height));
}

.lantana_list_view,
.lantana_detail_view {
    width: calc(320px);
    max-width: calc(320px);
    min-width: calc(320px);
}
</style>
<template>
    <div>
        <search_word_text_box />
        <mood_filter_view />
        <tag_list_view :option="option" />
    </div>
</template>
<script lang="ts" setup>
import { Ref, ref, watch, nextTick, defineExpose } from 'vue';
import search_word_text_box from './search_word_text_box.vue';
import mood_filter_view from './mood_filter_view.vue';
import tag_list_view from './tag_list_view.vue';
import { ApplicationConfig } from '@/lantana_data/application-config';
interface Props {
    option: ApplicationConfig
}
const props = defineProps<Props>()
</script>
<style scoped></style>
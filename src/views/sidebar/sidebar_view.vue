<template>
    <div class="sidebar">
        <search_word_text_box @errors="emit_errors" @updated_search_word="emit_updated_search_word"
            ref="search_word_text_box_ref" />
        <mood_filter_view @errors="emit_errors" @updated_mood_filter_query="emit_updated_mood_filter_query"
            ref="mood_filter_view_ref" />
        <tag_list_view :option="option" @errors="emit_errors" @updated_by_user="emit_updated_tags_by_user"
            @updated_checked_tags="emit_updated_checked_tags" ref="tag_list_view_ref" />
    </div>
</template>
<script lang="ts" setup>
import { Ref, ref, watch, nextTick, defineExpose } from 'vue';
import search_word_text_box from '@/views/sidebar/search_word_text_box.vue';
import mood_filter_view from '@/views/sidebar/mood_filter_view.vue';
import tag_list_view from '@/views/sidebar/tag_list_view.vue';
import { ApplicationConfig } from '@/lantana_data/application-config';
import MoodFilterQuery from '@/views/sidebar/mood_filter_query';
import { LantanaSearchQuery } from '@/lantana_data/lantana-search-query';

interface Props {
    option: ApplicationConfig
}

const props = defineProps<Props>()
const emits = defineEmits<{
    (e: 'errors', errors: Array<string>): void
    (e: 'updated_search_word', word: string): void
    (e: 'updated_mood_filter_query', query: MoodFilterQuery): void
    (e: 'updated_tags_by_user'): void
    (e: 'updated_checked_tags', tags: Array<string>): void
}>()

const search_word_text_box_ref = ref<InstanceType<typeof search_word_text_box> | null>(null);
const mood_filter_view_ref = ref<InstanceType<typeof mood_filter_view> | null>(null);
const tag_list_view_ref = ref<InstanceType<typeof tag_list_view> | null>(null);

defineExpose({
    construct_lantana_search_query,
    update_tag_struct_promise,
})

function emit_errors(errors: Array<string>) {
    emits("errors", errors)
}
function emit_updated_search_word(word: string) {
    emits("updated_search_word", word)
}
function emit_updated_mood_filter_query(query: MoodFilterQuery) {
    emits("updated_mood_filter_query", query)
}
function emit_updated_tags_by_user() {
    emits("updated_tags_by_user")
}
function emit_updated_checked_tags(tags: Array<string>) {
    emits("updated_checked_tags", tags)
}
function construct_lantana_search_query(): LantanaSearchQuery {
    const query = new LantanaSearchQuery()
    query.lantana_search_type = mood_filter_view_ref.value?.get_mood_query().type!
    query.mood = mood_filter_view_ref.value?.get_mood_query().mood!
    query.tags = tag_list_view_ref.value?.get_checked_tags()!
    query.words = search_word_text_box_ref.value?.get_search_word()!
    return query
}
function get_checked_tags(): Array<string> {
    return tag_list_view_ref.value?.get_checked_tags()!
}
function set_checked_tags_by_application(new_checked_tags: Array<string>) {
    tag_list_view_ref.value?.set_checked_tags_by_application(new_checked_tags)
}
async function update_tag_struct_promise() {
    const checked_tags = get_checked_tags()
    await tag_list_view_ref.value?.update_tags_promise()
    set_checked_tags_by_application(checked_tags)
}
</script>
<style scoped>
.sidebar {
    width: calc(300px);
    max-width: calc(300px);
    min-width: calc(300px);
}
</style>
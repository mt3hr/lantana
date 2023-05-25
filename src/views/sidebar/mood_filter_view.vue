<template>
    <h2>Lantana</h2>
    <lantana_flowers_view ref="lantana_flowers_view_ref" @updated_mood="update_mood_filter_query"/>
    <mood_filter_type_select_box ref="mood_filter_type_select_box_ref" @updated_filter_type="update_mood_filter_query"/>
</template>
<script lang="ts" setup>
import { Ref, ref, watch, nextTick, defineExpose } from 'vue';
import lantana_flowers_view from '../lantana/lantana_flowers_view.vue';
import mood_filter_type_select_box from './mood_filter_type_select_box.vue';
import MoodFilterQuery from './mood_filter_query';

const lantana_flowers_view_ref = ref<InstanceType<typeof lantana_flowers_view> | null>(null);
const mood_filter_type_select_box_ref = ref<InstanceType<typeof mood_filter_type_select_box> | null>(null);

const mood_filter_query: Ref<MoodFilterQuery> = ref(new MoodFilterQuery())

const emits = defineEmits<{
    (e: 'errors', errors: Array<String>): void
    (e: "emit_updated_mood_filter_query", query: MoodFilterQuery): void
}>()

defineExpose({
    get_mood_query,
    set_mood_query_by_application,
})

function update_mood_filter_query() {
    const query = new MoodFilterQuery()
    query.mood = lantana_flowers_view_ref.value?.get_mood()!
    query.type = mood_filter_type_select_box_ref.value?.get_filter_type()!
    mood_filter_query.value = query
    emit_updated_mood_filter_query()
}
function set_mood_query_by_application(query: MoodFilterQuery) {
    mood_filter_query.value = query
}
function get_mood_query(): MoodFilterQuery {
    return mood_filter_query.value
}
function emit_updated_mood_filter_query() {
    emits("emit_updated_mood_filter_query", mood_filter_query.value)
}
</script>
<style scoped></style>
<template>
    <v-select class="mood_filter_type_select_box" v-model="filter_type" :items="filter_types" item-title="title"
        item-value="value" @update:model-value="emit_updated_filter_type" />
</template>
<script lang="ts" setup>
import { LantanaSearchType } from '@/lantana_data/lantana-search-type';
import { Ref, ref, watch, nextTick, defineExpose } from 'vue';

class FilterType {
    title: string = ""
    value: LantanaSearchType = LantanaSearchType.all
}
const filter_type: Ref<LantanaSearchType> = ref(LantanaSearchType.all)
const filter_all = new FilterType()
filter_all.title = "全て"
filter_all.value = LantanaSearchType.all

const filter_match = new FilterType()
filter_match.title = "一致"
filter_match.value = LantanaSearchType.match

const filter_greater_than = new FilterType()
filter_greater_than.title = "以上"
filter_greater_than.value = LantanaSearchType.greater_than

const filter_less_than = new FilterType()
filter_less_than.title = "以下"
filter_less_than.value = LantanaSearchType.less_than

const filter_types: Ref<Array<FilterType>> = ref(new Array<FilterType>())
filter_types.value.push(filter_all)
filter_types.value.push(filter_match)
filter_types.value.push(filter_greater_than)
filter_types.value.push(filter_less_than)

const emits = defineEmits<{
    (e: 'updated_filter_type', search_type: LantanaSearchType): void
    (e: 'errors', errors: Array<String>): void
}>()

defineExpose({
    get_filter_type,
    set_filter_type_by_application,
})

function set_filter_type_by_application(new_filter_type: LantanaSearchType) {
    filter_type.value = new_filter_type
}
function get_filter_type(): LantanaSearchType {
    return filter_type.value
}
function emit_updated_filter_type() {
    emits("updated_filter_type", filter_type.value)
}
</script>
<style scoped>
.mood_filter_type_select_box {
    width: 300px;
    min-width: 300px;
    max-width: 300px;
    height: 56px;
    max-height: 56px;
    min-height: 56px;
}
</style>
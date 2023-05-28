<template>
    <v-card class="pa-5">
        <v-card-title>
            <v-row class="pa-0 ma-0">
                <v-col cols="10" class="pa-0 ma-0">
                    <add_text_to_lantana_type_select_box :text_data="text_data" @errors="emit_errors"
                        @updated_text_type="update_text_type" />
                </v-col>
                <v-col cols="2" class="pa-0 ma-0">
                    <v-btn icon="mdi-close" @click="emit_delete_text_request" />
                </v-col>
            </v-row>
        </v-card-title>
        <v-textarea v-model="text_content" @update:model-value="update_text_content" />
    </v-card>
</template>
<script setup lang="ts">
import { Ref, ref, defineExpose, watch } from 'vue';
import { AddTextToLantanaData } from './add_text_to_lantana_data';
import AddToTextLantanaType from './add_text_to_lantana_type';
import add_text_to_lantana_type_select_box from './add_text_to_lantana_type_select_box.vue';

const emits = defineEmits<{
    (e: 'errors', errors: Array<string>): void
    (e: 'delete_text_request'): void
    (e: 'updated_text_data', text_data: AddTextToLantanaData): void
}>()
const text_data: Ref<AddTextToLantanaData> = ref(new AddTextToLantanaData())
const text_content: Ref<string> = ref("")

watch(text_data, () => {
    emit_updated_text_data(text_data.value)
})

defineExpose({
    get_text_data,
})

function emit_errors(errors: Array<string>) {
    emits("errors", errors)
}
function emit_delete_text_request() {
    emits("delete_text_request")
}
function emit_updated_text_data(text_data: AddTextToLantanaData) {
    emits("updated_text_data", text_data)
}
function update_text_content(content: string) {
    text_data.value.content = content
    emit_updated_text_data(text_data.value)
}
function update_text_type(type: AddToTextLantanaType) {
    text_data.value.type = type
    emit_updated_text_data(text_data.value)
}
function get_text_data(): AddTextToLantanaData {
    return text_data.value
}
</script>
<style scoped></style>
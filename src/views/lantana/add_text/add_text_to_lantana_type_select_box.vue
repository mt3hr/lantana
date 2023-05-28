<template>
    <v-select v-model="text_type" :items="text_types" item-title="title" item-value="value"
        @update:model-value="emit_updated_text_type" />
</template>

<script setup lang="ts">
import { Ref, ref } from 'vue';
import AddToTextLantanaType from './add_text_to_lantana_type';

const emits = defineEmits<{
    (e: 'errors', errors: Array<string>): void
    (e: 'updated_text_type', type: AddToTextLantanaType): void
}>()

class TextTypeSelectModel {
    title: string = ""
    value: AddToTextLantanaType = AddToTextLantanaType.kmemo
}
const text_type_select_model = new Array<TextTypeSelectModel>()
const text_type_kmemo = new TextTypeSelectModel()
text_type_kmemo.title = "Kmemo"
text_type_kmemo.value = AddToTextLantanaType.kmemo
const text_type_text = new TextTypeSelectModel()
text_type_text.title = "テキスト"
text_type_text.value = AddToTextLantanaType.text
text_type_select_model.push(text_type_kmemo)
text_type_select_model.push(text_type_text)

const text_types: Ref<Array<TextTypeSelectModel>> = ref(text_type_select_model)
const text_type: Ref<AddToTextLantanaType> = ref(AddToTextLantanaType.kmemo)

function emit_updated_text_type() {
    emits("updated_text_type", text_type.value)
}
</script>

<style scoped></style>
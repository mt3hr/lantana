<template>
    <table>
        <tr class="lantana_icon_tr">
            <td class="lantana_icon_td">
                <lantana_flower :editable="editable" :state="flower_state1" @clicked_left="mood = 1"
                    @clicked_right="mood = 2" />
            </td>
            <td class="lantana_icon_td">
                <lantana_flower :editable="editable" :state="flower_state2" @clicked_left="mood = 3"
                    @clicked_right="mood = 4" />
            </td>
            <td class="lantana_icon_td">
                <lantana_flower :editable="editable" :state="flower_state3" @clicked_left="mood = 5"
                    @clicked_right="mood = 6" />
            </td>
            <td class="lantana_icon_td">
                <lantana_flower :editable="editable" :state="flower_state4" @clicked_left="mood = 7"
                    @clicked_right="mood = 8" />
            </td>
            <td class="lantana_icon_td">
                <lantana_flower :editable="editable" :state="flower_state5" @clicked_left="mood = 9"
                    @clicked_right="mood = 10" />
            </td>
        </tr>
    </table>
</template>
<script lang="ts" setup>
import { Ref, ref, watch, nextTick, defineExpose } from 'vue';
import LantanaFlowerState from '@/views/lantana/lantana_flower_state';
import lantana_flower from '@/views/lantana/lantana_flower.vue';

interface Props {
    mood: number
    editable: boolean
}

const props = defineProps<Props>()

const mood: Ref<number> = ref(props.mood)
const flower_state1: Ref<LantanaFlowerState> = ref(mood.value >= 2 ? LantanaFlowerState.fill : (mood.value >= 1 ? LantanaFlowerState.half : LantanaFlowerState.none))
const flower_state2: Ref<LantanaFlowerState> = ref(mood.value >= 4 ? LantanaFlowerState.fill : (mood.value >= 3 ? LantanaFlowerState.half : LantanaFlowerState.none))
const flower_state3: Ref<LantanaFlowerState> = ref(mood.value >= 6 ? LantanaFlowerState.fill : (mood.value >= 5 ? LantanaFlowerState.half : LantanaFlowerState.none))
const flower_state4: Ref<LantanaFlowerState> = ref(mood.value >= 8 ? LantanaFlowerState.fill : (mood.value >= 7 ? LantanaFlowerState.half : LantanaFlowerState.none))
const flower_state5: Ref<LantanaFlowerState> = ref(mood.value >= 10 ? LantanaFlowerState.fill : (mood.value >= 9 ? LantanaFlowerState.half : LantanaFlowerState.none))

watch(mood, () => {
    flower_state1.value = (mood.value >= 2 ? LantanaFlowerState.fill : (mood.value >= 1 ? LantanaFlowerState.half : LantanaFlowerState.none))
    flower_state2.value = (mood.value >= 4 ? LantanaFlowerState.fill : (mood.value >= 3 ? LantanaFlowerState.half : LantanaFlowerState.none))
    flower_state3.value = (mood.value >= 6 ? LantanaFlowerState.fill : (mood.value >= 5 ? LantanaFlowerState.half : LantanaFlowerState.none))
    flower_state4.value = (mood.value >= 8 ? LantanaFlowerState.fill : (mood.value >= 7 ? LantanaFlowerState.half : LantanaFlowerState.none))
    flower_state5.value = (mood.value >= 10 ? LantanaFlowerState.fill : (mood.value >= 9 ? LantanaFlowerState.half : LantanaFlowerState.none))
    emit_updated_mood()
})

defineExpose({
    get_mood,
    set_mood,
})

watch(() => props.mood, () => {
    mood.value = props.mood
})

const emits = defineEmits<{
    (e: 'updated_mood', mood: number): void
}>()

function get_mood(): number {
    return mood.value
}
function set_mood(mood_value: number) {
    mood.value = mood_value
}
function emit_updated_mood() {
    emits("updated_mood", mood.value)
}
</script>
<style scoped>
.lantana_icon_tr {
    width: calc(50px * 5);
    max-width: calc(50px * 5);
    min-width: calc(50px * 5);
}

.lantana_icon_td {
    width: 50px !important;
    height: 50px !important;
    max-width: 50px !important;
    min-width: 50px !important;
    max-height: 50px !important;
    min-height: 50px !important;
    display: inline-block;
}
</style>
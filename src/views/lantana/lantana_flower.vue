<template>
    <div class="lantana_icon">
        <img v-if="state == LantanaFlowerState.fill" class="lantana_icon_harf_left" :src="lantana_icon"
            @click="emit_clicked_left" />
        <img v-if="state == LantanaFlowerState.fill" class="lantana_icon_harf_right" :src="lantana_icon"
            @click="emit_clicked_right" />
        <img v-if="state == LantanaFlowerState.half" class="lantana_icon_harf_left" :src="lantana_icon"
            @click="emit_clicked_left" />
        <img v-if="state == LantanaFlowerState.half" class="lantana_icon_harf_right gray" :src="lantana_icon"
            @click="emit_clicked_right" />
        <img v-if="state == LantanaFlowerState.none" class="lantana_icon_harf_left gray" :src="lantana_icon"
            @click="emit_clicked_left" />
        <img v-if="state == LantanaFlowerState.none" class="lantana_icon_harf_right gray" :src="lantana_icon"
            @click="emit_clicked_right" />
    </div>
</template>
<script lang="ts" setup>
import { Ref, ref, watch, nextTick, defineExpose } from 'vue';
import LantanaFlowerState from './lantana_flower_state';
import lantana_icon from '/public/lantana_icon.png'

interface Props {
    state: LantanaFlowerState
}
const props = defineProps<Props>()
const emits = defineEmits<{
    (e: 'clicked_left'): void
    (e: 'clicked_right'): void
}>()

function emit_clicked_left() {
    emits('clicked_left')
}
function emit_clicked_right() {
    emits('clicked_right')
}
</script>
<style scoped>
.lantana_icon {
    position: relative;
    width: 50px !important;
    height: 50px !important;
    max-width: 50px !important;
    min-width: 50px !important;
    max-height: 50px !important;
    min-height: 50px !important;
    display: inline-block;
}

.lantana_icon_fill {
    width: 50px !important;
    height: 50px !important;
    max-width: 50px !important;
    min-width: 50px !important;
    max-height: 50px !important;
    min-height: 50px !important;
    display: inline-block;
    z-index: 10;
}

.lantana_icon_harf_left {
    position: absolute;
    left: 0px;
    width: 25px !important;
    height: 50px !important;
    max-width: 25px !important;
    min-width: 25px !important;
    max-height: 50px !important;
    min-height: 50px !important;
    object-fit: cover;
    object-position: 0 0;
    display: inline-block;
    z-index: 10;
}

.lantana_icon_harf_right {
    position: absolute;
    left: 0px;
    width: 50px !important;
    height: 50px !important;
    max-width: 50px !important;
    min-width: 50px !important;
    max-height: 50px !important;
    min-height: 50px !important;
    display: inline-block;
    z-index: 9;
}

.lantana_icon_none {
    width: 50px !important;
    height: 50px !important;
    max-width: 50px !important;
    min-width: 50px !important;
    max-height: 50px !important;
    min-height: 50px !important;
    display: inline-block;
    z-index: 10;
}

.gray {
    filter: grayscale(100%);
}
</style>
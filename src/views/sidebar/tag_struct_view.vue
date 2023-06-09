<template>
    <tr v-if="is_item()">
        <td>
            <table>
                <tr>
                    <td>
                        <input type="checkbox" v-model="check" @change="update_check_item_by_user"
                            :indeterminate.prop="struct.indeterminate" />
                    </td>
                    <td class="tree_item ml-1" @click="click_item_by_user">{{ struct.key }}</td>
                </tr>
            </table>
        </td>
    </tr>
    <tr v-else>
        <td>
            <table>
                <tr>
                    <td>
                        <input type="checkbox" v-model="check" @change="change_group_by_user"
                            :indeterminate.prop="indeterminate_group" />
                    </td>
                    <td>
                        <span v-if="open_group" style="cursor: default" @click="open_group = !open_group">▽</span>
                        <span v-else style="cursor: default" @click="open_group = !open_group">▷</span>
                    </td>
                    <td @click="click_group_by_user">
                        <div class="tree_item">{{ group_name }}</div>
                    </td>
                </tr>
            </table>
            <table class="ml-4">
                <structures v-show="open_group" v-for="(child_struct, index) in struct_list" :open="get_group_open(index)"
                    :key="index" :struct="child_struct" :group_name="get_group_name(index)"
                    @updated_check_items_by_user="emit_updated_check_items_by_user"
                    @click_items_by_user="emit_click_items_by_user" />
            </table>
        </td>
    </tr>
</template>

<script setup lang="ts">
import structures from './tag_struct_view.vue'
import { Ref, ref, watch } from 'vue';

interface Props {
    struct: any
    group_name: string
    open: boolean
}

const props = defineProps<Props>()
const emits = defineEmits<{
    (e: 'errors', errors: Array<string>): void
    (e: 'click_items_by_user', items: Array<string>): void
    (e: 'updated_check_items_by_user', items: Array<string>, check: boolean, indeterminate: boolean): void
}>()
defineExpose({ get_selected_items })

let open_group: Ref<boolean> = ref(props.open)
let check: Ref<boolean> = ref(false)
let struct_list: Ref<any> = ref(new Array<string>())
let indeterminate_group: Ref<boolean> = ref(false)

watch(() => props.open, () => {
    open_group.value = props.open
})
watch(() => props.struct, () => {
    updated_struct()
    update_check()
})

open_group.value = props.open
updated_struct()
update_check()

// チェックボックスのチェック状態を最新の状態に更新します。
// 親であれば子の状態を見ます
function update_check() {
    if (is_item()) {
        check.value = props.struct.check
    } else {
        let exist_checked = false
        let all_checked = true

        let f = (struct: any) => { }
        let func = (struct: any) => {
            if (struct.key !== undefined && struct.check !== undefined) {
                if (struct.check) {
                    exist_checked = true
                } else {
                    all_checked = false
                }
            } else {
                Object.keys(struct).forEach(name => {
                    f(struct[name])
                })
            }
        }
        f = func
        f(props.struct)

        if (all_checked) {
            indeterminate_group.value = false
            check.value = true
        } else if (exist_checked && !all_checked) {
            check.value = false
            indeterminate_group.value = true
        } else {
            indeterminate_group.value = false
            check.value = false
        }
    }
}
// アイテムではなくの場合に使われます。
// 子アイテムを子アイテム配列に変換してthis.struct_listに収めます。
// this.struct_listはv-forで回して子アイテムとして再帰的に読み込まれます。
function updated_struct() {
    struct_list.value = Object.values(props.struct)
}
// this.structがアイテムであればtrueを、そうではなくグループである場合はfalseを返します。
function is_item() {
    return props.struct.key !== undefined && props.struct.check !== undefined
}
function get_group_open(index: number) {
    let group_name = Object.keys(props.struct)[index]
    if (group_name.endsWith(',close') || group_name.endsWith(', close')) {
        return false
    } else if (group_name.endsWith(',open') || group_name.endsWith(', open')) {
        return true
    }
    return true
}
// 子アイテムのグループ名を取得するためにv-for内から使われます。
function get_group_name(index: number) {
    let group_name = Object.keys(props.struct)[index]
    if (group_name.endsWith(',close') || group_name.endsWith(', close')) {
        group_name = group_name.split(',').slice(0, -1).join(',')
    } else if (group_name.endsWith(',open') || group_name.endsWith(', open')) {
        group_name = group_name.split(',').slice(0, -1).join(',')
    }
    return group_name
}
// アイテムのチェック状態に変更があった場合に呼び出されます。
// すべての子アイテムのcheckの状態を、グループのチェック状態と同じにします。
function change_group_by_user() {
    let items = new Array()
    let f = (struct: any) => { }
    let func = (struct: any) => {
        if (struct.key !== undefined && struct.check !== undefined) {
            items.push(struct.key)
        } else {
            Object.keys(struct).forEach(name => {
                f(struct[name])
            })
        }
    }
    f = func
    f(props.struct)
    emit_updated_check_items_by_user(items, check.value, false)
}
// 子グループ内のアイテムに変更があったときに、それを上に伝えるために呼び出されます。
function emit_updated_check_items_by_user(items: Array<string>, check: boolean, indeterminate: boolean) {
    emits('updated_check_items_by_user', items, check, indeterminate)
}
// 子グループ内の複数のアイテムのみをチェックするように変更があったときに、それを上に伝えるために呼び出されます。
function emit_click_items_by_user(items: Array<string>) {
    emits('click_items_by_user', items)
}
// 子グループ内の一つのアイテムのみをチェックするよう変更があったときに、それを上に伝えるために呼び出されます。
function emit_click_item_by_user(item: string) {
    emit_click_items_by_user([item])
}
// アイテムのチェック状態に変更があったときに呼び出されます。
function update_check_item_by_user() {
    emit_updated_check_items_by_user([props.struct.key], check.value, props.struct.indeterminate)
}
// このアイテムがクリックされたときに呼び出されます。
// このアイテムのみにチェックが入るように上にemitします。
function click_item_by_user() {
    emit_click_item_by_user(props.struct.key)
}
// このアイテムがクリックされたときに呼び出されます。
// このアイテム内のアイテムのみにチェックが入るように上にemitします。
function click_group_by_user() {
    let items = new Array<string>()
    let f = (struct: any) => { }
    let func = (struct: any) => {
        if (struct.key !== undefined && struct.check !== undefined) {
            items.push(struct.key)
        } else {
            Object.keys(struct).forEach(name => {
                f(struct[name])
            })
        }
    }
    f = func
    f(props.struct)
    emit_click_items_by_user(items)
}
// 現在チェックの入っているアイテム名を配列で取得します。
function get_selected_items(): Array<string> {
    let items = new Array<string>()
    let f = (struct: any) => { }
    let func = (struct: any) => {
        if (struct.key !== undefined && struct.check !== undefined) {
            if (struct.check) {
                items.push(struct.key)
            }
        } else {
            Object.keys(struct).forEach(name => {
                f(struct[name])
            })
        }
    }
    f = func
    f(props.struct)
    return items
}
</script>

<style>
.tree_item {
    min-width: 200px;
    cursor: default;
}
</style>
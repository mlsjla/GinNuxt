<template>
    <div class="flex w-full flex-col justify-start items-center">
        <ClientOnly>
            <div class="flex flex-1 flex-col w-full min-h-screen">
                <admin-header></admin-header>
                <div class="flex flex-col flex-grow justify-between items-start relative w-full min-h-full overflow-hidden">
                    <admin-menu class="absolute top-0 left-0 h-full w-64"></admin-menu>
                    <div class="flex w-full flex-col justify-start items-start pl-64">
                        <slot />
                    </div>
                </div>
            </div>
        </ClientOnly>
    </div>
</template>


<script setup>
import { useDictionaryStore } from "@/stores/dictionary";
import { useRoleStore } from "@/stores/role";
import { useMenuStore } from "@/stores/menu";

const dictionaryStore = useDictionaryStore()
provide("dictionaryStore", dictionaryStore)

const roleStore = useRoleStore()
provide("roleStore", roleStore)

const menuStore = useMenuStore()
provide("menuStore", menuStore)

menuStore.init()
roleStore.init()
dictionaryStore.init()
</script>
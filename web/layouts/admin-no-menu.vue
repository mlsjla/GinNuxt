<template>
    <div class="flex flex-col justify-start items-center">
        <ClientOnly>
            <div class="flex w-full flex-col min-h-screen">
                <div class="flex-col h-full header justify-start items-center">
                    <admin-header></admin-header>
                </div>
                <div class="main flex flex-row h-full justify-between items-start overflow-hidden">
                    <slot />
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
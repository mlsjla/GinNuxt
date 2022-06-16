<template>
  <div class="flex flex-col pl-5 justify-start items-start">
    <ul class="menu px-5 bg-base-100 menu-horizontal">
      <li class="bordered">
        <a> 首页 </a>
      </li>
      <li>
        <a> 全局 </a>
      </li>
      <li>
        <a> 用户 </a>
      </li>
      
    </ul>
  </div>
</template>

<script setup>
// This will also work in `<script setup>`
definePageMeta({
  layout: "admin",
});
</script>

<style lang="scss" scoped>
</style>

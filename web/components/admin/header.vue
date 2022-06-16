<template>
  <div class="header p-5 flex flex-row justify-between items-center">
    <div class="logo flex flex-row justify-between items-center">
      <i class="bi-chat-dots-fill" style="font-size: 38px"></i>
      <h2>测试</h2>
      <!-- 爱上钓鱼 -->
      <!-- <img src="https://www.zhanmishu.com/template/zhanmishu_edu/images/logo.png" alt=""> -->
    </div>
    <div class="menu-list flex-grow">
      <ul class="menu p-3 menu-vertical lg:menu-horizontal">
        <li @click="tapMenu(item)" v-for="(item, index) in menuList" :key="index" class="">
          <a>
            {{ item.name }}
          </a>
        </li>
      </ul>
    </div>
    <div class="user-options">
      <span>您好，{{ user.username }}</span> <a class="cursor-pointer" @click="loginOut"><span>&nbsp;[退出]</span></a>
      &nbsp;&nbsp;<nuxt-link target="_blank" to="/">网站首页</nuxt-link>
      <a class="cursor-pointer" @click="removeCache"
        ><span style="margin-left: 20px">清空缓存</span></a
      >
    </div>
  </div>
</template>

<script setup>
import { useMenuStore } from "@/stores/menu";
import { useRoleStore } from "@/stores/role";
import { useUserStore } from "@/stores/user";
import { useRouter, useRoute } from "#imports";
import { ElMessage } from "element-plus";

const router = useRouter();
const menuStore = useMenuStore();
const useStore = useRoleStore()
const { user, token } = useUserStore();

onMounted(() => {
  menuStore.init()
  useStore.init()
})


const menuList = computed(() => {
  return menuStore.menu.filter(item => {
    return item.is_show == 1 && item.status == 1
  })
})


const removeCache = () => {
    ElMessage({
        message: "清除成功",
        type: "success",
    });
}

const tapMenu = (item) => {
  console.log('tapMenu start...')
    if(item.router) {
        router.push(item.router)
    }else if(item?.children && item?.children[0] && item?.children[0]?.router) {
        router.push(item?.children[0]?.router)
    }
    console.log('tapMenu end...')
}
const loginOut = () => {
    router.push('/admin/user/login')
}



</script>

<style lang="scss" scoped>
.header {
  height: 80px;
  background-color: #0054a5;
  color: #fff;

  .logo {
    // background-color: rgba(255, 255, 255, 0.5);
    font-size: 26px;
    color: #fff;
    user-select: none;

    h2 {
      padding-left: 10px;
    }
  }
}

.menu-list {
  padding-left: 34px;
}
</style>

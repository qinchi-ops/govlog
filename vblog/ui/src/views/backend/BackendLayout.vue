<template>
    <a-layout>
        <a-layout-header class="header">
            <div class="header-left">
                我的博客
            </div>
            <div class="header-right">
                <a-space>
                    <a-button type="text">前台</a-button>
                    <a-button :loading="logoutLoadding" type="text" @click="handleLogout"><span
                            style="margin-right: 12px;">退出</span><icon-export /></a-button>
                </a-space>
            </div>
        </a-layout-header>
        <a-layout>
            <a-layout-sider collapsible :width="260" class="sider-bar">侧边栏导航组件</a-layout-sider>
            <a-layout-content><router-view /></a-layout-content>
        </a-layout>
    </a-layout>
</template>



<script setup>

import { LOGOUT } from '@/api/vblog';
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import app from '@/stores/app';

const router =useRouter() 
const logoutLoading =ref(false)
const handleLogout = async () => {
    // 1. 调用logout接口

    try {
        logoutLoading.value =true
        await LOGOUT(app.value.token.refresh_token)
        app.value.token= undefined
        
    } finally {
        logoutLoading.value =false

    }
    // 2. 跳转到登录界面
    router.push({name:'LoginPage'})

}
</script>

<style lang="css" scoped>
.header {
    height: 59px;
    border-bottom: 1px solid var(--color-border);
    padding: 0px 12px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    color: var(--color-netural-8);
}
.header-left{
    font-size: 16px;
    font-weight: 600;
    display: flex;
    justify-content: left;
}
.header-right{
    display: flex;
    justify-content: right;
}

.sider-bar {
    height: calc(100vh-60px);
}
</style>
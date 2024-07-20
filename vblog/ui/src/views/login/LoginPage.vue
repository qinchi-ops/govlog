<template>
    <div class="login-container-wrapper">
        <!-- 登录表单容器 -->
        <div class="login-container">
            <!-- title -->
            <div class="title">欢迎登录博客管理系统</div>
            <!-- 登录表单 -->
            <div>
                <a-form :model="form" @submit="handleSubmit" auto-label-width>
                    <a-form-item field="username" label="用户">
                        <a-input placeholder="please enter your username..." >
                            <template #prefix>
                            <icon-user />
                        </template>
                        </a-input>
                    </a-form-item>
                    <a-form-item field="password" label="密码">
                        <a-input-password
                        v-model="form.password"
                        placeholder="please enter your password..."
                        allow-clear
                        >
                            <template #prefix>
                                <icon-lock />
                            </template>
                        </a-input-password>
                    </a-form-item>
                    <a-form-item>
                        <a-button html-type="submit"  style="width: 20%;">登录</a-button>
                    </a-form-item>
                </a-form>
            </div>
        </div>
    </div>

</template>


<script setup>
import { reactive } from 'vue';

const form = reactive({
    username: '',
    password: '',
    is_member: false,

})

const handleSubmit = (data) =>{
    console.log(data);
}
</script>

<style lang="css" scoped>

.title {
    font-size: 24px;  
    font-weight: 600;
    display: flex;
    justify-content: center;
    margin: 12px;
    color: var(--color-neutral-8);
}
.login-container-wrapper {
    height: 100%;
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
}

.login-container {
    height: 260px;
    width: 460px;
    border: solid 1px var(--color-neutral-3); 
    padding: 12px;
    box-shadow: 0cqh0 -2px 5px rgba(0, 0, 0, 0.1);

}
.login-container :deep(.arco-form-item-label-col){
    padding-right: 16px;
}
</style>
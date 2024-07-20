<template>
    <div>
        <!-- 登录表单容器 -->
        <div class="login-container">
            <!-- title -->
            <div>欢迎登录xxxx系统</div>
            <!-- 登录表单 -->
            <div>
                <a-form :model="form" @submit="handleSubmit">
                    <a-form-item field="username" label="用户">
                        <a-input v-model="form.username" placeholder="please enter your username..." />
                    </a-form-item>
                    <a-form-item field="password" label="密码">
                        <a-input v-model="form.password" placeholder="please enter your password..." />
                    </a-form-item>
                    <a-form-item>
                        <a-button html-type="submit">登录</a-button>
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

<style lang="css" scoped></style>
<template>
    <div class="login-container-wrapper">
        <!-- 登录表单容器 -->
        <div class="login-container">
            <!-- title -->
            <div class="title">欢迎登录博客管理系统</div>
            <!-- 登录表单 -->
            <div>
                <a-form :model="form" @submit="handleSubmit" auto-label-width>
                    <a-form-item  required hide-asterisk field="username" label="用户"
                    :rules="{required:true,message:'please input your username'}">
                        <a-input v-model="form.username" placeholder="please enter your username..." allow-clear>
                            <template #prefix>
                                <icon-user />
                            </template>
                        </a-input>
                    </a-form-item>
                    <a-form-item required hide-asterisk field="password" label="密码" :rules="{required:true,message:'please input your password'}">
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
                    <a-form-item field="checkbox" label="" auto-label-width >
                            <a-checkbox value="1">Is_member</a-checkbox>
                    </a-form-item>

                    <a-form-item>
                        <a-button html-type="submit"  :loading="LoginLoading"  style="width: 20%;">登录</a-button>
                    </a-form-item>
                </a-form>
            </div>
        </div>
    </div>

</template>


<script setup>
import { reactive,ref } from 'vue';
import { LOGIN } from '@/api/vblog';
import { useRouter } from 'vue-router';

const router=useRouter()
const LoginLoading = ref(false)
const form = reactive({
    username: '',
    password: '',
    is_member: false,

})

const handleSubmit = async (data) =>{
    if (!data.erros){
        console.log(data.values)
        try {
            LoginLoading.value = true
            const resp = await LOGIN(data.values)
            // 如果登陆成功了，需要跳转到后台页面
            router.push({name: 'BackendLayout'})
            console.log(resp)
        }
        catch (error){
            console.log(error)

        }
        finally {
            LoginLoading.value = false
        }
        
    }
    
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
    /* fallback for old browsers */
    background: -webkit-linear-gradient(to right, #D4D3DD, #EFEFBB);
    /* Chrome 10-25, Safari 5.1-6 */
    background: linear-gradient(to right, #D4D3DD, #EFEFBB);
    /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */
}

.login-container {
    height: 260px;
    width: 560px;
    border: solid 1px var(--color-neutral-3); 
    padding: 12px;
    box-shadow: 0cqh0 -2px 5px rgba(0, 0, 0, 0.1);
    background-color: #fff;

}
.login-container :deep(.arco-form-item-label-col){
    padding-right: 16px;
}

.login-container :deep(.arco-form-item){
    margin-bottom: 12px;
}
</style>
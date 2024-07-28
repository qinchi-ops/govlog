<template>
    <div>
        <a-page-header :title="isEdit ? '编辑文章':'创建文章'" @back="$router.go(-1)">
            <template #extra>
                <a-button :loading="createLoadding" type="outline" @click="handleSave">
                    <template #icon>
                        <icon-save />
                    </template>
                    保存
                </a-button>
            </template>
        </a-page-header>

        <!-- 提交form  -->
        <a-form ref="formRef" :model="form" layout="vertical" breakpoint="xl">
            <a-form-item field="title" label="文章标题" :rules="{required:true,message:'请输入文章标题'}" >
                <a-input v-model="form.title" placeholder="please enter title..." />
            </a-form-item>
            <a-form-item field="summary" label="文章摘要" :rules="{required:true,message:'请输入文章摘要'}">
                <a-input v-model="form.summary" placeholder="please enter summary..." />
            </a-form-item>
            <a-form-item field="content" label="文章内容" :rules="{required:true,message:'请输入内容'}">
                <!-- https://www.npmjs.com/package/md-editor-v3 -->
                <MdEditor v-model="form.content"  class="md-editor"> </MdEditor>
            </a-form-item>
        </a-form>    
    
    </div>
</template>

<script setup>
import { CREATE_BLOG,GET_BLOG,UPDATE_BLOG } from '@/api/vblog';
import { ref,watch } from 'vue';
import { MdEditor } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import { Notification } from '@arco-design/web-vue';
import { useRouter } from 'vue-router';

const router = useRouter()
const isEdit = ref(false)

//判断模式
const getBlogloadding = ref(false)
watch(
    ()=> router.currentRoute.value.query,
    async (v)=> {
        if (v.id) {
            isEdit.value =true
            //通过id获取博客内容
            try {
                getBlogloadding.value=true
                const resp = await GET_BLOG(v.id)
                form.value =resp
            }finally{
                getBlogloadding.value=false
            }
        }
    },
    {immediate:true}
)
const form = ref({
    title: '',
    author: '',
    content: '',
    summary: '',
    create_by: '',
    tags: {},
})


// form 表单示例
const formRef = ref(null);
const createLoadding = ref(false)
const handleSave = async () => {
    const resp = await formRef.value.validate()
    if (!resp) {
        try {
            createLoadding.value = true
            if (isEdit.value) {
                await UPDATE_BLOG(router.currentRoute.value.query.id,form.value)
                Notification.success('更新成功') 
                router.push({name:'BackendBlogList'})
            }else{
                await CREATE_BLOG(form.value)
                Notification.success('保存成功')
                router.push({name:'BackendBlogList'})
            }
        } finally {
            createLoadding.value = false
        }
    }
}

</script>

<style lang="css" scoped>
.md-editor {
    height: 660px;
}

</style>
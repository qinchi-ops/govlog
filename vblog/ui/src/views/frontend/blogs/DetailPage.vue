<template>
    <div>
        <a-page-header  title="文章详情" @back="$router.go(-1)"> 
        </a-page-header>
        <a-form-item field="title"  class="table-line" > {{form.title}}
            </a-form-item>
            <a-form-item field="author" class="author-line">作者：{{form.author}}
            </a-form-item>           
        <a-form ref="formRef" layout="vertical" breakpoint="xl">
            <a-form-item field="content"  >
                <!-- https://www.npmjs.com/package/md-editor-v3 -->
                <!-- <MdEditor v-model="form.content"  class="md-editor"> </MdEditor> -->
                <MdPreview :editorId="id" :modelValue="form.content" />
                <!-- <MdCatalog :editorId="id" :scrollElement="scrollElement" /> -->
            </a-form-item>
        </a-form>            
    </div>
</template>

<script setup>

import { ref,watch } from 'vue';
import { MdPreview, MdCatalog } from 'md-editor-v3';
import 'md-editor-v3/lib/preview.css';
import { GET_BLOG } from '@/api/vblog';
import { useRouter } from 'vue-router';
const router = useRouter()
const isEdit = ref(false)
// const blogid = ({})
const formRef = ref(null);

const form = ref({
    title: '',
    author: '',
    content: '',
    summary: '',
    create_by: '',
    tags: {},
})
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

const id = 'preview-only';
// const text = ref(GetBlog(29));
// const text = ref('# Hello Editor');
// console.log((GetBlog(29)))
// const scrollElement = document.documentElement;



</script>

<style lang="css" scoped>
.table-line :deep(.arco-form-item-content-flex ){
    font-size: 28px;
    font-weight: 700;
    display: flex;
    justify-content: center;
    margin-bottom: 6px;
}

.author-line :deep(.arco-form-item-content-flex){
    display: flex;
    justify-content: center;
    margin-bottom: 6px;
}
</style>
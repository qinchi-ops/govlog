<template>
    <div>
        文章列表
        <!-- 页头 -->
        <a-breadcrumb style="margin-bottom: 12px;">
            <a-breadcrumb-item>文章管理</a-breadcrumb-item>
            <a-breadcrumb-item>文章列表</a-breadcrumb-item>
        </a-breadcrumb>
        <!-- 数据创建和筛选 -->
        <div class="table-line">
            <div class="table-op">
                <a-button type="primary">
                    
                    <template #icon>
                        <icon-plus />
                    </template>
                    创建文章
                </a-button>
            </div>
            <div class="table-filter"> 
                <a-input-search :style="{ width: '320px' }" placeholder="Please enter something" search-button />
            </div>
        </div>
        <!-- 文章数据展示 -->
        <a-table   :loading="listBlogLoadding" :data="data.items" >  
            <template #columns>
                <a-table-column title="标题" data-index="title"></a-table-column>
                <a-table-column title="作者" data-index="author"></a-table-column>
                <a-table-column title="状态" data-index="status"></a-table-column>
                <a-table-column title="编号" data-index="id"></a-table-column>
                <a-table-column title="创建时间" >
                    <template #cell="{record}"> 
                        {{dayjs.unix(record.created_at).format('YYYY-MM-DD HH:mm:ss')}}
                    </template>
                </a-table-column>
                <a-table-column title="更新时间" >
                    <template #cell="{record}"> 
                        {{dayjs.unix(record.updated_at).format('YYYY-MM-DD HH:mm:ss')}}
                    </template>
                </a-table-column>                
                <a-table-column :width="300" align="center" title="操作"  >
                    <template #cell="{record}"> 
                        <a-button type="text">
                                <icon-edit />
                            编辑
                        </a-button>
                        <a-button type="text">
                            <icon-send />
                            发布
                        </a-button>
                        <a-button type="text" status="danger">
                            <icon-delete />
                            删除
                        </a-button>                                                
                    </template>
                </a-table-column>
            </template>
        </a-table>
    </div>
</template>

<script setup>

import { onMounted,ref } from 'vue';
import { LIST_BLOG } from '@/api/vblog';
import dayjs from 'dayjs'
//界面有关系

const data = ref({total:0,items:[]})
const listBlogLoadding = ref(false)
const ListBlog = async ()=> {
    try {
        listBlogLoadding.value=true
        data.value =await LIST_BLOG()
    }finally{
        listBlogLoadding.value=false
    }
}


// 声明时候加载数据
onMounted(()=> {
    ListBlog()
})
//通过API查询当前文章列表
// const columns = [
//     {
//     title: '编号',
//     dataIndex: 'id',
//     },
//     {
//     title: '时间',
//     dataIndex: 'created_at',
//     },
//     {
//     title: '标题',
//     dataIndex: 'title',
//     },        
// ];
// const data = reactive([{
//     key: '1',
//     name: 'Jane Doe',
//     salary: 23000,
//     address: '32 Park Road, London',
//     email: 'jane.doe@example.com'
// }, {
//     key: '2',
//     name: 'Alisa Ross',
//     salary: 25000,
//     address: '35 Park Road, London',
//     email: 'alisa.ross@example.com'
// }, {
//     key: '3',
//     name: 'Kevin Sandra',
//     salary: 22000,
//     address: '31 Park Road, London',
//     email: 'kevin.sandra@example.com'
// }, {
//     key: '4',
//     name: 'Ed Hellen',
//     salary: 17000,
//     address: '42 Park Road, London',
//     email: 'ed.hellen@example.com'
// }, {
//     key: '5',
//     name: 'William Smith',
//     salary: 27000,
//     address: '62 Park Road, London',
//     email: 'william.smith@example.com'
// }]);




</script>
<style lang="css" scoped>

.table-line{
    display: flex;
    justify-content: space-between;
    margin-bottom: 6px;
}
.table-op{

}
.table-filter{

}
</style>
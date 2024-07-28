<template>
    <div>
        <!-- 页头 -->
        <a-breadcrumb style="margin-bottom: 12px;">
            <a-breadcrumb-item>文章列表</a-breadcrumb-item>
        </a-breadcrumb>
        <!-- 数据创建和筛选 -->
        <div class="table-line">
            <div class="table-filter"> 
                <a-input-search class="kw-search" @search="handleSearch" placeholder="Please enter something" search-button />
            </div>
        </div>
        <!-- 文章数据展示 -->
        <!-- <a-table :loading="listBlogLoadding" :data="data.item"></a-table>  -->
        <!-- <a-table   :loading="listBlogLoadding" :data="data.items"  :pagination="pagination" @page-change="handlePageChange" @page-size="handlePageSizeChange">   -->
        <a-table :loading="listBlogLoadding" :data="data.items" :pagination="pagination" @page-change="handlePageChange"
            @page-size-change="handlePageSizeChange">

            <template #columns>
                <a-table-column title="标题" data-index="title"></a-table-column>
                <a-table-column title="作者" data-index="author"></a-table-column>
                <a-table-column title="状态" data-index="status">
                    <template #cell="{ record }">
                        {{ STATUS_MAP[record.status] }}
                    </template>
                </a-table-column>
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
                    <template #cell="{ record }"> 
                        <a-button type="text" @click="$router.push({name:'BackendBlogEdit',query:{id:record.id}})">
                                <icon-edit />
                            编辑
                        </a-button>
                        <a-popconfirm ok-text :content="`发布【${record.title}】这篇文章?`" type="success" 
                        :ok-loading="updatestatusLoadding"  @ok="handleUpdatestatus(record)">
                            <a-button type="text" >
                            <icon-send />
                            发布
                            </a-button> 
                        </a-popconfirm>
                        <a-popconfirm :content="`是否确认要删除【${record.title}】这篇文章?`" type="error" 
                        :ok-loading="deleteLoadding" @ok="handleDelete(record)">
                            <a-button type="text" status="danger" >
                            <icon-delete />
                            删除
                            </a-button> 
                        </a-popconfirm>
                                               
                    </template>
                </a-table-column>
            </template>
        </a-table>
    </div>
</template>

<script setup>

import { onMounted,ref,reactive } from 'vue';
import { LIST_BLOG,DELETE_BLOG,UPDATE_BLOG_STATUS } from '@/api/vblog';
import dayjs from 'dayjs'
import { Notification } from '@arco-design/web-vue';
import router from '@/router';
//界面有关系
//分页

const data = ref({total:0,items:[]})

const pagination = reactive({
    total: data.value.total,
    showTotal: true,
    showJumper: true,
    showMore: true,
    showPageSize: true,
    current: 1,
    pageSize: 10,
})
const listBlogLoadding = ref(false)
const ListBlog = async ()=> {
    try {
        listBlogLoadding.value=true
        data.value = await LIST_BLOG({
            keywords: params.value,
            page_size: pagination.pageSize,
            page_number: pagination.current
        })
        pagination.total = data.value.total

    }finally{
        listBlogLoadding.value=false
    }
}

const STATUS_MAP = {
    0:'草稿',
    1:'已发布'
}

//搜索参数
const params = ref({
    keywords:'',
    page_size: pagination.PageSize,
    page_number:pagination.current
})

// pageNumber有变化, 重新请求数据
const handlePageChange = (v) => {
    pagination.current = v
    ListBlog()
}

// PageSize发送变化, 重新请求数据
const handlePageSizeChange = (v) => {
    pagination.pageSize = v
    ListBlog()
}

const handleSearch = (v) => {
    params.value.keywords = v
    ListBlog()
}

const deleteLoadding =ref(false)
const handleDelete =async (v) => {
    try {
        deleteLoadding.value =true
        await DELETE_BLOG(v.id)
        Notification.info(`文章【${v.title}】删除成功`)
        //重新刷新页面
        // LIST_BLOG()
        router.push({name:'BackendBlogList'})

    }finally{
        deleteLoadding.value =false
    }
}

const status = {
    "status":1
}
const updatestatusLoadding =ref(false)
const handleUpdatestatus =async (v) => {
    try {
        updatestatusLoadding.value =true
        await UPDATE_BLOG_STATUS(v.id,status)
        Notification.success(`文章【${v.title}】发布成功`)
        //重新刷新页面
        // LIST_BLOG()
        

    }finally{
        updatestatusLoadding.value =false
    }
    router.push({name:'BackendBlogList'})
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
.kw-search {
    width: 400px;
}

</style>
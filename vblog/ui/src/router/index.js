import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'LoginPage',
      //没有嵌套，这个就是最终页面
      component: ()=> import('@/views/login/LoginPage.vue'),
    },
    {
      path: '/backend',
      name: 'BackendLayout',
      //没有嵌套，这个就是最终页面
      component: ()=> import('@/views/backend/BackendLayout.vue'),
      redirect:{name:'BackenddBlogList'},
      children:[
        {
          // /backend/vblogs
          path:'blogs/list',
          name: 'BackendBlogList',
          component: ()=> import('@/views/backend/blogs/ListPage.vue'),
          
        },
        {
          // /backend/vblogs
          path: 'blogs/edit',
          name: 'BackendBlogEdit',
          component: ()=> import('@/views/backend/blogs/EditPage.vue'),

        }
      ]
    },
    {
      path: '/frontend',
      name: 'FrontendLayout',
      //没有嵌套，这个就是最终页面
      component: ()=> import('@/views/frontend/FrontendLayout.vue'),
      redirect:{name:'FrontendBlogList'},
      children: [
        {
          // 
          path: 'blogs/list',
          name: 'FrontendBlogList',
          component: ()=> import('@/views/frontend/blogs/ListPage.vue')
        },
        {
          // 
          path: 'blogs/detail',
          name: 'FronendBlogDetail',
          component: ()=> import('@/views/frontend/blogs/DetailPage.vue')
        }
      ]
    }

  ]
})

export default router

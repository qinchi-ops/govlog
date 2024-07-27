import { createRouter, createWebHistory } from 'vue-router'
import app from '@/stores/app'
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
      redirect:{name:'BackendBlogList'},
      beforeEnter:()=>{
      // 怎么确认用户当前有没有登录喃?
      // 如果中断直接返回你要去向的页面
        if (!app.value.token){
          return {name:'LoginPage'}
        }

      },
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

        },
        {
          // /backend/vblogs
          path: 'comments/list',
          name: 'BackendCommentList',
          component: ()=> import('@/views/backend/comment/ListPage.vue'),
        },
        {
          // /backend/vblogs
          path: 'tags/list',
          name: 'BackendTagList',
          component: ()=> import('@/views/backend/tags/ListPage.vue'),
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
    },
    {
      path:'/:pathMatch(.*)*',
      name: 'NotFound',
      component: ()=> import('@/views/common/NotFound.vue')
    }

  ]
})

export default router
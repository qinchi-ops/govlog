<template>
  <div class="about">
    <table border="1">
      <tbody>
        <tr>
          <th>姓名</th>
          <th>年龄</th>
          <th>性别</th>
        </tr>        
        <tr>
          <th>{{person.name}}</th>
          <th>{{person.age}}</th>
          <th>{{person.profile.gender}}</th>
        </tr>
      </tbody>
    </table>
    <h1>View -> Module </h1>
    <div>用户名: <input type="text" id = "username" v-model="username"  /></div>
    <div>密  码: <input type="password" v-model="password"  /></div>
  </div>
</template>

<script setup>
import {reactive,onBeforeMount,onMounted,onBeforeUpdate,onUpdated,onBeforeUnmount,onUnmounted,ref,nextTick} from 'vue'
const person = reactive({name:'张山', age : 30,profile:{gender:'female'}})
// const loginReq = reactive({username:'', password : ''})
// person.name
const username=ref('init')
const password=ref('')
// reactive({value:v})
username.value = 'test'
onBeforeMount(() =>{
  console.log('onBeforeMount hook')
})
onMounted(() =>{
  console.log('onMounted hook')
  username.value ='test01'
  //在下次更新后，执行这段代码
  nextTick(() =>{
    console.log(document.getElementById('username').value)
  })
  // console.log(document.getElementById('username').value)
})
onBeforeUpdate(() =>{
  console.log('onBeforeUpdate hook')
})
onUpdated(() =>{
  console.log('onUpdated hook')
  console.log(document.getElementById('username').value)
})
onBeforeUnmount(() =>{
  console.log('onBeforeUnmount hook')
})
onUnmounted(() =>{
  console.log('onUnmounted hook')
})
</script>
<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>

<template>
  <div class="about">
    <h1>Model -> View</h1>
    <br />
    <table border="1">
      <thead>
        <tr>
          <th>姓名</th>
          <th>年龄</th>
          <!-- <th>性别</th> -->
        </tr> 
      </thead>
      <tbody>     
        <tr>
          <th>{{person.name}}</th>
          <th>{{person.age}}</th>
          <!-- <th>{{person.profile.gender}}</th> -->
        </tr>
      </tbody>
    </table>
    <!-- <h1>View -> Module </h1> -->
    <div>用户名: <input type="text" id = "username" v-model="username"  /></div>
    <div>密  码: <input type="password" v-model="password"  /></div>
    <!-- innerText -->
    <!-- <div>{{testID}}</div> -->
    <!-- 元素属性 -->
    <!-- ref 类型的变量，vue模版系统会自动解开 -->
    <!-- v-bind:id == :id -->
    <!-- <div v-bind="testID">{{testID}}</div> -->
    <!-- <div :id="testId">{{ testId }}</div> -->

    <!-- 元素方法 -->
    <!-- onClick -- v-on:click -->
    <!-- v-on: == @ -->
    <div style="color: red;"><button @click="handleClieck">按钮</button></div>

    <!-- 元素样式 -->
    <button :style="buttonStyle" @click="handleClieck">按钮</button>

    <div style="height: 40px;width: 40vw;background-color: bisque;" :class="divClass"></div>

    <!-- 表单绑定，用户输入绑定 -->
    <!-- v-model(双向绑定): input value -- name --- input value -->
    <input v-model="name" v-focus type="text"  @keydown.enter="pressEnter">

    <label for="pet-select">Choose a pet:</label>

    <select v-model="choiceValue" name="pets" id="pet-select">
      <!-- <option value="">--Please choose an option--</option>
      <option value="dog">Dog</option>
      <option value="cat">Cat</option>
      <option value="hamster">Hamster</option>
      <option value="parrot">Parrot</option>
      <option value="spider">Spider</option>
      <option value="goldfish">Goldfish</option> -->
      <option v-for="(item) in petOptions" :key="item.value" :value="item.value">{{item.label}}</option>

    </select>
    <div v-if="choiceValue === 'dog'">dog food</div>
    <div v-else-if="choiceValue === 'cat'">cat food</div>
    <div v-else>other food</div>

  </div>
</template>

<script setup>
import {reactive,onBeforeMount,onMounted,onBeforeUpdate,onUpdated,onBeforeUnmount,onUnmounted,ref,nextTick} from 'vue'
const person = reactive({name:'张山', age : 30,profile:{gender:'female'}})
// const loginReq = reactive({username:'', password : ''})
// person.name
const username=ref('init')
const password=ref('')
// const testID=ref('test')

// js sytle对象
const buttonStyle = reactive({ backgroundColor: 'grey' })

const handleClieck = () => {
  console.log('handleClieck')
  buttonStyle.backgroundColor = 'yellow'
  divClass.push('bg')
}
const divClass = reactive([])

const pressEnter = () => {
  console.log('press enter')
}
const choiceValue =ref('')


const petOptions = reactive([
  {label:'--Please choose an option--',value:''},
  {label:'Dog',value:'dog'},
  {label:'Cat',value:'cat'},
  {label:'Goldfish',value:'goldfish'},

])
// const vFocus = {
//   mounted: (el) => el.focus()
// }

// vFocus ---> v-focus
const vFocus = {
  mounted: (el) => { el.focus() }
}

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
.bg{
  background-color: aqua !important;
}
</style>

<template>
    <div>
        <a-breadcrumb style="margin-bottom: 12px;">
        <a-breadcrumb-item>文章管理</a-breadcrumb-item>
        <a-breadcrumb-item>标签列表</a-breadcrumb-item>
        </a-breadcrumb>    
    </div>



    <a-space wrap>
    <div v-for="(tag, index) of tags" :key="index" class="tag-container">
      <a-tag :color="colors[index % colors.length]" closable>{{ tag }}</a-tag>
    </div>
    </a-space>
    <!-- <a-space wrap>
    <a-tag v-for="(tag, index) of tags" :key="index" :tag="tag" closable>{{ tag }}</a-tag>
    <a-tag v-for="(color, index) of colors" :key="index" :color="color" closable>{{ color }}</a-tag> -->


    <!-- <a-tag v-for="(color, index) of colors" :key="index" :color="color" closable>{{ color }}</a-tag> -->
    <!-- </a-space> -->
    <!-- <h3>Custom Color:</h3>
    <a-space wrap>
        <a-tag v-for="(color, index) of custom" :key="index" :color="color" closable>{{ color }}</a-tag>
    </a-space> -->

</template>

<script setup>

const colors = [
      'red',
      'orangered',
      'orange',
      'gold',
      'lime',
      'green',
      'cyan',
      'blue',
      'arcoblue',
      'purple',
      'pinkpurple',
      'magenta',
      'gray'
    ];
const tags = [
    'tag1',
    'tag2',
    'default',
    'complete'

]    

</script>

<style lang="css" scoped>
.tag-container {
  display: flex;
  align-items: center;
  margin-bottom: 5px;
}
</style>
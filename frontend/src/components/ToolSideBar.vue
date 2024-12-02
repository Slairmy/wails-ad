<template>
  <el-row class="tac">
    <el-col :span="24">
      <el-menu
          default-active="0"
          class="el-menu-vertical-demo"
      >
        <el-menu-item index="1" @click="selectMenu('Archery2')">
          <template #title>
            <el-icon><location /></el-icon>
            <span>Archery SQL查询</span>
          </template>
        </el-menu-item>
        <el-menu-item index="2" @click="selectMenu('')">
          <el-icon><icon-menu /></el-icon>
          <span>其他待开放</span>
        </el-menu-item>
      </el-menu>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import {
  Menu as IconMenu,
  Location,
} from '@element-plus/icons-vue'
import { defineEmits } from 'vue'
// 在 Vue 3 中，使用 setup 语法时，你不能直接使用 this 关键字来访问组件的实例和数据
// 定义 emit 的类型
const menuClickEmit = defineEmits<{
  (e: 'menu-click', menuName: string): void;
}>()

const selectMenu = (menuName) => {
  menuClickEmit('menu-click', menuName)
}
</script>


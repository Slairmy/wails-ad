<template>
  <div class="common-layout">
    <el-container>
      <el-header>
        <tool-header />
      </el-header>
      <el-container>
        <el-aside width="200px">
          <tool-side-bar @menu-click="handleMenuClick" />
        </el-aside>
        <el-main>
          <!-- 动态加载组件 -->
          <component :is="currentComponent"></component>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import ToolHeader from "./ToolHeader.vue";
import ToolSideBar  from "./ToolSideBar.vue";
import Archery from "../navigation/archery/Archery.vue";
import Archery2 from "../navigation/archery/Archery2.vue";
import Welcome from "../navigation/welcome/Welcome.vue";

export default {
  name: 'Home',
  components: {
    ToolHeader,
    ToolSideBar,
    Welcome,
    Archery,
    Archery2
  },
  data() {
    return {
      currentComponent: 'Welcome'
    }
  },
  methods: {
    handleMenuClick(componentName) {
      console.log(componentName)
      if (componentName === '') {
        this.currentComponent = Welcome
        return
      }
      this.currentComponent = componentName
    }
  }
}
</script>

<template>
  <div class="home-container">
    <side-bar @menu-click="handleMenuClick" />
    <div class="content-area">
      <component :is="currentComponent" />
    </div>
  </div>
</template>

<script>
import SideBar from './SideBar.vue'
import Archery from "./Archery.vue";

export default {
  name: 'Home',
  components: {
    SideBar,
    Archery
  },
  data() {
    return {
      currentComponent: 'Archery'
    }
  },
  methods: {
    handleMenuClick(componentName) {
      this.currentComponent = componentName
    }
  }
}
</script>

<style scoped>
.home-container {
  display: flex;
  height: 100vh;
  background-color: white;
}

.content-area {
  flex: 1;
  padding: 20px;
}
</style>
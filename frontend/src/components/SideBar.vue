<template>
  <div class="sidebar">
    <div
        v-for="item in menuItems"
        :key="item.id"
        class="menu-item"
        :class="{ active: currentItem === item.component }"
        @click="selectMenuItem(item.component)"
        role="button"
        tabindex="0"
    >
      <span>{{ item.title }}</span>
    </div>
  </div>
</template>

<script>
export default {
  name: 'SideBar',
  data() {
    return {
      currentItem: 'Archery',
      menuItems: [
        { id: 1, title: 'Archery', component: 'Archery' }
      ]
    }
  },
  methods: {
    selectMenuItem(componentName) {
      if (!componentName) return
      this.currentItem = componentName
      this.$emit('menu-click', componentName)
    }
  }
}
</script>

<style scoped>
.sidebar {
  width: 200px;
  color: white;
  padding: 20px 0;
  user-select: none;
}

.menu-item {
  padding: 15px 20px;
  cursor: pointer;
  transition: background-color 0.3s;
  display: flex;
  align-items: center;
}

.menu-item:hover {
  background-color: #34495e;
}

.menu-item.active {
  background-color: #3498db;
}

.menu-item:focus {
  outline: 2px solid #3498db;
  outline-offset: -2px;
}
</style>
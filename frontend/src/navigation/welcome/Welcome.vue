<template>
  <div class="welcome">
    <span>Welcome to Ad-Tools</span>
  </div>
  <div>
    <el-select
        v-model="value"
        placeholder="Select"
        size="large"
        style="width: 240px"
    >
      <el-option
          v-for="item in options"
          :key="item.value"
          :label="item.label"
          :value="item.value"
      />
    </el-select>
  </div>


</template>

<script>
export default {
  name: "Welcome",
  data() {
    return {
      value: "aaa",
      options: [
        {
          value: 'Option5',
          label: 'Option5',
        }
      ]
    }
  }
}
</script>

<!--<script lang="ts" setup>-->
<!--import { ref } from 'vue'-->

<!--const value = ref('')-->

<!--const options = [-->
<!--  {-->
<!--    value: 'Option1',-->
<!--    label: 'Option1',-->
<!--  },-->
<!--  {-->
<!--    value: 'Option2',-->
<!--    label: 'Option2',-->
<!--  },-->
<!--  {-->
<!--    value: 'Option3',-->
<!--    label: 'Option3',-->
<!--  },-->
<!--  {-->
<!--    value: 'Option4',-->
<!--    label: 'Option4',-->
<!--  },-->
<!--  {-->
<!--    value: 'Option5',-->
<!--    label: 'Option5',-->
<!--  },-->
<!--]-->
<!--</script>-->

<style scoped>

.welcome {
  font-size: 24px;
  align-items: center;
  text-align: center;
}

</style>

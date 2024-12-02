<template>
  <query />
</template>

<script>
import AceEditor from "./editor/AceEditor.vue"
import Query from "./Query.vue";
import {SetCookie, GetCookie} from "../../../wailsjs/go/archery/Archery";

export default {
  name: 'Archery2',
  components: {
    Query
  },
  data() {
    return {
      sqlContent: '-- 在这里输入你的代码',
      queryResult: null,
      departmentId: null,
    }
  },
  methods: {
    executeQuery() {
      const sql = this.$refs.sqlEditor.editor.getValue()
      console.log(sql)
      this.queryResult = `sql语句: \n ${sql}`
    },
    handleEditorChange(value) {
      this.sqlContent = value
    },
    handleInputChange(value) {
      this.departmentId = value
    },
    handleBlur() {

      SetCookie("name","slairmy").then(() => {
        console.log("aaaaaa")
        GetCookie().then((result) => {
          console.log("输出get cookie")
          console.log(result)
        })
      })
    }
  },
  // 挂载一下自定义的提示词
  // mounted() {
  //   // 添加新表和列
  //   this.$refs.sqlEditor.addTable('profiles_2', [
  //     'id',
  //     'user_name',
  //     'phone',
  //     'address'
  //   ])
  // }
}
</script>

<style scoped>

.archery-container {
  border: 2px solid red;
  padding-left: 40px;
  padding-top: 20px;
  padding-right: 40px;
}


.archery-header {
  border: 2px solid yellow;
  font-size: 16px;
  text-align: left;
  height: 50px;
  padding-left: 20px;
  display: flex;
  align-items: center;
}

.editor-controls select {
  margin-right: 10px;
  padding: 5px 10px;
  border-radius: 4px;
}

.sql-area {
  display: grid;
  grid-template-columns: 70% 30%; /* 左右两部分占比 */

  border: 2px solid yellow;
}

.editor-wrapper {
  border-radius: 4px;
  overflow: hidden;
}

</style>
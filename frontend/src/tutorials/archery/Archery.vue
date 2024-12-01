<template>
  <div class="archery-container">

    <div>
      <div class="archery-header">
        <span>SQL查询</span>
      </div>
      <div class="sql-area">

        <div class="editor-wrapper">
          <ace-editor
              v-model="sqlContent"
              ref="sqlEditor"
              @change="handleEditorChange"
          />
        </div>
        <div class="operator-wrapper">

          <el-input v-model="departmentId" style="width: 240px" placeholder="请输入企业id" />
          <el-button type="success">SQL查询</el-button>

        </div>
      </div>

    </div>

  </div>
</template>

<script>
import AceEditor from "./editor/AceEditor.vue"


export default {
  name: 'Archery',
  components: {
    AceEditor
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
    }
  },
  // 挂载一下自定义的提示词
  mounted() {
    // 添加新表和列
    this.$refs.sqlEditor.addTable('profiles_2', [
      'id',
      'user_name',
      'phone',
      'address'
    ])
  }
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
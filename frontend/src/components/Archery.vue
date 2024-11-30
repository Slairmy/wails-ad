<template>
  <div class="function-container">
    <h2>代码编辑器</h2>
    <div class="editor-wrapper">
      <ace-editor
          v-model="sqlContent"
          ref="sqlEditor"
          @change="handleEditorChange"
      />
    </div>
    <div class="editor-controls">
      <button class="execute-btn" @click.prevent="executeQuery">执行查询</button>
    </div>

    <div v-if="queryResult" class="query-result">
      <pre>{{ queryResult }}</pre>
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
      queryResult: null
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
.function-container {
  padding: 20px;
}

.editor-controls {
  margin-bottom: 20px;
}

.editor-controls select {
  margin-right: 10px;
  padding: 5px 10px;
  border-radius: 4px;
}

.editor-wrapper {
  border-radius: 4px;
  overflow: hidden;
}

.query-result {
  padding: 10px;
  background-color: #f5f5f5;
  border-radius: 4px;
}

.query-result pre {
  margin: 0;
  white-space: pre-wrap;
  word-wrap: break-word;
  color: rebeccapurple;
}
</style>
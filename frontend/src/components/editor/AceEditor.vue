<template>
  <div :id="editorId" class="ace-editor"></div>
</template>

<script>
import ace from 'ace-builds'
import 'ace-builds/src-noconflict/mode-sql'
import 'ace-builds/src-noconflict/snippets/sql'
import 'ace-builds/src-noconflict/theme-monokai'
import 'ace-builds/src-noconflict/ext-language_tools'

export default {
  name: 'AceEditor',
  props: {
    value: {
      type: String,
      default: ''
    },
    lang: {
      type: String,
      default: 'sql'
    },
    theme: {
      type: String,
      default: 'chrome'
    },
    height: {
      type: String,
      default: '400px'
    }
  },
  watch: {
    value(newVal) {
      if (this.editor && this.editor.getValue() !== newVal) {
        this.editor.setValue(newVal, 1)
      }
    }
  },
  data() {
    return {
      /** @type {import('ace-builds').Ace.Editor|null} */
      editor: null,
      editorId: `ace-editor-${Math.random().toString(36).substr(2, 9)}`,
      tableSchema: {
        'profiles_1': ['id', 'name', 'email', 'created_at'],
        'campaigns_1': ['id', 'title', 'status', 'start_date', 'end_date'],
        'products_1': ['id', 'name', 'price', 'category']
      }
    }
  },
  mounted() {
    this.initEditor()
  },
  methods: {
    initEditor() {

      this.editor = ace.edit(this.editorId)
      this.editor.setTheme(`ace/theme/${this.theme}`)
      this.editor.session.setMode(`ace/mode/${this.lang}`)
      this.editor.setValue(this.value, 1)

      this.editor.setOptions({
        enableBasicAutocompletion: true,
        enableLiveAutocompletion: true,
        enableSnippets: true,
        showLineNumbers: true,
        tabSize: 2,
        fontSize: 14
      })

      this.setupCompleter()

      this.editor.on('change', () => {
        const content = this.editor.getValue()
        this.$emit('input', content)
        this.$emit('change', content)
      })
    },

    setupCompleter() {

      const langTools = ace.require('ace/ext/language_tools')

      const completer = {
        getCompletions: (editor, session, pos, prefix, callback) => {
          const line = session.getLine(pos.row)
          const textBeforeCursor = line.slice(0, pos.column).toLowerCase()

          if (/from\s+$/i.test(textBeforeCursor) || /from\s+\w*$/i.test(textBeforeCursor)) {
            const tables = Object.keys(this.tableSchema)
            callback(null, tables.map(table => ({
              caption: table,
              value: table,
              meta: 'table',
              score: 100
            })))
            return
          }

          const tableMatch = textBeforeCursor.match(/from\s+(\w+)/i)
          if (tableMatch) {
            const tableName = tableMatch[1]
            const columns = this.tableSchema[tableName] || []

            if (/select\s+\w*$/i.test(textBeforeCursor)) {
              callback(null, columns.map(column => ({
                caption: column,
                value: column,
                meta: 'column',
                score: 90
              })))
              return
            }

            if (/where\s+\w*$/i.test(textBeforeCursor)) {
              callback(null, columns.map(column => ({
                caption: `${tableName}.${column}`,
                value: `${tableName}.${column}`,
                meta: 'column',
                score: 90
              })))
              return
            }
          }

          callback(null, [])
        }
      }

      langTools.addCompleter(completer)
    },

    addTable(tableName, columns = []) {
      if (!this.tableSchema[tableName]) {
        this.tableSchema[tableName] = columns
      }
    }
  },
  beforeDestroy() {
    if (this.editor) {
      this.editor.destroy()
    }
  }
}
</script>

<style scoped>
.ace-editor {
  width: 100%;
  height: v-bind(height);
  border: 1px solid #ccc;
}
</style>
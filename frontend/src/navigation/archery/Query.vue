<template>
  <div class="common-layout">
    <el-container>
      <el-main  style="padding-left: 0; padding-top: 0;">
        <div class="common-layout">
          <el-container>
            <el-main>
              <ace-editor
                  ref="sqlEditor"
              />
            </el-main>
            <el-aside width="320px">

              <el-row :gutter="20">
                <div style="margin-left: 15px; margin-top: 20px">
                  <el-input style="width: 300px" size="large" placeholder="输入企业id" />
                </div>
              </el-row>
              <el-row :gutter="20">
                <div style="margin-left: 15px; margin-top: 20px; ">
                  <el-select
                      v-model="selectedOption"
                      size="large"
                      style="width: 300px"
                  >
                    <el-option
                        v-for="item in selectOptions"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                    />
                  </el-select>
                </div>
              </el-row>

              <el-row :gutter="20">
                <div style="margin-left: 15px; margin-top: 20px; ">
                  <el-button type="primary">美化</el-button>
                  <el-button type="warning">执行计划</el-button>
                  <el-button type="success">SQL查询</el-button>
                </div>

              </el-row>

              <el-row :gutter="20">

                <div style="margin-left: 15px; margin-top: 60px;">
                  <el-table :data="tableData" border style="font-size: 15px">
                    <el-table-column prop="instance" label="实例" width="150" />
                    <el-table-column prop="database" label="数据库" width="90" />
                    <el-table-column prop="index" label="分表" width="60"  />
                  </el-table>
                </div>
              </el-row>
            </el-aside>
          </el-container>
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import AceEditor from "./editor/AceEditor.vue"

export default {
  name: "Query",
  components: {
    AceEditor,
  },
  data() {
    return {
      selectedOption: "Option1",
      selectOptions: [
        {
          value: "Option1",
          label: "Option1"
        },
        {
          value: "Option2",
          label: "Option2"
        },
        {
          value: "Option3",
          label: "Option3"
        }
      ],
      tableData: [
        {
          instance: 'newad-prd-19',
          database: 'apadata',
          index: '1',
        }
      ]
    }
  }
}
</script>


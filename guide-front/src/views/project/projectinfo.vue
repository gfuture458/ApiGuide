<template>
  <el-container>
    <MyHeader></MyHeader>
    <el-main>
      <Left></Left>
      <div class="right">
        <el-page-header content="项目详情" style="padding: 20px;padding-bottom: 0" @back="back()"></el-page-header>
        <el-row>
          <span style="margin-left: 20px; align-items: center;display: flex">模块列表</span>
        </el-row>
        <div style="margin: 20px;">
          <div v-for="(m, index) in modular" :key="m" class="text item" style="padding-bottom: 10px">
            <el-collapse accordion>
              <el-collapse-item :title="m.name" :name="index">
                <el-table :border="true" :data="m.interface">
                  <el-table-column label="接口名称" prop="name"></el-table-column>
                  <el-table-column label="路由" prop="path" >
                    <template slot-scope="{row}">
                      <div v-if="!row.edit">{{row.path}}</div>
                      <el-input v-else v-model="row.path"></el-input>
                    </template>
                  </el-table-column>
                  <el-table-column label="请求方式" prop="method"></el-table-column>
                  <el-table-column label="是否必填" prop="require"></el-table-column>
                  <el-table-column label="示例" prop="example"></el-table-column>
                  <el-table-column label="类型" prop="type"></el-table-column>
                  <el-table-column label="描述" prop="desc"></el-table-column>
                  <el-table-column label="操作" width="200">
                    <template slot-scope="{row}">
                      <el-button
                        size="mini"
                        @click="updateRow(row)">修改</el-button>
                      <el-button
                        size="mini"
                        type="danger"
                        @click="delTpis(scope.$index, scope.row)"
                      >删除</el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </el-collapse-item>
            </el-collapse>
          </div>
        </div>
      </div>
    </el-main>
  </el-container>
</template>

<script>
import MyHeader from '../../components/MyHeader'
import Left from '../../components/left'
export default {
  name: 'projectinfo',
  components: {MyHeader, Left},
  data () {
    return {
      modular: [
        {
          name: '模块一',
          version: '2.0',
          prefix: '/user',
          interface: [{
            name: '接口一',
            path: '/user_info',
            method: 'get',
            require: 'true',
            example: '{id: 1}',
            type: 'object',
            edit: false,
            desc: '获取用户详情'
          }]
        },
        {
          name: '模块一',
          version: '2.0',
          prefix: '/user',
          interface: [{
            name: '接口一',
            path: '/user_info',
            method: 'get',
            require: 'true',
            example: '{id: 1}',
            type: 'object',
            desc: '获取用户详情'
          }]
        },
        {
          name: '模块一',
          version: '2.0',
          prefix: '/user',
          interface: [{
            name: '接口一',
            path: '/user_info',
            method: 'get',
            require: 'true',
            example: '{id: 1}',
            type: 'object',
            desc: '获取用户详情'
          }]
        }
      ]
    }
  },
  methods: {
    back () {
      this.$router.push('/project')
    },
    updateRow (row) {
      console.log('----')
      console.log(this.modular)
      row.edit = !row.edit
      console.log(this.modular)
    }
  }
}
</script>

<style scoped>
  @import "../../assets/css/index.css";
</style>

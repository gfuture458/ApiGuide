<template>
  <el-container>
    <MyHeader></MyHeader>
    <el-main>
      <Left></Left>
      <div class="right">
        <el-page-header content="项目详情" style="padding: 20px 20px 0;"></el-page-header>
        <el-row>
          <span style="margin-left: 20px; align-items: center;display: flex">项目详情</span>
        </el-row>
        <el-form>
          <div style="margin: 30px">
            <el-form-item label="项目名称">
              <el-select placeholder="请选择" v-model="value">
                <el-option value="全部">
                </el-option>
                <el-option
                  v-for="item in tableData"
                  :value=item.Name
                  :label=item.Name
                  :key=item.Id
                  >
                </el-option>
              </el-select>
              <el-button type="primary" style="margin-left: 50px">查询</el-button>
              <el-button type="primary" style="float: right; margin-right: 20px" @click="add_project()">新增项目</el-button>
            </el-form-item>
          </div>
        </el-form>
        <el-table style="width: auto; margin: 30px 30px 0" :border="true" :data="tableData">
          <el-table-column label="项目名称" prop="Name"></el-table-column>
          <el-table-column label="域名" prop="Host"></el-table-column>
          <el-table-column label="所属用户" prop="User.Username"></el-table-column>
          <el-table-column label="创建时间" prop="CreateAt"></el-table-column>
          <el-table-column label="操作">
            <template slot-scope="scope">
              <el-button
                size="mini"
                @click="to_detail(scope.row)">详情</el-button>
              <el-button
                size="mini"
                type="danger"
                @click="delTpis(scope.$index, scope.row)"
                >删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-main>
  </el-container>
</template>

<script>
import MyHeader from '../../components/MyHeader'
import Left from '../../components/left'
import request from '../../utils/request'
import apis from '../../config/apis'
export default {
  name: 'project',
  components: {MyHeader, Left},
  data: function () {
    return {
      tableData: [],
      value: ''
    }
  },
  mounted () {
    this.get_all_project()
  },
  methods: {
    add_project () {
      this.$router.push('/addProject')
    },
    get_all_project: async function () {
      let result = await request({
        url: apis.project,
        method: 'get'
      })
      if (result.data.code === 200) {
        this.tableData = result.data.data
      }
    },
    to_detail (pid) {
      this.$router.push('/project_info')
      localStorage.setItem('pid', pid.Id)
      this.$store.dispatch('changePidFun', pid)
    },
    async delete_project (index, obj) {
      let data = {
        id: obj.Id
      }
      let result = await request({
        url: apis.del_project,
        method: 'delete',
        params: data
      })
      if (result.data.code === 200) {
        console.log('delete success ')
        this.tableData.splice(index, 1)
      } else {
        this.$message({
          type: 'error',
          message: '删除失败'
        })
      }
    },
    delTpis (index, row) {
      this.$confirm('将永久删除该项目， 是否继续', '提示', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning',
        center: true
      }).then(() => {
        console.log('tips')
        this.delete_project(index, row)
        this.$message({
          type: 'success',
          message: '删除成功'
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消操作'
        })
      })
    }
  }
}
</script>

<style scoped>
  @import "../../assets/css/index.css";
  /*.el-row {*/
    /*width: auto;*/
    /*background-color: #B3C0D1;*/
    /*margin: 10px;*/
    /*height: 50px;*/
    /*display: flex;*/
  /*}*/
</style>

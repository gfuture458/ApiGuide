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
                @click="to_detail()">详情</el-button>
              <el-button
                size="mini"
                type="danger"
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
      console.log('project', result.data.code)
      if (result.data.code === 200) {
        this.tableData = result.data.data
        console.log(result.data.data)
      } else {
        this.$message.error('错误的请求')
      }
    },
    to_detail () {
      this.$router.push('/project_info')
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

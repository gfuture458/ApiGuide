<template>
  <el-container>
    <MyHeader></MyHeader>
    <el-main>
      <Left></Left>
      <div class="right">
        <el-page-header content="项目详情" style="padding: 20px;padding-bottom: 0"></el-page-header>
        <el-row>
          <span style="margin-left: 20px; align-items: center;display: flex">项目详情</span>
        </el-row>
        <el-form>
          <div style="margin: 30px">
            <el-form-item label="项目名称">
              <el-select placeholder="请选择">
                <el-option
                  value="选项1"
                  label="黄金"
                  >
                </el-option>
              </el-select>
              <el-button type="primary" style="margin-left: 50px">查询</el-button>
              <el-button type="primary" style="float: right; margin-right: 20px">新增项目</el-button>
            </el-form-item>
          </div>
        </el-form>
        <el-table style="width: auto; margin: 30px 30px 0;" border="true" :data="tableData">
          <el-table-column label="项目名称" prop="name"></el-table-column>
          <el-table-column label="域名" prop="host"></el-table-column>
          <el-table-column label="所属用户" prop="username"></el-table-column>
          <el-table-column label="创建时间" prop="created_at"></el-table-column>
          <el-table-column label="操作">
            <template slot-scope="scope">
              <el-button
                size="mini"
                >查看</el-button>
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
export default {
  name: 'project',
  components: {MyHeader, Left},
  data: function () {
    return {
      tableData: [
        {
          name: '项目一',
          host: 'www.baidu.com',
          username: '郭家宏',
          created_at: '2018-08-01'
        }
      ]
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

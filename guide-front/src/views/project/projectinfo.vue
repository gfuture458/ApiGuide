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

    }
  }
}
</script>

<style scoped>
  @import "../../assets/css/index.css";
</style>

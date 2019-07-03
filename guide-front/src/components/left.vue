<template>
  <div class="left">
    <el-col :span="12">
      <el-menu
        class="el-menu-vertical-demo"
        background-color="#545c64"
        text-color="#fff"
        active-text-color="#ffd04b"
        @open="open">
        <el-menu-item @click="to_upload()">
          <i class="el-icon-s-custom"></i>
          <span>用户信息</span>
        </el-menu-item>
        <el-submenu index="1">
          <template slot="title">
            <i class="el-icon-menu"></i>
            <span>项目名称</span>
          </template>
          <label v-for="item in target_json" :key="item.Name">
            <el-menu-item @click="mytest(item.Name)">{{item.Name}}</el-menu-item>
          </label>
        </el-submenu>
      </el-menu>
    </el-col>
  </div>
</template>
<script>
import request from '../utils/request'
import apis from '../config/apis'
// return {
//   target_json: []
// }
export default {
  data () {
    return {
      target_json: []
    }
  },
  mounted () {
    this.get_projects()
  },
  methods: {
    open (key, keypath) {
      console.log(key, keypath)
      this.$router.push('/project')
    },
    mytest (mystr) {
      console.log(mystr)
    },
    to_upload () {
      this.$router.push('/index')
    },
    get_projects: async function () {
      let data = {
        id: this.$store.getters.getuser.Id
      }
      let result = await request({
        url: apis.project,
        method: 'get',
        data: data
      })
      console.log('project', result.data.code)
      if (result.data.code === 200) {
        this.target_json = result.data.data
        console.log('base', result.data.data)
      } else {
        this.$message.error('错误的请求')
        this.$router.push('/')
      }
    }
  }
}
</script>
<style scoped>
  @import "../assets/css/index.css";
</style>

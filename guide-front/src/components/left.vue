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
          <label v-for="item in target_json" :key="item">
            <el-menu-item @click="mytest(item)">{{item}}</el-menu-item>
          </label>
        </el-submenu>
      </el-menu>
    </el-col>
  </div>
</template>
<script>
// import apis from '@/config/apis'
// import request from '@/utils/request'
export default {
  data () {
    return {
      target_json: [
        '项目1',
        '项目2',
        '项目3'
      ]
    }
  },
  methods: {
    open (key, keypath) {
      console.log(key, keypath)
    },
    mytest (mystr) {
      console.log(mystr)
    },
    to_upload () {
      this.$router.push('/index')
    }
  }
}
</script>
<style scoped>
  @import "../assets/css/index.css";
</style>

<template>
  <el-container>
    <MyHeader></MyHeader>
    <el-main>
      <Left></Left>
      <div class="right">
        <h2 style="padding: 20px">用户信息</h2>
        <el-divider></el-divider>
        <div style="height: auto;padding: 20px;">
            <div>
              <label>姓名</label>
              <li>{{user_info.username ? user_info.username: "无姓名"}}</li>
              <label>电话</label>
              <li>{{user_info.email ? user_info.phone: "无电话"}}</li>
              <label>年龄</label>
              <li>{{user_info.age ? user_info.age: "无数据"}}</li>
            </div>
          </div>
        </div>
      </el-main>
    </el-container>
</template>

<script>
import request from '@/utils/request'
import apis from '@/config/apis'
import Left from '../../components/left'
import MyHeader from '../../components/MyHeader'
export default {
  name: 'index',
  components: {MyHeader, Left},
  data: function () {
    return {
      // target_json: [
      //   '项目1',
      //   '项目2',
      //   '项目3'
      // ],
      user_info: {
        username: '',
        email: '',
        age: ''
      }
    }
  },
  mounted () {
    this.mytest2()
  },
  methods: {
    open (key, keypath) {
      console.log(key, keypath)
    },
    async mytest2 () {
      let result = await request({
        url: apis.user_msg,
        method: 'get'
      })
      console.log(result.data)
      if (result.data.code === 200) {
        let response = result.data.data
        this.user_info.username = response.Username
        this.user_info.email = response.Email
        this.user_info.age = response.Age
        this.$store.dispatch('changeFun', response)
      } else {
        this.$router.push('/')
      }
    }
  }
}
</script>

<style scoped>
  @import "../../assets/css/index.css";
</style>

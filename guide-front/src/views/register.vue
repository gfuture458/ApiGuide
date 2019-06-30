<template>
  <div class="hole">
    <el-container>
      <el-header>接口管理平台</el-header>
      <el-main>
        <el-form :model="registerForm" status-icon :rules="rules" ref="registerForm"  class="demo-registerForm">
          <el-form-item label="账号" prop="username">
            <el-input type="text" auto-complete="off" v-model="registerForm.username"></el-input>
          </el-form-item>
          <el-form-item label="密码" prop="password">
            <el-input type="password" auto-complete="off" v-model="registerForm.password"></el-input>
          </el-form-item>
          <el-form-item label="确认密码" prop="repassword">
            <el-input type="password" auto-complete="off" v-model="registerForm.repassword"></el-input>
          </el-form-item>
          <el-form-item class="register-button">
            <el-button type="primary" @click="submitForm('registerForm')">提交</el-button>
            <el-button @click="ToLogin()">登陆</el-button>
          </el-form-item>
        </el-form>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import request from '@/utils/request'
import apis from '@/config/apis'
export default {
  data () {
    var validateUsername = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('账号不能为空'))
      }
      setTimeout(() => {
        if (value.length < 5 || value.length > 12) {
          callback(new Error('长度错误'))
        } else {
          callback()
        }
      }, 10)
    }
    var validatePassword = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('密码不能为空'))
      }
      setTimeout(() => {
        if (value.length < 5 || value.length > 20) {
          callback(new Error('请输入长度5-20字符的密码'))
        } else {
          callback()
        }
      }, 10)
    }
    var validateRePassword = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('重复密码不能为空'))
      }
      setTimeout(() => {
        if (value !== this.registerForm.password) {
          callback(new Error('两次密码不一致'))
        } else if (value.length !== this.registerForm.password.length) {
          return callback(new Error('两次密码长度不一致'))
        } else {
          callback()
        }
      }, 10)
    }
    return {
      registerForm: {
        username: '',
        password: '',
        repassword: ''
      },
      rules: {
        username: [
          { validator: validateUsername, trigger: 'blur' }
        ],
        password: [
          { validator: validatePassword, trigger: 'blur' }
        ],
        repassword: [
          { validator: validateRePassword, trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    submitForm (formName) {
      this.$refs[formName].validate(async (valid) => {
        if (valid) {
          let data = {
            username: this.registerForm.username,
            password: this.registerForm.password
          }
          let result = await request({
            url: apis.register,
            method: 'post',
            data: data
          })
          if (result.data.code === 200) {
            this.$message(result.data.msg)
            this.$router.push('/')
          } else {
            this.$message.error(result.data.msg)
          }
        } else {
          this.$message.error('错误的操作')
        }
      })
    },
    ToLogin () {
      this.$router.push('/')
    }
  }
}
</script>

<style scoped>
  @import "../assets/css/main.css";
</style>

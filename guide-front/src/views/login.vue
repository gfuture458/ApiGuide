<template>
  <div class="hole">
    <el-container>
      <el-header>接口管理平台</el-header>
      <el-main>
        <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm"  class="demo-ruleForm">
          <el-form-item label="账号" prop="username">
            <el-input type="text" v-model="ruleForm.username" auto-complete="off"></el-input>
          </el-form-item>
          <el-form-item label="密码" prop="password">
            <el-input type="password" v-model="ruleForm.password" auto-complete="off"></el-input>
          </el-form-item>
          <el-form-item class="login-button">
            <el-button type="primary" @click="submitForm('ruleForm')">提交</el-button>
            <el-button @click="resetForm('ruleForm')">重置</el-button>
            <el-button @click="ToRegister()">注册</el-button>
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
        if (value.length > 12 || value.length < 6) {
          callback(new Error('账号长度错误'))
        } else {
          callback()
        }
      }, 10)
    }
    var validatePassword = (rule, value, callback) => {
      if (!value) {
        callback(new Error('请输入密码'))
      } else {
        callback()
      }
    }
    return {
      ruleForm: {
        username: '',
        password: ''
      },
      rules: {
        username: [
          { validator: validateUsername, trigger: 'blur' }
        ],
        password: [
          { validator: validatePassword, trigger: 'blur' }
        ]
      }
    }
  },
  mounted () {
  },
  methods: {
    submitForm (formName) {
      this.$refs[formName].validate(async (valid) => {
        if (valid) {
          let Base64 = require('js-base64').Base64
          let data = {
            username: this.ruleForm.username,
            password: Base64.encode(this.ruleForm.password)
          }
          let result = await request({
            url: apis.login,
            method: 'post',
            data: data
          })
          if (result.data.code === 200) {
            this.$message(result.data.msg)
            this.$router.push('/index')
            return true
          } else {
            this.$message.error(result.data.msg)
            return false
          }
        } else {
          return false
        }
      })
    },
    resetForm (formName) {
      this.$refs[formName].resetFields()
    },
    ToRegister () {
      this.$router.push('/register')
    }
  }
}
</script>

<style scoped>
  @import "../assets/css/main.css";
</style>

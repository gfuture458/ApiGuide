<template>
  <el-container>
    <MyHeader></MyHeader>
    <el-main>
      <Left></Left>
      <div class="right">
        <el-page-header content="新增项目" style="padding: 20px;padding-bottom: 0"></el-page-header>
        <el-row>
          <span style="margin-left: 20px; align-items: center;display: flex">新增项目</span>
        </el-row>
        <div class="form" style="margin-top: 60px">
          <el-form style="width: 60%" :model="projectForm" status-icon :rules="prules" ref="projectForm">
            <el-form-item label="项目名称" prop="name">
              <span><el-input v-model="projectForm.name" auto-complete="off"></el-input></span>
            </el-form-item>
            <el-form-item label="项目地址" prop="host">
              <span><el-input v-model="projectForm.host" auto-complete="off"></el-input></span>
            </el-form-item>
            <el-form-item class="add-button">
              <el-button type="primary" @click="submitForm('projectForm')">提交</el-button>
            </el-form-item>
          </el-form>
        </div>
      </div>
    </el-main>
  </el-container>
</template>

<script>
import MyHeader from '../../components/MyHeader'
import Left from '../../components/left'
import request from '../../utils/request'
import apis from '@/config/apis'
export default {
  name: 'project_add',
  components: {MyHeader, Left},

  data: function () {
    var validatePro = (rule, value, callback) => {
      console.log('ssss')
      if (!value) {
        return callback(new Error('项目名称不能为空'))
      }
      setTimeout(() => {
        if (value.length > 12 || value.length < 1) {
          callback(new Error('账号长度错误，请输入1-12个字符'))
        } else {
          callback()
        }
      }, 10)
    }
    var validateHost = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请正确填写邮箱'))
      } else {
        if (value !== '') {
          let reg = /^(https?:\/\/)([0-9a-z.]+)(:[0-9]+)?([/0-9a-z.]+)?(\?[0-9a-z&=]+)?(#[0-9-a-z]+)?/i
          if (!reg.test(value)) {
            callback(new Error('请输入有效的地址'))
          }
        }
        callback()
      }
    }
    return {
      projectForm: {
        name: '',
        host: ''
      },
      prules: {
        name: [
          {validator: validatePro, trigger: 'blur'}
        ],
        host: [
          {validator: validateHost, trigger: 'blur'}
        ]
      }
    }
  },
  methods: {
    submitForm (formName) {
      this.$refs[formName].validate(async (valid) => {
        if (valid) {
          let data = {
            name: this.projectForm.name,
            host: this.projectForm.host
          }
          let result = await request({
            url: apis.add_pro,
            method: 'post',
            data: data
          })
          if (result.data.code === 200) {
            this.$message(result.data.msg)
            this.$router.push('/project_info')
          }
        }
      })
    }
  }
}
</script>

<style scoped>
  @import "../../assets/css/index.css";
  .el-form-item {
    width: 40%;
  }
  .add-button {
    display: flex;
    justify-content: center;
    padding-top: 20px;
  }
  .form {
    display: flex;
    justify-content: center;
  }
  .el-form {
    display: flex;
    justify-content: center;
    flex-direction: column;
    align-items: center;
  }
</style>

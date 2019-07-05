<template>
  <el-container>
    <MyHeader></MyHeader>
    <el-main>
      <Left></Left>
      <div style="height: 100%; width: 80%">
        <div class="right-top" style="height: 10%;">
          <div style="position: fixed; width: 100%;background-color: white;z-index: 999">
            <el-page-header content="项目详情" style="padding: 20px 20px 0;" @back="back()"></el-page-header>
            <el-row>
              <span style="margin-left: 20px; align-items: center;display: flex">模块列表</span>
            </el-row>
            <el-row style="background-color: white">
              <span><el-button @click="showAddModular()">添加模块</el-button></span>
            </el-row>
          </div>
        </div>
        <div class="right-bottom" style="height: 90%;overflow: auto; width: 100%">
          <div style="overflow: auto; overflow-style: marquee-line; margin: 120px 20px 20px; width: 100%">
            <div v-for="(m, index) in modular" :key="index" class="text item" style="padding-bottom: 10px">
              <el-collapse accordion style="width: 90%">
                <el-collapse-item :title="m.name" :name="index">
                  <el-button style="margin-bottom: 10px; float: right;" @click="addInterface(m)">添加接口</el-button>
                  <el-input placeholder="输入接口名称" style="width: 180px"></el-input>
                  <!--<el-select placeholder="选择版本号"></el-select>-->
                  <el-table :border="true" :data="m.interface">
                    <el-table-column label="接口名称" prop="name">
                      <template slot-scope="{row}">
                        <div v-if="!row.edit">{{row.name}}</div>
                        <el-input v-else v-model="row.name"></el-input>
                      </template>
                    </el-table-column>
                    <el-table-column label="路由" prop="path" >
                      <template slot-scope="{row}">
                        <div v-if="!row.edit">{{row.path}}</div>
                        <el-input v-else v-model="row.path"></el-input>
                      </template>
                    </el-table-column>
                    <el-table-column label="版本" prop="version" >
                      <template slot-scope="{row}">
                        <div v-if="!row.edit">{{row.version}}</div>
                        <el-input v-else v-model="row.version"></el-input>
                      </template>
                    </el-table-column>
                    <el-table-column label="请求方式" prop="method">
                      <template slot-scope="{row}">
                        <div v-if="!row.edit" style="">{{row.method}}</div>
                        <el-input v-else v-model="row.method"></el-input>
                      </template>
                    </el-table-column>
                    <el-table-column label="是否必填" prop="require">
                      <template slot-scope="{row}">
                        <div v-if="!row.edit">{{row.require}}</div>
                        <el-input v-else v-model="row.require"></el-input>
                      </template>
                    </el-table-column>
                    <el-table-column label="示例" prop="example">
                      <template slot-scope="{row}">
                        <div v-if="!row.edit">{{row.example}}</div>
                        <el-input v-else v-model="row.example"></el-input>
                      </template>
                    </el-table-column>
                    <el-table-column label="类型" prop="type">
                      <template slot-scope="{row}">
                        <div v-if="!row.edit">{{row.type}}</div>
                        <el-input v-else v-model="row.type"></el-input>
                      </template>
                    </el-table-column>
                    <el-table-column label="描述" prop="desc">
                      <template slot-scope="{row}">
                        <div v-if="!row.edit">{{row.desc}}</div>
                        <el-input v-else v-model="row.desc"></el-input>
                      </template>
                    </el-table-column>
                    <el-table-column label="操作" width="200">
                      <template slot-scope="{row}">
                        <el-button
                          size="mini"
                          @click="updateRow(row, m)">修改</el-button>
                        <el-button
                          size="mini"
                          type="danger"
                          @click="delTpis(row.$index, row.row, m)"
                        >删除</el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-collapse-item>
              </el-collapse>
            </div>
          </div>
        </div>
        <el-dialog
          title="添加模块"
          :visible.sync="showDialog"
          width="30%"
          center>
          <el-form :model="newModular" ref="newModular" status-icon :rules="ModularRule">
            <el-form-item label="模块名称" prop="name">
              <el-input type="text" v-model="newModular.name" auto-complete="off"></el-input>
            </el-form-item>
            <el-form-item label="请求前缀" prop="prefix">
              <el-input type="text" v-model="newModular.prefix" auto-complete="off"></el-input>
            </el-form-item>
          </el-form>
          <span slot="footer" class="dialog-footer">
            <el-button @click="showDialog = false">取 消</el-button>
            <el-button type="primary" @click="AddModular('newModular')">确 定</el-button>
          </span>
        </el-dialog>
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
  name: 'projectinfo',
  components: {MyHeader, Left},
  data () {
    var validateName = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('项目名称不能为空'))
      }
      setTimeout(() => {
        callback()
      }, 10)
    }
    var validatePrefix = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('前缀不能为空'))
      }
      setTimeout(() => {
        callback()
      }, 10)
    }
    return {
      modular: [],
      showDialog: false,
      newModular: {
        name: '',
        prefix: '',
        interface: []
      },
      ModularRule: {
        name: [{
          validator: validateName, trigger: 'blur'
        }],
        prefix: [{
          validator: validatePrefix, trigger: 'blur'
        }]
      }
    }
  },
  mounted () {
    this.getModular()
  },
  methods: {
    back () {
      this.$router.push('/project')
    },
    async updateRow (row, m) {
      console.log(row, m)
      if (row.edit === true) {
        m.interface = ''
        let data = row
        data.modular = m
        console.log(data)
        let result = await request({
          url: apis.add_interface,
          method: 'post',
          data: data
        })
        console.log('result', result)
      }
      row.edit = !row.edit
    },
    addInterface (m) {
      let newRow = {
        name: '',
        path: '',
        method: '',
        require: '',
        example: '',
        type: '',
        desc: '',
        edit: true
      }
      m.interface.unshift(newRow)
    },
    showAddModular () {
      this.showDialog = true
    },
    AddModular (formName) {
      this.$refs[formName].validate(async (valid) => {
        if (valid) {
          let newM = {
            name: this.newModular.name,
            prefix: this.newModular.prefix,
            project: this.$store.getters.getPid
          }
          console.log('newM', newM)
          let result = await request({
            url: apis.add_modular,
            method: 'post',
            data: newM
          })
          console.log('result', result)
          if (result.data.code === 200) {
            this.modular.unshift(newM)
            this.showDialog = false
            this.newModular.name = ''
            this.newModular.prefix = ''
            this.newModular.interface = []
            this.$message.success(result.data.msg)
          } else if (result.data.code === 600) {
            this.$message.error(result.data.msg)
            this.$router.push('/')
          } else {
            this.$message.error(result.data.msg)
            this.showDialog = false
            this.newModular.name = ''
            this.newModular.prefix = ''
            this.newModular.interface = []
          }
        }
      })
    },
    delTpis (index, row, m) {
      console.log(index, row)
      this.$confirm('将永久删除该项目， 是否继续', '提示', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning',
        center: true
      }).then(() => {
        m.interface.splice(index, 1)
        console.log('tips')
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
    },
    async getModular () {
      let params = {
        pid: localStorage.getItem('pid')
      }
      console.log(params, '数据')
      let result = await request({
        url: apis.get_all_modular,
        method: 'get',
        params: params
      })
      if (result.data.code === 200) {
        this.modular = result.data.data
        for (let i in result.data.data) {
          this.modular[i].interface = []
        }
      } else if (result.data.code === 600) {
        this.$router.push('/')
      }
    },
    getInter () {
      console.log('----')
    }
  }
}
</script>

<style scoped>
  @import "../../assets/css/index.css";
</style>

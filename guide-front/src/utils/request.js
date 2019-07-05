import axios from 'axios'
import { Toast } from 'vant'
const service = axios.create({
  baseURL: 'http://localhost:9090'
})

service.defaults.headers.common['Authorization'] = localStorage.getItem('token')

service.interceptors.request.use(function (config) {
  let url = config.url
  if (url.indexOf('login') > -1) {
    localStorage.setItem('token', '')
    config.headers.Authorization = ''
  } else {
    config.headers.Authorization = localStorage.getItem('token')
  }
  return Promise.resolve(config)
},
function (error) {
  return Promise.reject(error)
}
)

service.interceptors.response.use(function (res) {
  if (res.data.token) {
    localStorage.setItem('token', 'basic ' + res.data.token)
  }
  if (res.status !== 200 || res.data.code !== 200) {
    Toast(res.data.msg)
  } else if (res.data.code === 600) {
    localStorage.setItem('token', '')
  }
  return res
},
function (error) {
  if (error.response.status === 600) {
    localStorage.setItem('token', '')
    Toast(error.response.data.msg)
    this.$router.push('/')
  } else if (error.request.status === 502) {
    Toast('服务器繁忙')
  } else {
    Toast(error.response.data.msg)
  }
  return Promise.reject(error)
})

export default service

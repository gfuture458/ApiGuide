const apiHead = '/v1'
const Apis = {
  login: '/user/login', // 登陆
  register: '/user/register', // 注册
  user_msg: '/user/get_user_detail', // 用户详情
  project: '/project/get_all_project', // 获取所有项目
  add_pro: '/project/add_project', // 添加项目
  project_info: '/project/info', // 项目详情
  del_project: '/project/delete', // 删除项目
  get_all_modular: '/modular/get_all_modular', // 获取项目模块
  add_modular: '/modular/add_modular', // 添加模块
  add_interface: '/interface/add_interface' // 添加模块
}

for (let key in Apis) {
  Apis[key] = apiHead + Apis[key]
}
export default Apis

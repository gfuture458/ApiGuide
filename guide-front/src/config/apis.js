const apiHead = '/v1'
const Apis = {
  login: '/user/login', // 登陆
  register: '/user/register', // 注册
  user_msg: '/user/get_user_detail', // 用户详情
  project: '/project/get_all_project', // 获取所有项目
  add_pro: '/project/add_project', // 添加项目
  project_info: '/project/info/' // 项目详情
}

for (let key in Apis) {
  Apis[key] = apiHead + Apis[key]
}
export default Apis

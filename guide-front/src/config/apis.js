const apiHead = '/v1'
const Apis = {
  login: '/user/login',
  register: '/user/register',
  user_msg: '/user/get_user_detail'
}

for (let key in Apis) {
  Apis[key] = apiHead + Apis[key]
}
export default Apis

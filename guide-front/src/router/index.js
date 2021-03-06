import Vue from 'vue'
import Router from 'vue-router'
import login from '@/views/login'
import register from '@/views/register'
import index from '@/views/index/index'
import project from '../views/project/project'
import addProject from '../views/project/project_add'
import projectInfo from '../views/project/projectinfo'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'login',
      component: login
    },
    {
      path: '/register',
      name: 'register',
      component: register
    },
    {
      path: '/index',
      name: 'index',
      component: index
    },
    {
      path: '/project',
      name: 'project',
      component: project
    },
    {
      path: '/addProject',
      name: 'addProject',
      component: addProject
    },
    {
      path: '/project_info',
      name: 'projectInfo',
      component: projectInfo
    }
  ]
})

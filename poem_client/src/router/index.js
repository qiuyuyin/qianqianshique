import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    redirect: 'dashboard',
  },
  {
    path: '/dashboard',
    name: 'dashboard',
    component: () => import('@/views/dashboard/Dashboard.vue'),
  },

  {
    path: '/author',
    name: 'author',
    component: () => import('@/views/author/AuthorList.vue'),
  },
  {
    path: '/author/:id',
    name: 'author-detail',
    component: () => import('@/views/author/AuthorDetail.vue'),
  },
  {
    path: '/poem',
    name: 'poem',
    component: () => import('@/views/poem/PoemList.vue'),
  },
  {
    path: '/poem/:poemType/:id',
    name: 'poem-detail',
    component: () => import('@/views/poem/PoemDetail.vue'),
  },
  {
    path: '/sentence',
    name: 'sentence',
    component: () => import('@/views/sentence/SentenceList.vue'),
  },
  {
    path: '/ai-poem/shi',
    name: 'ai-poem-shi',
    component: () => import('@/views/ai-poem/Shi.vue'),
  },
  {
    path: '/pages/user',
    name: 'user',
    component: () => import('@/views/pages/User.vue'),
  },
  {
    path: '/about',
    name: 'about',
    component: () => import('@/views/pages/About.vue'),
  },
  {
    path: '/pages/login',
    name: 'pages-login',
    component: () => import('@/views/pages/Login.vue'),
    meta: {
      layout: 'blank',
    },
  },
  {
    path: '/pages/register',
    name: 'pages-register',
    component: () => import('@/views/pages/Register.vue'),
    meta: {
      layout: 'blank',
    },
  },
  {
    path: '/error-404',
    name: 'error-404',
    component: () => import('@/views/Error.vue'),
    meta: {
      layout: 'blank',
    },
  },
  {
    path: '*',
    redirect: 'error-404',
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
})

export default router

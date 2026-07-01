import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import RegisterView from '../views/RegisterView.vue'
import ShareView from '../views/ShareView.vue'
import GameView from '../views/GameView.vue'
const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/register',
    name: 'register',
    component: RegisterView,
    beforeEnter: (to, from, next) => {
      const email = localStorage.getItem('user_email')
      if (!email) {
        next('/')
      } else {
        next()
      }
    }
  },
  {
    path: '/share-card/:id/:nama_lengkap?',
    name: 'share',
    component: ShareView
  },
  {
    path: '/card/:id/:nama_lengkap?',
    name: 'card',
    component: ShareView
  },
  {
    path: '/game',
    name: 'game',
    component: GameView
  }
]
const router = createRouter({
  history: createWebHistory(),
  routes
})
export default router

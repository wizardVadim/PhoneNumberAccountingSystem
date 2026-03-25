import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import LoginView from '@/views/LoginView.vue'
import UserList from '@/views/UserList.vue'
import PersonList from '@/views/PersonList.vue'
import PhoneList from '@/views/PhoneList.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView,
    },
    {
      path: '/users',
      name: 'users',
      component: UserList,
    },
    {
      path: '/persons',
      name: 'persons',
      component: PersonList,
    },
    {
      path: '/phones',
      name: 'phones',
      component: PhoneList,
    },
  ],
})

export default router

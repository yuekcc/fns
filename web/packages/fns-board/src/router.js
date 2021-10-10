import { createRouter, createWebHistory } from 'vue-router'

import CommonLayout from './components/common-layout.vue'
import { Overview } from './views/overview'
import { Services } from './views/services'
import { Functions } from './views/functions'

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      component: CommonLayout,
      children: [
        {
          path: '',
          redirect: { name: 'overview' }
        },
        {
          path: 'overview',
          name: 'overview',
          component: Overview
        },
        {
          path: 'services',
          name: 'service_list',
          component: Services
        },
        {
          path: 'services/new',
          name: 'create_service',
          component: Services
        },
        {
          path: 'functions',
          name: 'function_list',
          component: Functions
        },
        {
          path: 'functions/new',
          name: 'create_function',
          component: Functions
        },
        {
          path: 'setup',
          name: 'system_setup',
          component: Overview
        },
      ]
    },
    {
      path: '/:pathMatch(.*)*',
      redirect: "/"
    }
  ]
})
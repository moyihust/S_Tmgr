import { createRouter, createWebHistory } from 'vue-router';
import student from './pages/StudentMgr.vue';
import course from './pages/CourseMgr.vue';
import grade from './pages/GradeMgr.vue';
import sum from './pages/GradeSum.vue';
const routes = [
    { path: '/', component: student },
    { path: '/course', component: course },
    { path: '/grade', component: grade },
    { path: '/sum', component: sum },
  ];
  
  const router = createRouter({
    history: createWebHistory(),
    routes,
  });
  
  export default router;
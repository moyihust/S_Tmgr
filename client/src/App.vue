<template>
  <a-layout style="min-height: 100vh">
    <a-layout-sider v-model:collapsed="collapsed" collapsible>
      <div class="logo" >
        <img src="./assets/logo2.jpg" alt="logo" />
      </div>
      <a-menu v-model:selectedKeys="selectedKeys" theme="dark" mode="inline">
        <a-menu-item key="1">
          <router-link to="/">
            <IdcardOutlined />
            <span>学生管理</span>
          </router-link>
        </a-menu-item>
        <a-menu-item key="2">
          <router-link to="/course">
            <desktop-outlined />
            <span>课程管理</span>
          </router-link>
        </a-menu-item>
        <a-sub-menu key="sub2">
          <template #title>
            <span>
              <team-outlined />
              <span>成绩管理</span>
            </span>
          </template>
          <a-menu-item key="6" @click="$router.push('/grade')">成绩录入</a-menu-item>
          <a-menu-item key="8" @click="$router.push('/sum')">成绩查询</a-menu-item>
        </a-sub-menu>
      </a-menu>
    </a-layout-sider>
    <a-layout>
      <a-layout-header style="background: #fff; padding: 0">
        <h1>学生课程成绩管理系统</h1>
      </a-layout-header>
      <a-layout-content style="margin: 0 16px">
        <router-view></router-view>
      </a-layout-content>
      <a-layout-footer style="text-align: center">
        moyi ©2021 Created by moyi
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>
<script setup>
import { ref } from 'vue';
import{ IdcardOutlined, DesktopOutlined, TeamOutlined } from '@ant-design/icons-vue';
const collapsed = ref(false);
const selectedKeys = ref(['1']);
</script>
<style scoped>
#components-layout-demo-side .logo {
  height: 32px;
  margin: 16px;
  background: rgba(255, 255, 255, 0.3);
}

.site-layout .site-layout-background {
  background: #fff;
}

[data-theme='dark'] .site-layout .site-layout-background {
  background: #141414;
}
</style>
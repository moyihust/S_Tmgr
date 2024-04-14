<template>
    <!-- 搜索学生信息的输入框和按钮 -->
    <a-input-search
      placeholder="输入学生学号进行搜索"
      @search="handleSearchStudent"
    />
    <!-- 学生信息展示的抽屉 -->
    <a-drawer
      :visible="showStudentInfoDrawer"
      :title="'学生信息'"
      :closable="true"
      placement="right"
      :width="520"
      @close="handleCloseDrawer"
    >
      <p>学号：{{ studentInfo.studentId }}</p>
      <p>姓名：{{ studentInfo.name }}</p>
      <p>性别：{{ studentInfo.gender }}</p>
      <p>年龄：{{ studentInfo.age }}</p>
      <p>专业：{{ studentInfo.major }}</p>
    </a-drawer>
  </template>
  
  <script>
  export default {
    data() {
      return {
        showStudentInfoDrawer: false,
        studentInfo: {}
      };
    },
    methods: {
      handleSearchStudent(value) {
        if (value) {
          // 在这里模拟查询学生信息
          this.studentInfo = {
            studentId: value,
            name: '张三',
            gender: '男',
            age: 20,
            major: '计算机科学与技术'
          };
          // 显示学生信息抽屉
          this.showStudentInfoDrawer = true;
        } else {
          this.studentInfo = {};
          this.showStudentInfoDrawer = false;
        }
      },
      handleCloseDrawer() {
        // 关闭抽屉时可以执行的操作，比如重置学生信息
        this.studentInfo = {};
        this.showStudentInfoDrawer = false;
      }
    }
  };
  </script>
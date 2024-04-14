<template>
    <div>
      <!-- 导入学生的按钮 -->
      <a-button type="primary" @click="showImportDrawer = true">导入学生</a-button>
  
      <!-- 导入学生信息的抽屉 -->
      <a-drawer
        :visible="showImportDrawer"
        :title="'导入学生信息'"
        :closable="true"
        :width="520"
        @close="handleCloseDrawer"
      >
        <a-form
          :form="studentForm"
          @submit="handleImportStudent"
          layout="vertical"
        >
          <a-form-item label="学号" :rules="[{ required: true, message: '请输入学号' }]">
            <a-input v-model="studentForm.Sno" placeholder="请输入学号" />
          </a-form-item>
          <a-form-item label="姓名" :rules="[{ required: true, message: '请输入姓名' }]">
            <a-input v-model="studentForm.Sname" placeholder="请输入姓名" />
          </a-form-item>
          <a-form-item label="性别">
            <a-radio-group v-model="studentForm.Ssex">
              <a-radio value="男">男</a-radio>
              <a-radio value="女">女</a-radio>
            </a-radio-group>
          </a-form-item>
          <a-form-item label="年龄" :rules="[{ required: true, message: '请输入年龄' }]">
            <a-input-number v-model="studentForm.Sage" placeholder="请输入年龄" />
          </a-form-item>
          <a-form-item label="专业">
            <a-input v-model="studentForm.Sdept" placeholder="请输入专业" />
          </a-form-item>
          <a-form-item>
            <a-button type="primary" html-type="submit">提交</a-button>
          </a-form-item>
        </a-form>
      </a-drawer>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        showImportDrawer: false,
        studentForm: {
          Sno: '',
          Sname: '',
          Ssex: '男',
          Sage: null,
          Sdept: ''
        }
      };
    },
    methods: {
      handleImportStudent() {
        // 在这里处理导入学生逻辑
        console.log('导入的学生信息:', this.studentForm);
        // 关闭抽屉并重置表单
        this.showImportDrawer = false;
        this.studentForm = {
          Sno: '',
          Sname: '',
          Ssex: '男',
          Sage: null,
          Sdept: ''
        };
      },
      handleCloseDrawer() {
        // 可以在这里执行抽屉关闭时的清理工作，如果需要的话
        this.showImportDrawer = false;
      }
    }
  };
  </script>
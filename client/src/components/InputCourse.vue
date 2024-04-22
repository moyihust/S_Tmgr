<template>
    <a-button type="primary" @click="showDrawer">
      <template #icon><PlusOutlined /></template>
      New Course
    </a-button>
    <a-drawer
      title="Create a new course"
      :width="720"
      :open="open"
      :body-style="{ paddingBottom: '80px' }"
      :footer-style="{ textAlign: 'right' }"
      @close="onClose"
    >
      <a-form :model="form" :rules="rules" layout="vertical">
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item label="Course No" name="Cno">
              <a-input v-model:value="form.Cno" placeholder="Please enter course number" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="Course Name" name="Cname">
              <a-input v-model:value="form.Cname" placeholder="Please enter course name" />
            </a-form-item>
          </a-col>
        </a-row>
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item label="Prerequisite Course No" name="Cpno">
              <a-input v-model:value="form.Cpno" placeholder="Please enter prerequisite course number" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="Credit" name="Ccredit">
              <a-input-number v-model:value="form.Ccredit" placeholder="Please enter credit" />
            </a-form-item>
          </a-col>
        </a-row>
      </a-form>
      <template #extra>
        <a-space>
          <a-button @click="onClose">Cancel</a-button>
          <a-button type="primary" @click="onSubmit">Submit</a-button>
        </a-space>
      </template>
    </a-drawer>
  </template>
  
  <script setup>
  import { reactive, ref } from 'vue';
  import axios from 'axios';
  const axiosInstance = axios.create({
    baseURL: 'http://10.12.180.78:1145'
  });
  const form = reactive({
    Cno: '',
    Cname: '',
    Cpno: null,
    Ccredit: null,
  });
  const rules = {
    Cno: [
      {
        required: true,
        message: 'Please enter course number',
      },
    ],
    Cname: [
      {
        required: true,
        message: 'Please enter course name',
      },
    ],
    Cpno: [
      {
        required: false,
        message: 'Please enter prerequisite course number',
      },
    ],
    Ccredit: [
      {
        required: true,
        message: 'Please enter credit',
      },
    ],
  };
  const open = ref(false);
  const showDrawer = () => {
    open.value = true;
  };
  const onClose = () => {
    open.value = false;
  };
  const onSubmit = () => {
    console.log('submit!');
    console.log(form);
    if(form.Cpno === '') {
      form.Cpno = null;
    }
    axiosInstance.post('/inputCourse', form)
      .then(response => {
        console.log('Response:', response.data);
        // 关闭抽屉并重置表单
        onClose();
        form.Cno = '';
        form.Cname = '';
        form.Cpno = null;
        form.Ccredit = null;
      })
      .catch(error => {
        console.error('Error submitting form:', error);
      });
  };
  </script>
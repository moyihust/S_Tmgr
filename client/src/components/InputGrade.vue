<template>
    <a-button type="primary" @click="showDrawer">
      <template #icon><PlusOutlined /></template>
      Enter Grade
    </a-button>
    <a-drawer
      title="Enter a new grade"
      :width="720"
      :open="open"
      :body-style="{ paddingBottom: '80px' }"
      :footer-style="{ textAlign: 'right' }"
      @close="onClose"
    >
      <a-form :model="form" :rules="rules" layout="vertical">
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item label="Student No" name="Sno">
              <a-input v-model:value="form.Sno" placeholder="Please enter student number" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="Course No" name="Cno">
              <a-input v-model:value="form.Cno" placeholder="Please enter course number" />
            </a-form-item>
          </a-col>
        </a-row>
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item label="Grade" name="Grade">
              <a-input-number v-model:value="form.Grade" placeholder="Please enter grade" />
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
    baseURL: 'http://localhost:1145'
  });
  const form = reactive({
    Sno: '',
    Cno: '',
    Grade: null,
  });
  const rules = {
    Sno: [
      {
        required: true,
        message: 'Please enter student number',
      },
    ],
    Cno: [
      {
        required: true,
        message: 'Please enter course number',
      },
    ],
    Grade: [
      {
        required: true,
        message: 'Please enter grade',
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
    axiosInstance.post('/inputSC', form)
      .then(response => {
        console.log('Response:', response.data);
        // 关闭抽屉并重置表单
        onClose();
        form.Sno = '';
        form.Cno = '';
        form.Grade = null;
      })
      .catch(error => {
        console.error('Error submitting form:', error);
      });
  };
  </script>
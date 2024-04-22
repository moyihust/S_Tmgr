<template>
    <a-button type="danger" @click="showDrawer">
      <template #icon><PlusOutlined /></template>
      编辑
    </a-button>
    <a-drawer
      title="Edit grade"
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
  import { defineProps, defineEmits } from 'vue';
  const props = defineProps({
    record: {
      type: Object,
      required: true,
    },
  });
  
  const emit = defineEmits(['refresh']);
  import { reactive, ref } from 'vue';
  import axios from 'axios';
  
  const axiosInstance = axios.create({
    baseURL: 'http://10.12.180.78:1145'
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
    console.log(props.record.Sno);
    open.value = true;
    form.Sno = props.record.Sno;
    form.Cno = props.record.Cno;
    form.Grade = props.record.Grade;
  };
  const onClose = () => {
    open.value = false;
  };
  const onSubmit = async () => {
    console.log('submit!');
    console.log(form);
    try {
      const response = await axiosInstance.put(`/updateSC/${props.record.Sno}/${props.record.Cno}`, form);
      console.log(response.data);
      // 关闭抽屉并清空表单
      open.value = false;
      // 刷新成绩信息
      emit('refresh');
      Object.keys(form).forEach(key => {
        form[key] = '';
      });
    } catch (error) {
      console.error(error);
    }
  };
  </script>
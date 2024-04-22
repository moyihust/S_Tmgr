<template>
    <a-button type="danger" @click="showDrawer">
      <template #icon><PlusOutlined /></template>
      编辑
    </a-button>
    <a-drawer
      title="Edit student account"
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
            <a-form-item label="Name" name="Sname">
              <a-input v-model:value="form.Sname" placeholder="Please enter student name" />
            </a-form-item>
          </a-col>
        </a-row>
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item label="Sex" name="Ssex">
              <a-select v-model:value="form.Ssex" placeholder="Please select sex">
                <a-select-option value="男">Male</a-select-option>
                <a-select-option value="女">Female</a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="Age" name="Sage">
              <a-input-number v-model:value="form.Sage" placeholder="Please enter age" />
            </a-form-item>
          </a-col>
        </a-row>
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item label="Major" name="Sdept">
              <a-input v-model:value="form.Sdept" placeholder="Please enter major" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="Scholarship" name="scholarship">
              <a-select v-model:value="form.scholarship" placeholder="Please select scholarship status">
                <a-select-option value="是">Yes</a-select-option>
                <a-select-option value="否">No</a-select-option>
              </a-select>
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
    Sname: '',
    Ssex: '',
    Sage: null,
    Sdept: '',
    scholarship: '',
  });
  const rules = {
    Sno: [
      {
        required: true,
        message: 'Please enter student number',
      },
    ],
    Sname: [
      {
        required: true,
        message: 'Please enter student name',
      },
    ],
    Ssex: [
      {
        required: true,
        message: 'Please select sex',
      },
    ],
    Sage: [
      {
        required: true,
        message: 'Please enter age',
      },
    ],
    Sdept: [
      {
        required: true,
        message: 'Please enter major',
      },
    ],
    scholarship: [
      {
        required: true,
        message: 'Please select scholarship status',
      },
    ],
  };
  const open = ref(false);
  const showDrawer = () => {
    console.log(props.studentId);
    open.value = true;
    form.Sno = props.record.Sno;
    form.Sname = props.record.Sname;
    form.Ssex = props.record.Ssex;
    form.Sage = props.record.Sage;
    form.Sdept = props.record.Sdept;
    form.scholarship = props.record.scholarship;
  };
  const onClose = () => {
    open.value = false;
  };
const onSubmit = async () => {
  console.log('submit!');
  console.log(form);
  try {
    const response = await axiosInstance.put(`/updateStudent/${props.record.Sno}`, form);
    console.log(response.data);
    // 关闭抽屉并清空表单
    open.value = false;
    // 刷新学生信息
    emit('refresh');
    Object.keys(form).forEach(key => {
      form[key] = '';
    });
  } catch (error) {
    console.error(error);
  }
};
  
  </script>


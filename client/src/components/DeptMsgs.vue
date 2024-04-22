<template>
    <a-table :columns="columns" :dataSource="grades" rowKey="Sno">

       
    </a-table>
</template>
<script>
import { defineProps } from 'vue';
const props = defineProps({
    selectedSdept: {
      type: String,
      default: ''
    }
});
import axios from 'axios';
const axiosInstance = axios.create({
  baseURL: 'http://10.12.180.78:1145'
});
export default {
    data() {
        return {
            showEditDrawer: false,
            currentSno: null,
            grades: [
                // 这里填充成绩数据
                { Sno: '001', Sname: '学生一', Cno: '002', Cname: '课程二', Grade: 85 },
                { Sno: '002', Sname: '学生二', Cno: '001', Cname: '课程一', Grade: 90 },
                // ...更多成绩数据
            ],
            columns: [
                {
                    title: '学号',
                    dataIndex: 'Sno',
                    key: 'Sno'
                },
                {
                    title: '学生名字',
                    dataIndex: 'Sname',
                    key: 'Sname'
                },
                {
                    title: '课程号',
                    dataIndex: 'Cno',
                    key: 'Cno'
                },
                {
                    title: '课程名',
                    dataIndex: 'Cname',
                    key: 'Cname'
                },
                {
                    title: '成绩',
                    dataIndex: 'Grade',
                    key: 'Grade'
                },
                {
                    title: 'Action',
                    key: 'action',
                },
            ],
        };
    },
    methods: {
        async fetchStats(selectedSdept) {
            try {
                console.log(this.selectedSdept);
                const response = await axiosInstance.get(`/studentStats/${selectedSdept}`);
                this.stats = response.data[0];
            } catch (error) {
                console.error('Error fetching student stats:', error);
                // 这里可以添加更多的错误处理逻辑
            }
        }
    },
    mounted() {
        axiosInstance.get('/sc/'+props.selectedSdept)
      .then(response => {
        this.grades = response.data;
      })
      .catch(error => {
        console.error('Error fetching grades:', error);
      });
    },
};
</script>
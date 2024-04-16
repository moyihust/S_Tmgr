<template>
    <div>
        <div class="button-container">
            <a-button type="danger" @click="refresh">
                <RedoOutlined /> 
            </a-button>
        </div>
    <a-table :columns="columns" :dataSource="grades" rowKey="Sno">
        <!-- 添加操作列 -->
        <template #headerCell="{ column }">
            <template v-if="column.key === 'Sname'">
                <span>
                    <smile-outlined />
                    Sname
                </span>
            </template>
        </template>

        <template #bodyCell="{ column, record }">
            <template v-if="column.key === 'Sname'">
                <a>
                    {{ record.Sname }}
                </a>
            </template>
            <template v-else-if="column.key === 'action'">
                <span>
                    <EditGrade :record="record" @refresh="refresh"/>
                    <a-divider type="vertical" />
                    <a-popconfirm title="确定要删除吗?" @confirm="handleDeleteGrade(record.Sno)">
                        <a-button  type="danger">删除</a-button>
                    </a-popconfirm>
                </span>
            </template>
        </template>
    </a-table>
</div>
</template>
<script>
import axios from 'axios';
import EditGrade from './EditGrade.vue';
import { RedoOutlined, SmileOutlined } from '@ant-design/icons-vue';
const axiosInstance = axios.create({
  baseURL: 'http://localhost:1145'
});
export default {
    components: {
        EditGrade,
        RedoOutlined,
        SmileOutlined
    },
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
        openEditDrawer(record) {
            this.showEditDrawer = true;
            this.currentSno = record.Sno;
            this.editGradeForm = {
                Sno: record.Sno,
                Sname: record.Sname,
                Cno: record.Cno,
                Cname: record.Cname,
                Grade: record.Grade
            };
        },
        handleEditGrade() {
            console.log('编辑后的成绩信息:', this.editGradeForm);
            this.showEditDrawer = false;
            this.currentSno = null;
            const index = this.grades.findIndex(g => g.Sno === this.editGradeForm.Sno);
            if (index !== -1) {
                this.grades.splice(index, 1, this.editGradeForm);
            }
        },
        handleDeleteGrade(sno) {
            console.log('删除成绩:', sno);
            axiosInstance.delete(`/deleteSC/${sno}`)
                .then(response => {
                    console.log('Deleted grade:', response.data);
                    this.refresh();
                })
                .catch(error => {
                    console.error('Error deleting grade:', error);
                });
        },
        closeEditDrawer() {
            this.showEditDrawer = false;
            this.currentSno = null;
        },
        submitEditGrade() {
            this.$refs.editGradeForm.validate((valid) => {
                if (valid) {
                    this.handleEditGrade();
                } else {
                    console.log('error submit!!');
                    return false;
                }
            });
            this.closeEditDrawer();
        },
        refresh() {
            console.log('refresh');
            axiosInstance.get('/sc')
            .then(response => {
                this.grades = response.data;
            })
            .catch(error => {
                console.error('Error fetching grades:', error);
            });
        },
    },
    mounted() {
        axiosInstance.get('/sc')
      .then(response => {
        this.grades = response.data;
      })
      .catch(error => {
        console.error('Error fetching grades:', error);
      });
    },
};
</script>
<style scoped>
.button-container {
    text-align: right;
}
</style>
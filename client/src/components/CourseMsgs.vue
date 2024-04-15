<template>
    <a-table :columns="columns" :dataSource="courses" rowKey="Cno">
        <!-- 添加操作列 -->
        <template #headerCell="{ column }">
            <template v-if="column.key === 'Cname'">
                <span>
                    <smile-outlined />
                    Cname
                </span>
            </template>
        </template>

        <template #bodyCell="{ column, record }">
            <template v-if="column.key === 'Cname'">
                <a>
                    {{ record.Cname }}
                </a>
            </template>
            <template v-else-if="column.key === 'action'">
                <span>
                    <EditCourse :record="record" @refresh="refresh"/>
                    <a-divider type="vertical" />
                    <a-popconfirm title="确定要删除吗?" @confirm="handleDeleteCourse(record.Cno)">
                        <a-button  type="danger">删除</a-button>
                    </a-popconfirm>
                </span>
            </template>
        </template>
    </a-table>
</template>
<script>
import axios from 'axios';
import EditCourse from './EditCourse.vue';
const axiosInstance = axios.create({
  baseURL: 'http://localhost:1145'
});
export default {
    components: {
        EditCourse
    },
    data() {
        return {
            showEditDrawer: false,
            currentCno: null,
            courses: [
                // 这里填充课程数据
                { Cno: '001', Cname: '课程一', Cpno: '002', Ccredit: 3 },
                { Cno: '002', Cname: '课程二', Cpno: '001', Ccredit: 2 },
                // ...更多课程数据
            ],
            columns: [
                {
                    title: '课程编号',
                    dataIndex: 'Cno',
                    key: 'Cno'
                },
                {
                    title: '课程名称',
                    dataIndex: 'Cname',
                    key: 'Cname'
                },
                {
                    title: '先修课程编号',
                    dataIndex: 'Cpno',
                    key: 'Cpno'
                },
                {
                    title: '学分',
                    dataIndex: 'Ccredit',
                    key: 'Ccredit'
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
            this.currentCno = record.Cno;
            this.editCourseForm = {
                Cno: record.Cno,
                Cname: record.Cname,
                Cpno: record.Cpno,
                Ccredit: record.Ccredit
            };
        },
        handleEditCourse() {
            console.log('编辑后的课程信息:', this.editCourseForm);
            this.showEditDrawer = false;
            this.currentCno = null;
            const index = this.courses.findIndex(c => c.Cno === this.editCourseForm.Cno);
            if (index !== -1) {
                this.courses.splice(index, 1, this.editCourseForm);
            }
        },
        handleDeleteCourse(cno) {
            console.log('删除课程:', cno);
            axiosInstance.delete(`/deleteCourse/${cno}`)
                .then(response => {
                    console.log('Deleted course:', response.data);
                    this.refresh();
                })
                .catch(error => {
                    console.error('Error deleting course:', error);
                });
        },
        closeEditDrawer() {
            this.showEditDrawer = false;
            this.currentCno = null;
        },
        submitEditCourse() {
            this.$refs.editCourseForm.validate((valid) => {
                if (valid) {
                    this.handleEditCourse();
                } else {
                    console.log('error submit!!');
                    return false;
                }
            });
            this.closeEditDrawer();
        },
        refresh() {
            console.log('refresh');
            axiosInstance.get('/course')
            .then(response => {
                this.courses = response.data;
            })
            .catch(error => {
                console.error('Error fetching courses:', error);
            });
        },
    },
    mounted() {
        axiosInstance.get('/course')
      .then(response => {
        this.courses = response.data;
      })
      .catch(error => {
        console.error('Error fetching courses:', error);
      });
    },
};
</script>
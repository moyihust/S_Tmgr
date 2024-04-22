<template>
    <div>
        <div class="button-container">
            <a-button type="danger" @click="refresh">
                <RedoOutlined />
            </a-button>
        </div>
        <a-table :columns="columns" :dataSource="students" :pagination="students.length > pageSize" rowKey="studentId" >
            <!-- 添加操作列 -->
            <template #headerCell="{ column }">
                <template v-if="column.key === 'Sname'">
                    <span>
                        <smile-outlined />
                        Name
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
                        <EditStudent :record="record" @refresh="refresh" />
                        <a-divider type="vertical" />
                        <a-popconfirm title="确定要删除吗?" @confirm="handleDeleteStudent(record.Sno)">
                            <a-button type="danger">删除</a-button>
                        </a-popconfirm>
                    </span>
                </template>
            </template>
        </a-table>
    </div>
</template>
<script>
import axios from 'axios';
import EditStudent from './EditStudent.vue';
import { RedoOutlined, SmileOutlined } from '@ant-design/icons-vue';
const axiosInstance = axios.create({
    baseURL: 'http://10.12.180.78:1145'
});
export default {
    components: {
        EditStudent,
        RedoOutlined,
        SmileOutlined
    },
    data() {
        return {
            // ...其他数据
            pageSize: 10, // 每页显示的学生数量
            showEditDrawer: false, // 确保这个变量被定义
            currentSno: null, // 用于记录当前编辑的学生学号
            students: [
                // 这里填充学生数据，包括新添加的奖学金字段
                { Sno: '001', Sname: '张三', Ssex: '男', Sage: 20, Sdept: 'MA', scholarship: '是' },
                { Sno: '002', Sname: '李四', Ssex: '女', Sage: 19, Sdept: 'IS', scholarship: '否' },
                // ...更多学生数据
            ],
            columns: [
                // ...其他列定义
                {
                    title: '学号',
                    dataIndex: 'Sno',
                    key: 'Sno'
                },
                {
                    title: '姓名',
                    dataIndex: 'Sname',
                    key: 'Sname'
                },
                {
                    title: '性别',
                    dataIndex: 'Ssex',
                    key: 'Ssex'
                },
                {
                    title: '年龄',
                    dataIndex: 'Sage',
                    key: 'Sage'
                },
                {
                    title: '专业',
                    dataIndex: 'Sdept',
                    key: 'Sdept'
                },
                {
                    title: '奖学金',
                    dataIndex: 'Scholarship',
                    key: 'Scholarship'
                },
                {
                    title: 'Action',
                    key: 'action',
                },
            ],
        };
    },
    methods: {
        // ...其他方法
        openEditDrawer(record) {
            // 显示编辑抽屉，并填充当前学生信息
            this.showEditDrawer = true;
            this.currentSno = record.Sno;
            this.editStudentForm = {
                Sno: record.Sno,
                Sname: record.Sname,
                Ssex: record.Ssex,
                Sage: record.Sage,
                Sdept: record.Sdept,
                scholarship: record.scholarship // 添加奖学金字段
            };
        },
        handleEditStudent() {
            // 处理编辑学生信息的逻辑
            console.log('编辑后的学生信息:', this.editStudentForm);
            // 关闭抽屉
            this.showEditDrawer = false;
            this.currentSno = null;
            // 更新学生信息（示例代码，需要根据实际情况调整）
            const index = this.students.findIndex(s => s.Sno === this.editStudentForm.Sno);
            if (index !== -1) {
                this.students.splice(index, 1, this.editStudentForm);
            }
        },
        handleDeleteStudent(sno) {
            console.log('删除学生:', sno);
            axiosInstance.delete(`/deleteStudent/${sno}`)
                .then(response => {
                    console.log('Deleted student:', response.data);
                    this.refresh(); // 删除学生后刷新列表
                })
                .catch(error => {
                    console.error('Error deleting student:', error);
                });
        },
        // ...其他方法
        closeEditDrawer() {
            // 关闭编辑抽屉时的操作
            this.showEditDrawer = false;
            this.currentSno = null;
        },
        submitEditStudent() {
            // 提交编辑学生信息的表单
            this.$refs.editStudentForm.validate((valid) => {
                if (valid) {
                    this.handleEditStudent();
                } else {
                    console.log('error submit!!');
                    return false;
                }
            });
            this.closeEditDrawer();
        },
        refresh() {
            console.log('refresh');
            axiosInstance.get('/student')
                .then(response => {
                    this.students = response.data;
                })
                .catch(error => {
                    console.error('Error fetching students:', error);
                });
        },
    },
    mounted() {
        // ...其他初始化操作
        axiosInstance.get('/student')
            .then(response => {
                this.students = response.data;
            })
            .catch(error => {
                console.error('Error fetching students:', error);
            });
    },
};
</script>
<style scoped>
.button-container {
    margin-bottom: 16px;
    text-align: right;
}
</style>
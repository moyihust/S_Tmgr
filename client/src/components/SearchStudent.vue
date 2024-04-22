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
    <p>学号：{{ studentInfo.Sno }}</p>
    <p>姓名：{{ studentInfo.Sname }}</p>
    <p>性别：{{ studentInfo.Ssex }}</p>
    <p>年龄：{{ studentInfo.Sage }}</p>
    <p>专业：{{ studentInfo.Sdept }}</p>
    <p>奖学金：{{ studentInfo.Scholarship === '是' ? '有' : '无' }}</p>
    <h3>课程信息</h3>
    <a-table :columns="courseColumns" :dataSource="studentInfo.courses" rowKey="Cno">
      <a-table-column title="课程号" dataIndex="Cno" key="Cno" />
      <a-table-column title="课程名" dataIndex="Cname" key="Cname" />
      <a-table-column title="成绩" dataIndex="Grade" key="Grade" />
    </a-table>
  </a-drawer>
</template>

<script>
import axios from 'axios';
const axiosInstance = axios.create({
  baseURL: 'http://10.12.180.78:1145'
});
export default {
  data() {
    return {
      showStudentInfoDrawer: false,
      studentInfo: {
        Sno: '',
        Sname: '',
        Ssex: '',
        Sage: 0,
        Sdept: '',
        Scholarship: '',
        courses: [] // 用于存储课程信息的数组
      },
      courseColumns: [
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
        }
      ]
    };
  },
  methods: {
    handleSearchStudent(sno) {
            if (sno) {
                axiosInstance.get(`/studentCourses/${sno}`)
                    .then(response => {
                        this.studentInfo = {
                            Sno: response.data[0].Sno,
                            Sname: response.data[0].Sname,
                            Ssex: response.data[0].Ssex,
                            Sage: response.data[0].Sage,
                            Sdept: response.data[0].Sdept,
                            Scholarship: response.data[0].Scholarship,
                            courses: response.data.map(item => ({
                                Cno: item.Cno,
                                Cname: item.Cname,
                                Grade: item.Grade
                            }))
                        };
                        this.showStudentInfoDrawer = true;
                    })
                    .catch(error => {
                        console.error('Error getting student courses:', error);
                    });
            } else {
                this.studentInfo = {};
                this.showStudentInfoDrawer = false;
            }
        },
    handleCloseDrawer() {
      this.studentInfo = {};
      this.showStudentInfoDrawer = false;
    }
  }
};
</script>
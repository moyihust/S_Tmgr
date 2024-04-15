<template>
    <div>
        <a-select v-model="selectedSdept" placeholder="请选择系列" @change="handleGrade">
            <a-select-option v-for="sdept in sdepts" :key="sdept.Sdept" :value="sdept.Sdept">
                {{ sdept.Sdept }}
            </a-select-option>
        </a-select>
        <a-row v-if="stats">
            <a-col span="8">
                <a-statistic title="平均成绩" :value="stats.AverageGrade" precision="2" />
            </a-col>
            <a-col span="8">
                <a-statistic title="最好成绩" :value="stats.BestGrade" />
            </a-col>
            <a-col span="8">
                <a-statistic title="最差成绩" :value="stats.WorstGrade" />
            </a-col>
            <a-col span="8">
                <a-statistic title="优秀人数" :value="stats.Excellent" />
            </a-col>
            <a-col span="8">
                <a-statistic title="不及格人数" :value="stats.Fail" />
            </a-col>
        </a-row>
    </div>
    <a-table :columns="columns" :dataSource="grades" rowKey="Sno">
</a-table>
</template>

<script>
import axios from 'axios';
const axiosInstance = axios.create({
    baseURL: 'http://localhost:1145'
});
export default {
    
    data() {
        return {
            sdepts: [],
            selectedSdept: '',
            stats: null,
            showEditDrawer: false,
            currentSno: null,
            grades: [
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
            ],
        };
    },
    async created() {
        try {
            const response = await axiosInstance.get('/sdeptlist');
            this.sdepts = response.data;
        } catch (error) {
            console.error('Error fetching sdepts:', error);
            // 这里可以添加更多的错误处理逻辑
        }
    },
    methods: {
        async handleGrade(sdept){
            this.fetchStats(sdept);
            this.fetchGrades(sdept);
        },
        async fetchStats(sdept) {
            try {
                console.log(this.selectedSdept);
                const response = await axiosInstance.get(`/studentStats/${sdept}`);
                this.stats = response.data[0];
            } catch (error) {
                console.error('Error fetching student stats:', error);
                // 这里可以添加更多的错误处理逻辑
            }
        },
        async fetchGrades(sdept) {
            try {
                console.log(this.selectedSdept);
                const response = await axiosInstance.get(`/search/${sdept}`);
                this.grades = response.data;
            } catch (error) {
                console.error('Error fetching student grades:', error);
                // 这里可以添加更多的错误处理逻辑
            }
        }
    }
};
</script>
<template>
    <div class="container">
        <div class="header">
            <h1>课程管理</h1>
        </div>

        <div class="content">
            <div class="input-course">
                <InputCourse/>
            </div>
        </div>
        <div class="course-msgs">
                <CourseMsgs />
            </div>
    </div>
</template>

<style scoped>
.container {
    max-width: 1200px;
    /* 根据需要调整最大宽度 */
    margin: 0 auto;
    /* 水平居中 */
    padding: 20px;
    box-sizing: border-box;
}

.header {
    text-align: center;
    margin-bottom: 30px;
    /* 增加底部间距 */
}

.content {
    display: flex;
    /* 使用Flexbox布局 */
    justify-content: space-between;
    /* 子元素之间的间隔均匀分布 */
    flex-wrap: wrap;
    /* 允许子元素换行 */
}

.input-course,
.search-student,
.course-msgs {
    flex: 1;
    /* 每个组件占据相同的空间 */
    min-width: 300px;
    /* 根据需要设置最小宽度 */
    margin: 0 10px;
    /* 给组件添加间距 */
}

/* 响应式调整 */
@media (max-width: 768px) {
    .content {
        flex-direction: column;
    }

    .input-course,
    .search-student,
    .course-msgs {
        margin: 10px 0;
        /* 在较小屏幕上，只保留上下间距 */
    }
}
</style>
<script setup>
import CourseMsgs from '../components/CourseMsgs.vue';
import InputCourse from '../components/InputCourse.vue';


</script>
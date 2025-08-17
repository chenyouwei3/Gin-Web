<template>
  <NavigationBar />
  <div class="operation-container">
    <a-card title="操作日志" :bordered="false">
      <!-- 搜索表单 -->
      <a-form layout="inline" :model="searchForm" class="search-form">
        <a-form-item label="时间范围">
          <a-range-picker
            v-model:value="searchForm.timeRange"
            show-time
            format="YYYY-MM-DD HH:mm:ss"
            :placeholder="['开始时间', '结束时间']"
            @change="onTimeRangeChange"
          />
        </a-form-item>
        <a-form-item label="账号">
          <a-input
            v-model:value="searchForm.account"
            placeholder="请输入账号"
            allow-clear
          />
        </a-form-item>
        <a-form-item>
          <a-space>
            <a-button type="primary" @click="handleSearch">搜索</a-button>
            <a-button @click="resetSearch">重置</a-button>
          </a-space>
        </a-form-item>
      </a-form>
      <!-- 表格 -->
      <a-table
        :columns="columns"
        :data-source="tableData"
        :loading="loading"
        :pagination="pagination"
        @change="handleTableChange"
        :scroll="{ x: 1300 }"
        class="operation-table"
      >
        <template #bodyCell="{ column, text }">
          <template v-if="column.dataIndex === 'status'">
            <a-tag :color="text === 200 ? 'success' : 'error'">
              {{ text }}
            </a-tag>
          </template>
          <template v-else-if="column.dataIndex === 'startTime'">
            {{ formatDate(text) }}
          </template>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useTable } from '@/tools/page/common'
import NavigationBar from '@/components/NavigationBar.vue'
import {logByOperationColumns} from '@/tools/page/columns'
import {formatDate,newTimeRangeHandler}from '@/tools/page/time'
/*----------------------------------------全局变量----------------------------------------*/
const {
  tableData,
  loading,
  pagination,
  searchForm,
  fetchData,
  resetSearch,
  handleSearch,
  handleTableChange,
} = useTable('logs');
const onTimeRangeChange = newTimeRangeHandler(searchForm)
const columns = logByOperationColumns //表格定义
// 组件挂载时获取数据
onMounted(() => {
  fetchData()
})
</script>
<style scoped>
</style> 
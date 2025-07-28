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
            @change="handleTimeRangeChange"
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
import { ref, reactive, onMounted } from 'vue'
import { getOperationLog } from '@/tools/api'
import NavigationBar from '@/components/NavigationBar.vue'
/*----------------------------------------全局变量----------------------------------------*/
const loading = ref(false)
const searchForm = reactive({timeRange: [],startTime: '',endTime: '',account:''}) //赛选条件
const pagination = reactive({current: 1,pageSize: 10,total: 0,showSizeChanger: true,showQuickJumper: true,showTotal: (total) => `共 ${total} 条记录`})//分页数据
const tableData = ref([]) //渲染数据
const columns = [// 表格列定义
  {
    title: '用户账号',
    dataIndex: 'account',
    width: 80,
  },
  {
    title: 'IP地址',
    dataIndex: 'ip',
    width: 120
  },
  {
    title: '请求方法',
    dataIndex: 'method',
    width: 80
  },
  {
    title: '状态',
    dataIndex: 'status',
    width: 80
  },
  {
    title: '请求路径',
    dataIndex: 'path',
    width: 100,
  },
  {
    title: '耗时',
    dataIndex: 'costTime',
    width: 120
  },
  {
    title: '请求参数',
    dataIndex: 'query',
    ellipsis: true
  },
  {
    title: '请求负载',
    dataIndex: 'body',
    ellipsis: true
  },
  {
    title: '浏览器标识',
    dataIndex: 'userAgent',
  },
  {
    title: '开始时间',
    dataIndex: 'startTime',
    width: 180
  },
]

// 获取数据
const fetchData = async () => {
  loading.value = true
  try {
    const params = {
      currPage: pagination.current,
      pageSize: pagination.pageSize,
      startTime: searchForm.startTime,
      endTime: searchForm.endTime,
      account:searchForm.account
    }
    const res = await getOperationLog(params)
    if (res.data.code === 2000) {
      tableData.value = res.data.data.logs
      pagination.total = res.data.data.total
    } else {
      tableData.value = []
      pagination.total = 0
    }
  } catch (error) {
    tableData.value = []
    pagination.total = 0
  } finally {
    loading.value = false
  }
}

// 时间范围变化处理
const handleTimeRangeChange = (dates) => {
  if (dates) {
    searchForm.startTime = dates[0].format('YYYY-MM-DDTHH:mm:ssZ')
    searchForm.endTime = dates[1].format('YYYY-MM-DDTHH:mm:ssZ')
  } else {
    searchForm.startTime = ''
    searchForm.endTime = ''
  }
}

// 搜索
const handleSearch = () => {
  pagination.current = 1
  fetchData()
}

// 重置搜索
const resetSearch = () => {
  searchForm.timeRange = []
  searchForm.startTime = ''
  searchForm.endTime = ''
  handleSearch()
}

// 表格变化处理
const handleTableChange = (pag) => {
  pagination.current = pag.current
  pagination.pageSize = pag.pageSize
  console.log("测试1",pagination)
  fetchData()
}

// 日期格式化
const formatDate = (dateString) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  return date.toLocaleString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit',
    hour12: false
  })
}

// 组件挂载时获取数据
onMounted(() => {
  fetchData()
})
</script>

<style scoped>
.operation-container {
  padding: 24px;
}
.search-form {
  margin-bottom: 24px;
}
.operation-table {
  margin-top: 16px;
}
</style> 
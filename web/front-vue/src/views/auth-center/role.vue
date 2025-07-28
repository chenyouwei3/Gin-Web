<!-- 角色管理页面组件 -->
<template>
    <NavigationBar />
    <div class="role-container">
      <!-- 角色管理卡片 -->
      <a-card title="角色管理" :bordered="false">
        <!-- 搜索表单区域 -->
        <a-form layout="inline" :model="searchForm" class="search-form">
          <!-- 角色名称搜索框 -->
          <a-form-item label="角色名称">
            <a-input
              v-model:value="searchForm.name"
              placeholder="请输入角色名称"
              allow-clear
            />
          </a-form-item>
          <!-- 创建时间范围选择器 -->
          <a-form-item label="创建时间">
            <a-range-picker
              v-model:value="searchForm.timeRange"
              @change="onTimeRangeChange"
              allow-clear
            />
          </a-form-item>
          <!-- 搜索操作按钮组 -->
          <a-form-item>
            <a-space>
              <a-button type="primary" @click="handleSearch">查询</a-button>
              <a-button @click="resetSearch">重置</a-button>
            </a-space>
          </a-form-item>
        </a-form>

        <!-- 操作按钮区域 -->
        <div class="operation-bar">
          <a-button type="primary" @click="showAddModal">新增角色</a-button>
        </div>

        <!-- 角色列表表格 -->
        <a-table
          :columns="columns"
          :data-source="tableData"
          :loading="loading"
          :pagination="pagination"
          @change="handleTableChange"
          row-key="id"
        >
          <!-- 自定义列渲染 -->
          <template #bodyCell="{ column, record }">
            <!-- 操作列渲染 -->
            <template v-if="column.key === 'action'">
              <a-space>
                <a @click="showEditModal(record)">编辑</a>
                <a-divider type="vertical" />
                <a-popconfirm
                  title="确定要删除这个角色吗？"
                  @confirm="handleDelete(record)"
                >
                  <a class="danger-link">删除</a>
                </a-popconfirm>
              </a-space>
            </template>
            <!-- 状态列渲染 -->
            <template v-else-if="column.dataIndex === 'status'">
              <a-tag :color="record.status === 200 ? 'success' : 'error'">
                {{ record.status }}
              </a-tag>
            </template>
            <!-- 时间列渲染 -->
            <template v-else-if="column.dataIndex === 'created_at'||column.dataIndex==='updated_at'">
              {{ formatDate(record[column.dataIndex]) }}
            </template>
          </template>
        </a-table>
      </a-card>

      <!-- 新增/编辑角色对话框 -->
      <a-modal
        v-model:open="modalVisible"
        :title="modalTitle"
        @ok="handleModalOk"
        @cancel="handleModalCancel"
        :confirmLoading="modalLoading"
      >
        <!-- 角色表单 -->
        <a-form :model="formData" :rules="rules" ref="formRef">
          <!-- 角色名称输入框 -->
          <a-form-item label="角色名称" name="name">
            <a-input v-model:value="formData.name" placeholder="请输入角色名称" />
          </a-form-item>
          <!-- 角色描述输入框 -->
          <a-form-item label="角色描述" name="desc">
            <a-textarea v-model:value="formData.desc" placeholder="请输入角色描述" />
          </a-form-item>
        </a-form>
      </a-modal>
    </div>
  </template>
  
<script setup>
import {reactive, onMounted } from 'vue'
import NavigationBar from '@/components/NavigationBar.vue'
import {roleInsert, roleEdit} from '@/tools/api'
import { msgSuccess } from '@/tools/message'
import { useTable } from '@/tools/page/common'
import {formatDate,newTimeRangeHandler}from '@/tools/page/time'
import {roleColumns,roleRuleColumns} from '@/tools/page/columns'
/*----------------------------------------全局变量----------------------------------------*/
const {
    modalVisible,  //编辑表单确认
    modalTitle,   //编辑表单主题
    modalLoading,   //编辑表单加载
    formRef,     //编辑表单引用
    tableData,//表格数据
    loading,    //表格加载状况
    pagination,  //分页数据
    searchForm, //搜索表单
    fetchData,  //获取数据
    handleDelete,   //删除数据
    resetSearch,  //重置搜索表达
    handleSearch,  //处理搜索操作
    handleTableChange,  // 处理表格变化（分页、排序等）
    handleModalCancel,   // 处理弹窗取消操作
} = useTable('roles')
const onTimeRangeChange = newTimeRangeHandler(searchForm)   //
const columns = roleColumns// 表格列定义
const rules =roleRuleColumns
// 表单数据
const formData = reactive({
  name: '',
  desc: '',
})
// 显示编辑角色弹窗
const showEditModal = async (record) => {
  modalTitle.value = '编辑角色'
  formData.id = record.id
  formData.name = record.name
  formData.desc = record.desc
  modalVisible.value = true
}
// 处理弹窗确认操作
const handleModalOk = async () => {
  try {
    await formRef.value.validate()
    modalLoading.value = true
    const params = {
      id: formData.id,
      name: formData.name,
      desc: formData.desc, 
    }
    if (formData.id) {
      await roleEdit(params)
      msgSuccess('编辑成功')
    }else{
      await roleInsert(params)
      msgSuccess('新增成功')
    }
    modalVisible.value = false
    fetchData(pagination,searchForm)
    }catch(error) {
      console.error('操作失败:', error)
    } finally {
      modalLoading.value = false
  }
}
/*[增加]*/
const showAddModal = () => {
  modalTitle.value = '新增角色'
  formData.name = ''
  formData.desc = ''
  modalVisible.value = true
}

/*[默认函数]*/  
onMounted(() => {
  fetchData(pagination,searchForm)
})
</script>
  
<style scoped>
</style> 
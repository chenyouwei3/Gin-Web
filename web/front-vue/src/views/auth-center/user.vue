<!-- 用户管理页面组件 -->
<template>
    <NavigationBar />
    <div class="user-container">
      <!-- 用户管理卡片 -->
      <a-card title="用户管理" :bordered="false">
        <!-- 搜索表单区域 -->
        <a-form layout="inline" :model="searchForm" class="search-form">
          <!-- 用户名搜索框 -->
          <a-form-item label="用户名">
            <a-input
              v-model:value="searchForm.name"
              placeholder="请输入用户名"
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
        </div>

        <!-- 用户列表表格 -->
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
                  title="确定要删除这个用户吗？"
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
            <template v-else-if="column.dataIndex === 'create_time'||column.dataIndex==='update_time'">
              {{ formatDate(record[column.dataIndex]) }}
            </template>
          </template>
        </a-table>
      </a-card>

      <!-- 新增/编辑用户对话框 -->
      <a-modal
        v-model:open="modalVisible"
        :title="modalTitle"
        @ok="handleModalOk"
        @cancel="handleModalCancel"
        :confirmLoading="modalLoading"
      >
        <!-- 用户表单 -->
        <a-form :model="formData" :rules="rules" ref="formRef">
          <!-- 用户名输入框 -->
          <a-form-item label="用户名" name="name">
            <a-input v-model:value="formData.name" placeholder="请输入用户名" />
          </a-form-item>
          <!-- 邮箱输入框 -->
          <a-form-item label="邮箱" name="email">
            <a-input v-model:value="formData.email" placeholder="请输入邮箱" />
          </a-form-item>
          <!-- 角色选择穿梭框 -->
          <a-form-item label="角色">
            <a-transfer
              v-model:target-keys="formData.roleIds"
              :data-source="roleListData"
              :titles="['可选角色', '已选角色']"
              :render="item => item.title"
            />
          </a-form-item>
        </a-form>
      </a-modal>
    </div>
  </template>
  
<script setup>
// 导入必要的组件和工具
import { ref, reactive, onMounted } from 'vue'
import {  userEdit, getUserByRoles, roleList } from '@/tools/api'
import { msgSuccess, msgError } from '@/tools/message'
import NavigationBar from '@/components/NavigationBar.vue'
import { useTable } from '@/tools/page/common'
import {formatDate,newTimeRangeHandler}from '@/tools/page/time'
import {userColumns,userRuleColumns} from '@/tools/page/columns'
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
} = useTable('users')
const onTimeRangeChange = newTimeRangeHandler(searchForm)
const columns = userColumns //表格定义
const rules =userRuleColumns
 // 表单数据
const formData = reactive({ 
  id: null,
  name: '',
  email: '',
  roleIds: [],
  originalRoleIds: [],
})
// 角色列表数据
const roleListData = ref([])
// 显示编辑用户弹窗
const showEditModal = async (record) => {
  modalTitle.value = '编辑用户'
  formData.id = record.id
  formData.name = record.name
  formData.email = record.email
    
  try {
    // 获取用户已有的角色
    const res = await getUserByRoles(record.id)
    if(res?.data?.code === 2000) {
      formData.roleIds = res.data.data.map(role => String(role.id))
      // 保存原始角色列表用于比较
      formData.originalRoleIds = [...formData.roleIds]
    }else {
      formData.roleIds = []
      formData.originalRoleIds = []
    }
  }catch (error) {
    console.error('获取用户角色失败:', error)
    formData.roleIds = []
    formData.originalRoleIds = []
  }
  modalVisible.value = true
}
// 处理弹窗确认操作
const handleModalOk = async () => {
  try {
    await formRef.value.validate()
    modalLoading.value = true 
    // 计算新增和删除的角色
    const addRoles = formData.roleIds.filter(id => !formData.originalRoleIds?.includes(id))
    const deletedRoles = formData.originalRoleIds?.filter(id => !formData.roleIds.includes(id)) || []
    const params = {
      user: {
        id: formData.id,
        name: formData.name,
        email: formData.email,
      },
      addRoles: addRoles.map(id => Number(id)),
      deletedRoles: deletedRoles.map(id => Number(id))
    }
    await userEdit(params)
    msgSuccess('编辑成功')
    modalVisible.value = false
    fetchData()
  } catch (error) {
    console.error('操作失败:', error)
    msgError('操作失败')
  } finally {
    modalLoading.value = false
  }
}
// 获取角色列表
const fetchRoleList = async () => {
  try {
    const res = await roleList({ currPage: 1, pageSize: 1000 })
    if (res?.data?.code === 2000 && res?.data?.data?.roles) {
      roleListData.value = res.data.data.roles.map(role => ({
        key: String(role.id),
        title: role.name,
        description: role.desc,
      }))
    } else {
      console.error('获取角色列表失败: 数据结构不正确', res)
    }
  } catch (error) {
    console.error('获取角色列表失败:', error)
  }
} 
/*[默认函数]*/  
onMounted(() => {
  fetchData()
  fetchRoleList()
})
</script>
  
  <style scoped>
  /* 容器样式 */
  .user-container {
    padding: 24px;
  }
  /* 搜索表单样式 */
  .search-form {
    margin-bottom: 24px;
  }
  /* 操作按钮区域样式 */
  .operation-bar {
    margin-bottom: 16px;
  }
  /* 危险操作链接样式 */
  .danger-link {
    color: #ff4d4f;
  }
  .danger-link:hover {
    color: #ff7875;
  }
  </style> 
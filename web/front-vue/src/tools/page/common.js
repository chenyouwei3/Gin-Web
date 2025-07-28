import {ref, reactive } from 'vue'
import { roleList, userList,roleRemove,userRemove} from '@/tools/api'
export const useTable = ( dataKey = 'list') => {
  const loading = ref(false)//数据加载情况
  const tableData = ref([])//表格情况
  const modalVisible = ref(false)// 弹窗显示状态
  const modalTitle = ref('新增') // 默认值
  const modalLoading = ref(false)// 弹窗加载状态
  const formRef = ref(null)// 表单引用
  const pagination = reactive({
    current: 1,
    pageSize: 10,
    total: 0,
    showSizeChanger: true,
    showQuickJumper: true,
    showTotal: total => `共 ${total} 条记录`
  })
  let getFn = userList    
  let deleteFn=userRemove
  let keywordField = 'email'
  switch (dataKey){
  case "roles":
    getFn=roleList
    deleteFn=roleRemove
    keywordField = ''
    modalTitle.value='新增角色'
    break
  case "user":
    getFn=userList
    deleteFn=userRemove
    deleteFn=userRemove
    keywordField = 'email'
    modalTitle.value='新增用户'
    break
  default:

  }
  const searchForm = reactive({
  timeRange: [],
    startTime: '',
    endTime: '',
    name:'',
    [keywordField]: ''  // 动态字段名：可以是 name 或 username
  })

  const fetchData = async () => {
    loading.value = true
    try {
      const params = {
        currPage: pagination.current,
        pageSize: pagination.pageSize,
        startTime: searchForm.startTime,
        endTime: searchForm.endTime,
        name: searchForm.name,
        [keywordField]: searchForm[keywordField]
      }

      const res = await getFn(params)
      if (res?.data?.code === 2000) {
        tableData.value = res.data.data[dataKey] || []
        pagination.total = res.data.data.total || 0
      } else {
        tableData.value = []
        pagination.total = 0
      }
    } catch (error) {
      console.error('请求失败:', error)
      tableData.value = []
      pagination.total = 0
    } finally {
      loading.value = false
    }
  }

  const handleDelete = async (record) => {
  try {
      await deleteFn({ id: record.id })
      fetchData()
    } catch (error) {
    }
  }
  
  const resetSearch = () => {
    searchForm.timeRange = []
    searchForm.startTime = ''
    searchForm.endTime = ''
    searchForm.name = ''
    handleSearch()
  }

  const handleSearch = () => {
    pagination.current = 1
    fetchData()
  }

  const handleTableChange = (pag) => {
    pagination.current = pag.current
    pagination.pageSize = pag.pageSize
    fetchData()
  }
  const handleModalCancel = () => {
    modalVisible.value = false
    formRef.value?.resetFields()
  }

  return {
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
  }
}
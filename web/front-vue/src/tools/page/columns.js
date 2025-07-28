// src/utils/tableColumns.js

//角色渲染
export const roleColumns = [
  {
    title: '角色名称',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: '描述',
    dataIndex: 'desc',
    key: 'desc',
  },
  {
    title: '创建时间',
    dataIndex: 'created_at',
    key: 'createTime',
  },
  {
    title: '更新时间',
    dataIndex: 'updated_at',
    key: 'updateTime',
  },
  {
    title: '操作',
    key: 'action',
    width: 150,
  },
]


export const roleRuleColumns = {
    name: [{ required: true, message: '请输入角色名称', trigger: 'blur' }],
    desc: [{ required: true, message: '请输入角色描述', trigger: 'blur' }],
}


export const userColumns = [
  {
    title: '用户名',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: '账号',
    dataIndex: 'account',
    key: 'account',
  },
  {
    title: '性别',
    dataIndex: 'sex',
    key: 'sex',
  },
  {
    title: '邮箱',
    dataIndex: 'email',
    key: 'email',
  },
  {
    title: '创建时间',
    dataIndex: 'create_time',
    key: 'createTime',
  },
  {
    title: '更新时间',
    dataIndex: 'update_time',
    key: 'updateTime',
  },
  {
    title: '操作',
    key: 'action',
    width: 150,
  },
]


export const userRuleColumns = {
   name: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
   email: [
     { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
  ],
}
  
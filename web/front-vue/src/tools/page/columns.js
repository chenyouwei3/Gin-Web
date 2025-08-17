//登录界面
export const loginRules={
  account: [{ required: true, message: '用户名必填', trigger: 'blur' }],
  password: [{ required: true, message: '密码必填', trigger: 'blur' }]
}
//角色界面
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

//用户界面
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
    title: '邮箱',
    dataIndex: 'email',
    key: 'email',
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

export const userRuleColumns = {
   name: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
   email: [
     { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
  ],
}
  
//操作日志界面
export const logByOperationColumns=[// 表格列定义
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
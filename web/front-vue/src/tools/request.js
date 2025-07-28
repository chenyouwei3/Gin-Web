import axios from 'axios';
import storage from 'store'; 
import { msgSuccess,msgError } from './message'; 
import { VueAxios } from './axios'; 
// 创建 axios 实例
const request = axios.create({
  baseURL: "http://localhost:8080", // 正确的 API 地址
  timeout: 5000, //请求超时时间
});

// 异常拦截处理器
const errorHandler = (error) => {
  if (error.response) {
    if (error.response.status === 401) {//返回响应错误处理
      // 如果是 401 错误（授权失败）
      msgError('授权验证失败'); // 提示用户授权失败
      const token = storage.get('Access-Token'); // 获取存储中的 token
      if (token) {
        storage.remove('Access-Token'); // 从本地存储清除 token
      }
      setTimeout(() => {  //1s后执行一个函数
        window.location.reload(); // 刷新页面
      }, 1000);
    } else {
      msgError(error.response.statusText); // 显示其他错误
    }
  }
  return Promise.reject(error); // 返回拒绝的 promise
};

// 请求拦截器
request.interceptors.request.use(
  (config) => {
    const token = storage.get('Access-Token'); // 从 localStorage 中获取 token
    if (token) {
      config.headers['Access-Token'] = token; // 如果 token 存在，添加到请求头中
    }
    return config; // 返回配置
  },
  errorHandler // 如果请求发生错误，调用 errorHandler
);

// 响应拦截器
request.interceptors.response.use(
  (response) => {
   /*  if (response.data.message['en-US'] == 'success')  {
      msgSuccess(response.data.message['zh-CN']);
    } else {
      msgError(response.data.message['zh-CN']);
    } */
    return response; // 如果没有 success 字段，直接返回响应
  },
  errorHandler // 如果响应发生错误，调用 errorHandler
);

// 插件安装器
const installer = {
  install(app) {//当插件
    app.use(VueAxios, request); // 使用 VueAxios 插件，将 Axios 实例传给 Vue
  },
};

export default request; // 导出配置好的 Axios 实例
//installer命名为VueAxios  request命名为axios
export { installer as VueAxios, request as axios }; // 导出 VueAxios 插件和 Axios 实例

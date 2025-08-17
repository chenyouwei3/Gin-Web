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
request.interceptors.request.use((config) => {
  const access_token = storage.get('access_token'); 
  const refresh_token = storage.get('refresh_token'); 
  if (access_token) {
    config.headers['access_token'] = access_token; 
    config.headers['refresh_token'] = refresh_token; 
  }
    return config; 
  },
  errorHandler 
);

// 响应拦截器
// 响应拦截器（加入防重处理）
let lastMessageTime = 0;
let lastMessageContent = '';

request.interceptors.response.use(
  (response) => {
    const msgEn = response.data.message?.['en-US'];
    const msgZh = response.data.message?.['zh-CN'];
    const respCode=response.data.code
    // 只有在非登录接口时才自动显示消息
    if  (respCode) {
      const now = Date.now();
      // 1秒内相同内容不重复提示
      if (now - lastMessageTime < 1000 && respCode === lastMessageContent) {
        // 重复消息，跳过
      } else {
        if (msgEn === 'success') {
          msgSuccess(msgZh);
        } else {
          msgError(msgZh);
        }
        lastMessageContent = respCode;
        lastMessageTime = now;
      }
    }

    return response.data;
  },
  errorHandler
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

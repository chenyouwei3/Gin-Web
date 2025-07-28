import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import Antd from 'ant-design-vue'
import 'ant-design-vue/dist/reset.css'
import './assets/css/common.css'
//全局导航守卫
//参数解释: 
//to：即将要进入的目标路由对象。
//from：当前导航正要离开的路由对象。
//next：必须调用的函数，用来控制路由的跳转。你可以传递以下几种参数给 next：
router.beforeEach((to,from,next)=>{
    if (to.meta.title) {
        document.title = to.meta.title // 更新页面标题
    }
    next() // 跳转到目标路由
})


const app=createApp(App)//创造vue3应用实例
app.use(Antd) // 全局注册 Ant Design Vue 组件
app.use(router)//注册路由插件
app.mount('#app')// 挂载 Vue 应用到页面中的 #app 元素

import { createRouter, createWebHashHistory, createWebHistory } from "vue-router"

const routes = [
	//默认目录
	{
		path:'/',
		redirect:'/role-center'
	},
	{
		name:"404",
		path:"/404",
		meta:{
			title:"404"
		},
		component: () => import("@/views/default/404.vue")
	},
	/*---------------------首页---------------------*/
	{
		name:"login",
		path:"/login",
		meta:{
			title:"账号登录"
		},
		component: () => import("@/views/login.vue")
	},
	/*---------------------功能页面---------------------*/

	//日志中心
	{
		name:"log-operation",
		path:"/log-operation",
		meta:{
			title:"操作日志"
		},
		component: () => import('@/views/log-center/operation.vue')
	},
	//权限中心
	{
		name:"role-center",
		path:"/role-center",
		meta:{
			title:"角色管理"
		},
		component: () => import('@/views/auth-center/role.vue')
	},
	{
		name:"user-center",
		path:"/user-center",
		meta:{
			title:"用户管理"
		},
		component: () => import('@/views/auth-center/user.vue')
	},
	// 捕获所有未匹配的路由，重定向到 404 页面
	{
		path: "/:pathMatch(.*)*",
		redirect: "/404"
	}
]

const router = createRouter({
	//使用url的#符号之后的部分模拟url路径的变化,因为不会触发页面刷新,所以不需要服务端支持
	//history: createWebHashHistory(), 
	history: createWebHistory(),
	routes
})

export default router
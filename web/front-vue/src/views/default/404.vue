<template>
	<div class="not-found-page">
		<!-- 顶部动画区域 -->
		<div class="top-section">
			<div class="clouds">
				<div class="cloud cloud1"></div>
				<div class="cloud cloud2"></div>
				<div class="cloud cloud3"></div>
				<div class="cloud cloud4"></div>
			</div>
			<div class="animation-container">
				<img src="@/assets/logo.png" alt="logo" class="logo" />
				<div class="title-container">
					<div class="title-wrapper">
						<h1 class="main-title">404</h1>
						<div class="author-tag">
							<span class="author-name">页面未找到</span>
						</div>
					</div>
					<p class="sub-title">抱歉，您访问的页面不存在</p>
				</div>
			</div>
		</div>

		<!-- 中间内容区域 -->
		<div class="middle-section">
			<div class="not-found-content">
				<div class="not-found-left">
					<img src="@/assets/title.png" alt="title" class="title-image" />
				</div>
				<div class="not-found-divider"></div>
				<div class="not-found-right">
					<div class="not-found-message">
						<p>您可能输入了错误的网址，或者该页面已被删除。</p>
						<p>请检查您的输入，或返回首页继续浏览。</p>
					</div>
					<a-button 
						type="primary" 
						size="large" 
						class="login-button" 
						href="/"
					>
						返回首页
					</a-button>
				</div>
			</div>
		</div>

		<!-- 底部信息区域 -->
		<div class="bottom-section">
			<div class="bottom-content">
				<div class="register-link">
					<a href="https://github.com/chenyouwei3" target="_blank">个人开发者作品</a>
				</div>
				<div class="footer-links">
					<a href="/help-center">帮助中心</a>
					<a href="/privacy-policy">隐私政策</a>
					<a href="/terms-of-service">服务条款</a>
				</div>
			</div>
		</div>
	</div>
</template>

<script setup>
// Vue 3 的 script setup 语法

</script>

<style scoped>
.not-found-page {
	min-height: 100vh;
	display: flex;
	flex-direction: column;
	background-color: #f0f2f5;
}

/* 顶部动画区域样式 */
.top-section {
	position: relative;
	height: 300px;
	background: linear-gradient(135deg, #1890ff 0%, #36cfc9 100%);
	overflow: hidden;
}

.clouds {
	position: absolute;
	top: 0;
	left: 0;
	width: 100%;
	height: 100%;
}

.cloud {
	position: absolute;
	background: rgba(255, 255, 255, 0.8);
	border-radius: 50%;
	animation: float 15s infinite linear;
}

.cloud1 {
	width: 100px;
	height: 60px;
	top: 20%;
	left: 10%;
	animation-duration: 20s;
}

.cloud2 {
	width: 150px;
	height: 90px;
	top: 40%;
	left: 30%;
	animation-duration: 25s;
}

.cloud3 {
	width: 120px;
	height: 70px;
	top: 60%;
	left: 50%;
	animation-duration: 18s;
}

.cloud4 {
	width: 80px;
	height: 50px;
	top: 30%;
	left: 70%;
	animation-duration: 22s;
}

@keyframes float {
	0% {
		transform: translateX(-100%);
	}
	100% {
		transform: translateX(100vw);
	}
}

.animation-container {
	position: relative;
	z-index: 1;
	display: flex;
	flex-direction: column;
	align-items: center;
	padding-top: 40px;
}

.logo {
	width: 80px;
	height: 80px;
	margin-bottom: 20px;
}

.title-container {
	text-align: center;
	color: white;
}

.main-title {
	font-size: 48px;
	font-weight: bold;
	margin: 0;
}

.author-tag {
	margin-top: 10px;
	font-size: 16px;
}

.author-name {
	color: rgba(255, 255, 255, 0.8);
}

.sub-title {
	margin-top: 10px;
	font-size: 18px;
	color: rgba(255, 255, 255, 0.8);
}

/* 中间内容区域样式 */
.middle-section {
	flex: 1;
	display: flex;
	justify-content: center;
	align-items: center;
	padding: 40px 20px;
}

.not-found-content {
	display: flex;
	background-color: white;
	border-radius: 10px;
	box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
	width: 800px;
	min-height: 400px;
}

.not-found-left {
	flex: 1;
	display: flex;
	justify-content: center;
	align-items: center;
	padding: 40px;
}

.title-image {
	max-width: 100%;
	max-height: 300px;
}

.not-found-divider {
	width: 1px;
	background-color: #f0f0f0;
	margin: 40px 0;
}

.not-found-right {
	flex: 1;
	display: flex;
	flex-direction: column;
	justify-content: center;
	padding: 40px;
}

.not-found-message {
	text-align: center;
	margin-bottom: 30px;
}

.not-found-message p {
	color: #666;
	font-size: 16px;
	line-height: 1.6;
	margin: 10px 0;
}

.login-button {
	width: 100%;
	height: 40px;
	font-size: 16px;
	background: linear-gradient(135deg, #1890ff 0%, #36cfc9 100%);
	border: none;
	border-radius: 4px;
	transition: all 0.3s;
}

.login-button:hover {
	opacity: 0.9;
	transform: translateY(-2px);
	box-shadow: 0 4px 12px rgba(24, 144, 255, 0.4);
}

/* 底部信息区域样式 */
.bottom-section {
	background-color: white;
	padding: 20px;
	text-align: center;
}

.bottom-content {
	max-width: 1200px;
	margin: 0 auto;
}

.register-link {
	color: #666;
	margin-bottom: 10px;
}

.footer-links {
	margin-top: 10px;
}

.footer-links a {
	color: #666;
	margin: 0 10px;
	text-decoration: none;
}

.footer-links a:hover {
	color: #1890ff;
}
</style>
  
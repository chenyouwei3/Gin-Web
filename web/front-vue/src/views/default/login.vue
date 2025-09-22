<template>
  <div class="login-page">
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
            <h1 class="main-title">个人云盘</h1>
            <div class="author-tag">
              <span class="author-name">D-16</span>
              
            </div>
          </div>
          <p class="sub-title">安全、高效、便捷的云存储服务</p>
        </div>
      </div>
    </div>

    <!-- 中间登录区域 -->
    <div class="middle-section">
      <div class="login-content">
        <div class="login-left">
          <img src="@/assets/title.png" alt="title" class="title-image" />
        </div>
        <div class="login-divider"></div>
        <div class="login-right">
          <div class="login-methods">
            <div 
              class="method-tab" 
              :class="{ active: activeTab === 'password' }"
              @click="activeTab = 'password'"
            >
              登录
            </div>
            <div class="method-divider"></div>
            <div 
              class="method-tab" 
              :class="{ active: activeTab === 'sms' }"
              @click="activeTab = 'sms'"
            >
              注册
            </div>
          </div>

          <!-- 密码登录表单 -->
          <div v-show="activeTab === 'password'" class="login-form">
            <a-form ref="loginFormModel" :model="loginInfo" :rules="rules">
              <a-form-item name="account">
                <a-input 
                  size="large" 
                  v-model:value="loginInfo.account" 
                  placeholder="请输入账号" 
                  @keyup.enter="handleLogin"
                  class="custom-input"
                >
                  <template #prefix>
                    <UserOutlined class="input-icon" />
                  </template>
                </a-input>
              </a-form-item>

              <a-form-item name="password">
                <a-input-password 
                  size="large" 
                  v-model:value="loginInfo.password" 
                  placeholder="请输入密码" 
                  @keyup.enter="handleLogin"
                  class="custom-input"
                  id="login_password"
                >
                  <template #prefix>
                    <LockOutlined class="input-icon" />
                  </template>
                </a-input-password>
              </a-form-item>

              <div class="login-options">
                <a-checkbox>记住密码</a-checkbox>
                <a class="forget-password">忘记密码？</a>
              </div>

              <a-form-item>
                <a-button 
                  type="primary" 
                  size="large" 
                  class="login-button" 
                  @click="handleLogin" 
                  :loading="loadingLogin"
                >
                  登录
                </a-button>
              </a-form-item>
            </a-form>
          </div>

          <!-- 短信登录表单 -->
          <div v-show="activeTab === 'sms'" class="login-form">
            <a-form ref="smsFormModel" :model="smsInfo" :rules="smsRules">
             <a-form-item name="name">
                <a-input 
                  size="large" 
                  v-model:value="smsInfo.name" 
                  placeholder="请输入用户名" 
                  class="custom-input"
                >
                  <template #prefix>
                    <UserOutlined class="input-icon" />
                  </template>
                </a-input>
              </a-form-item>
              <a-form-item name="email">
                <a-input 
                  size="large" 
                  v-model:value="smsInfo.email" 
                  placeholder="请输入邮箱" 
                  class="custom-input"
                >
                  <template #prefix>
                    <MailOutlined class="input-icon" />
                  </template>
                </a-input>
              </a-form-item>

              <a-form-item name="password">
                <a-input-password 
                  size="large" 
                  v-model:value="smsInfo.password" 
                  placeholder="请输入密码" 
                  class="custom-input"
                  id="register_password"
                >
                  <template #prefix>
                    <LockOutlined class="input-icon" />
                  </template>
                </a-input-password>
              </a-form-item>

              <a-form-item name="confirmPassword">
                <a-input-password 
                  size="large" 
                  v-model:value="smsInfo.confirmPassword" 
                  placeholder="请确认密码" 
                  class="custom-input"
                >
                  <template #prefix>
                    <LockOutlined class="input-icon" />
                  </template>
                </a-input-password>
              </a-form-item>
              <a-form-item>
                <a-button 
                  type="primary" 
                  size="large" 
                  class="login-button" 
                  @click="handleRegister" 
                  :loading="loadingLogin"
                >
                  注册
                </a-button>
              </a-form-item>
            </a-form>
          </div>
        </div>
      </div>
    </div>

    <!-- 底部信息区域 -->
    <div class="bottom-section">
      <div class="bottom-content">
        <div class="developer-link">
          <a href="https://github.com/chenyouwei3" target="_blank">个人开发者作品</a>
        </div>
        <div class="other-login">
          <span>其他登录方式：</span>
          <a-space>
            <a-button type="link" class="social-login">
              <template #icon><WechatOutlined /></template>
              微信
            </a-button>
            <a-button type="link" class="social-login">
              <template #icon><QqOutlined /></template>
              QQ
            </a-button>
            <a-button type="link" class="social-login">
              <template #icon><TwitterOutlined /></template>
              X
            </a-button>
            <a-button type="link" class="social-login">
              <template #icon><InstagramOutlined /></template>
              Instagram
            </a-button>
            <a-button type="link" class="social-login">
              <template #icon><GoogleOutlined /></template>
              Google
            </a-button>
            <a-button type="link" class="social-login">
              <template #icon><FacebookOutlined /></template>
              Facebook
            </a-button>
            <a-button type="link" class="social-login">
              <template #icon><GithubOutlined /></template>
              GitHub
            </a-button>
          </a-space>
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
import storage from 'store'   
import { login, register } from '@/tools/api'; 
import {msgError, msgSuccess } from '@/tools/message';
import { UserOutlined, LockOutlined, WechatOutlined, QqOutlined,MobileOutlined,SafetyOutlined,MailOutlined,TwitterOutlined,InstagramOutlined,GoogleOutlined,FacebookOutlined,GithubOutlined} from '@ant-design/icons-vue';   
import { ref,getCurrentInstance } from 'vue';  
import { useRouter } from 'vue-router'
import {loginRules} from '@/tools/page/columns'
/*--------------------------------- 全局变量 ---------------------------------*/
const activeTab = ref('password');
const { proxy } = getCurrentInstance() // 获取 Vue 实例（this）
const router = useRouter()
const rules = loginRules
const smsRules = {
  email: [
    { required: true, type: 'email', message: '请输入正确的邮箱地址', trigger: 'blur' }
  ],
  password: [
    { required: true, min: 6, message: '密码必填且不能少于6位', trigger: 'blur' }
  ],
  confirmPassword: [
    { required: true, message: '请确认密码', trigger: 'blur' },
    {
      validator: (_, value) => 
        value !== smsInfo.value.password 
          ? Promise.reject('两次输入的密码不一致') 
          : Promise.resolve(),
      trigger: 'blur'
    }
  ]
}

/* [处理注册]*/
const smsFormModel = ref(null);// 表单引用
const smsInfo = ref({email: "",password: "",confirmPassword: "",name: ""})//注册信息
const loadingLogin = ref(false);// 登录请求加载状态
const handleRegister = async () => {
  if (!smsInfo.value.email || !smsInfo.value.password || !smsInfo.value.confirmPassword || !smsInfo.value.name ) {
    msgError('请填写完整信息')
    return
  }
  if (smsInfo.value.password !== smsInfo.value.confirmPassword) {
    msgError('两次密码不一致')
    return
  }
  try {
    loadingLogin.value = true;
    const userRequest = {
    user: {
      name: smsInfo.value.name,
      email: smsInfo.value.email,
      account: smsInfo.value.email.split('@')[0],
      password: smsInfo.value.password,
      avatarUrl:""
    },
    addRoles: proxy.$default_sign_up_roles,        // 要添加的角色 ID 数组
    deletedRoles: []        // 要删除的角色 ID 数组
    };  
    const res = await register(userRequest);
    if (res?.code==2000){
      setTimeout(() => {   // 延迟跳转
        activeTab.value = 'password';  //切换代码框
        smsInfo.value = {
          email: "",
          password: "",
          confirmPassword: "",
          name: "",
          account: "",
        };}, 500);
    }  
  } finally {
    loadingLogin.value = false;
  }
};

/* [登录]*/
const loginFormModel = ref(null);// 表单引用
const loginInfo = ref({account: "",password: ""})//初始化用户信息
const handleLogin = async () => {
  try {
    await loginFormModel.value.validate();
    loadingLogin.value = true;    
    const ret = await login(loginInfo.value);
    if (ret?.code === 2000){
      storage.set("access_token", ret.data.access_token, 2 * 60 * 60 * 1000);//2 hour
      storage.set("refresh_token", ret.data.refresh_token, 1 * 60 * 60 * 1000);
      storage.set("user_info", {name: ret.data.user.name, account: ret.data.user.account}, 1 * 60 * 60 * 1000);
      setTimeout(() => {
        loadingLogin.value = false;
        router.push('/role-center');
      }, 500);
    } else {
      loadingLogin.value = false;
    }
  } catch (error) {
    loadingLogin.value = false;
  }
};
</script>


































<style scoped>
.login-page {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background: #f4f5f6;
}

.top-section {
  background: linear-gradient(135deg, #00a1d6 0%, #00b5e5 100%);
  padding: 40px 0;
  color: white;
  position: relative;
  overflow: hidden;
}

.clouds {
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  pointer-events: none;
}

.cloud {
  position: absolute;
  background: rgba(255, 255, 255, 0.8);
  border-radius: 50px;
}

.cloud::before,
.cloud::after {
  content: '';
  position: absolute;
  background: rgba(255, 255, 255, 0.8);
  border-radius: 50%;
}

.cloud1 {
  width: 100px;
  height: 30px;
  top: 20%;
  left: 10%;
}

.cloud1::before {
  width: 50px;
  height: 50px;
  top: -25px;
  left: 10px;
}

.cloud1::after {
  width: 30px;
  height: 30px;
  top: -15px;
  left: 40px;
}

.cloud2 {
  width: 150px;
  height: 40px;
  top: 40%;
  right: 15%;
}

.cloud2::before {
  width: 60px;
  height: 60px;
  top: -30px;
  left: 15px;
}

.cloud2::after {
  width: 40px;
  height: 40px;
  top: -20px;
  left: 50px;
}

.cloud3 {
  width: 80px;
  height: 25px;
  top: 60%;
  left: 25%;
}

.cloud3::before {
  width: 40px;
  height: 40px;
  top: -20px;
  left: 8px;
}

.cloud3::after {
  width: 25px;
  height: 25px;
  top: -12px;
  left: 35px;
}

.cloud4 {
  width: 120px;
  height: 35px;
  top: 30%;
  right: 30%;
}

.cloud4::before {
  width: 55px;
  height: 55px;
  top: -27px;
  left: 12px;
}

.cloud4::after {
  width: 35px;
  height: 35px;
  top: -17px;
  left: 45px;
}

.animation-container {
  position: relative;
  z-index: 1;
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 20px;
}

.logo {
  width: 60px;
  height: 60px;
}

.title-container {
  text-align: left;
}

.title-wrapper {
  display: flex;
  align-items: center;
  gap: 12px;
}

.main-title {
  font-size: 32px;
  font-weight: 600;
  margin: 0;
}

.author-tag {
  display: flex;
  align-items: center;
  background: rgba(255, 255, 255, 0.1);
  padding: 4px 12px;
  border-radius: 20px;
  backdrop-filter: blur(5px);
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.author-name {
  font-size: 14px;
  font-weight: 500;
  color: #fff;
  letter-spacing: 0.5px;
}

.author-dot {
  margin: 0 6px;
  color: rgba(255, 255, 255, 0.6);
  font-weight: 300;
}

.author-role {
  font-size: 12px;
  color: rgba(255, 255, 255, 0.8);
  font-style: italic;
  letter-spacing: 0.5px;
}

.sub-title {
  font-size: 16px;
  opacity: 0.8;
  margin: 8px 0 0;
}

.middle-section {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 40px 20px;
  background: #fff;
}

.login-content {
  display: flex;
  min-height: 500px;
  width: 800px;
}

.login-left {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 40px;
}

.title-image {
  max-width: 100%;
  max-height: 400px;
  object-fit: contain;
}

.login-divider {
  width: 1px;
  background: #e5e9ef;
}

.login-right {
  flex: 1;
  padding: 40px;
}

.login-methods {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
  border-bottom: 1px solid #e5e9ef;
}

.method-tab {
  padding: 10px 20px;
  cursor: pointer;
  color: #6d757a;
  font-size: 14px;
  flex: 1;
  text-align: center;
}

.method-divider {
  width: 1px;
  height: 16px;
  background: #e5e9ef;
}

.method-tab.active {
  color: #00a1d6;
  border-bottom: 2px solid #00a1d6;
}

.login-form {
  margin-top: 20px;
}

.custom-input {
  height: 40px;
  border-radius: 4px;
}

.input-icon {
  color: #6d757a;
}

.login-options {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
}

.forget-password {
  color: #00a1d6;
}

.login-button {
  width: 100%;
  height: 40px;
  background: #00a1d6;
  border: none;
  border-radius: 4px;
}

.login-button:hover {
  background: #00b5e5;
}

.sms-code-input {
  display: flex;
  gap: 10px;
}

.send-code-btn {
  white-space: nowrap;
  padding: 0 15px;
}

.bottom-section {
  background: #fff;
  padding: 20px 0;
  border-top: 1px solid #e5e9ef;
}

.bottom-content {
  max-width: 1200px;
  margin: 0 auto;
  text-align: center;
}

.register-link {
  color: #6d757a;
  margin-bottom: 15px;
}

.register-link a {
  color: #00a1d6;
  margin-left: 5px;
}

.other-login {
  margin-bottom: 15px;
}

.social-login {
  color: #6d757a;
}

.footer-links {
  display: flex;
  justify-content: center;
  gap: 20px;
}

.footer-links a {
  color: #6d757a;
  text-decoration: none;
}

.footer-links a:hover {
  color: #00a1d6;
}

.developer-link {
  margin-bottom: 20px;
}

.developer-link a {
  color: #666;
  text-decoration: none;
  font-size: 16px;
}

.developer-link a:hover {
  color: #1890ff;
}

</style>

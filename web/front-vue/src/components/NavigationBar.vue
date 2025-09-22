<template>
  <div class="nav-container">
    <div class="nav-buttons">
      <div class="nav-buttons-left">
        <a-button class="nav-btn" @click="handleNavigation('file-test')">
          <template #icon><cloud-outlined /></template>
          首页
        </a-button>
        <a-dropdown>
          <a-button class="nav-btn">
            <template #icon><setting-outlined /></template>
            权限中心
            <down-outlined />
          </a-button>
          <template #overlay>
            <a-menu>
              <a-menu-item key="role" @click="handleNavigation('role-center')">
                <template #icon><team-outlined /></template>
                角色管理
              </a-menu-item>
              <a-menu-item key="user" @click="handleNavigation('user-center')">
                <template #icon><user-outlined /></template>
                用户管理
              </a-menu-item>
            </a-menu>
          </template>
        </a-dropdown>
        <!-- 日志中心下拉菜单 -->
        <a-dropdown>
          <a-button class="nav-btn">
            <template #icon><file-text-outlined /></template>
            日志中心
            <down-outlined />
          </a-button>
          <template #overlay>
            <a-menu>
              <a-menu-item key="operation" @click="handleNavigation('log-operation')">
                <template #icon><history-outlined /></template>
                操作日志
              </a-menu-item>
            </a-menu>
          </template>
        </a-dropdown>
      </div>
    </div>
    <!-- 添加用户信息显示 -->
    <a-dropdown>
      <div class="user-info">
        <a-avatar :size="32" style="background-color: #00a1d6">
          {{ userInfo.name ? userInfo.name.charAt(0).toUpperCase() : 'U' }}
        </a-avatar>
        <span class="user-name">{{ userInfo.name }}</span>
        <down-outlined class="dropdown-icon" />
      </div>
      <template #overlay>
        <a-menu>
          <a-menu-item key="logout" @click="handleNavigation('personal')">
            <template #icon><ProfileOutlined /></template>
            个人信息
          </a-menu-item>
          <a-menu-item key="logout" @click="handleLogout">
            <template #icon><LogoutOutlined /></template>
            退出登录
          </a-menu-item>
        </a-menu>
      </template>
    </a-dropdown>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  CloudOutlined,
  ProfileOutlined,
  DownOutlined,
  LogoutOutlined,
  SettingOutlined,
  TeamOutlined,
  UserOutlined,
  FileTextOutlined,
  HistoryOutlined} from '@ant-design/icons-vue'
import storage from 'store'
import { useRouter } from 'vue-router'

const router = useRouter()
const userInfo = ref(storage.get('User-Info') || {})

const handleNavigation = (route) => {
  router.push(`/${route}`)
}

const handleLogout = () => {
  storage.remove('User-Info')
  storage.remove('Access-Token')
  handleNavigation('login')
}

</script>

<style lang="less" scoped>
.nav-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
  height: 60px;
  background-color: #fff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.nav-buttons {
  display: flex;
  justify-content: space-between;
  width: 100%;
}

.nav-buttons-left {
  display: flex;
  gap: 10px;
}

.nav-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 16px;
  border-radius: 4px;
  transition: all 0.3s;
}

.nav-btn:hover {
  background-color: #f0f0f0;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 0 12px;
  cursor: pointer;
  transition: all 0.3s;
}

.user-info:hover {
  background-color: #f6f7f8;
  border-radius: 4px;
}

.user-name {
  font-size: 14px;
  color: #18191c;
  font-weight: 500;
}

.dropdown-icon {
  font-size: 12px;
  color: #9499a0;
  transition: transform 0.3s;
}

.user-info:hover .dropdown-icon {
  color: #00a1d6;
}

:deep(.ant-dropdown-menu) {
  padding: 4px;
  border-radius: 4px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
}

:deep(.ant-dropdown-menu-item) {
  padding: 8px 12px;
  border-radius: 4px;
  transition: all 0.3s;
}

:deep(.ant-dropdown-menu-item:hover) {
  background-color: #f6f7f8;
  color: #00a1d6;
}

:deep(.ant-dropdown-menu-item .anticon) {
  margin-right: 8px;
}
</style> 
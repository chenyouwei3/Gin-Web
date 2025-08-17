<template>
  <div class="error-page">
    <div class="error-container">
      <div class="error-icon">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <circle cx="12" cy="12" r="10"></circle>
          <line x1="15" y1="9" x2="9" y2="15"></line>
          <line x1="9" y1="9" x2="15" y2="15"></line>
        </svg>
      </div>
      
      <h1 class="error-code">404</h1>
      <h2 class="error-title">页面未找到</h2>
      <p class="error-description">
        抱歉，您访问的页面不存在或已被移除。
      </p>
      
      <div class="error-actions">
        <button @click="goBack" class="btn-back">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M19 12H5M12 19l-7-7 7-7"></path>
          </svg>
          返回上一页
        </button>
        
        <button @click="goHome" class="btn-home">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path>
            <polyline points="9,22 9,12 15,12 15,22"></polyline>
          </svg>
          返回首页
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

const goBack = () => {
  if (window.history.length > 1) {
    router.go(-1)
  } else {
    router.push('/')
  }
}

const goHome = () => {
  router.push('/')
}
</script>

<style scoped>
.error-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #fb7299 0%, #fc8bab 100%);
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  position: relative;
  overflow: hidden;
}

/* Bilibili 风格的装饰元素 */
.error-page::before {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(circle, rgba(255, 255, 255, 0.1) 1px, transparent 1px);
  background-size: 20px 20px;
  animation: float 20s linear infinite;
}

@keyframes float {
  0% {
    transform: translate(0, 0);
  }
  100% {
    transform: translate(-20px, -20px);
  }
}

.error-container {
  text-align: center;
  padding: 3rem;
  background: rgba(255, 255, 255, 0.95);
  border-radius: 16px;
  box-shadow: 0 8px 32px rgba(251, 114, 153, 0.3);
  max-width: 500px;
  width: 90%;
  border: 2px solid #fb7299;
  position: relative;
  backdrop-filter: blur(10px);
}

.error-container::before {
  content: '';
  position: absolute;
  top: -2px;
  left: -2px;
  right: -2px;
  bottom: -2px;
  background: linear-gradient(45deg, #fb7299, #fc8bab, #fb7299);
  border-radius: 16px;
  z-index: -1;
  opacity: 0.3;
}

.error-icon {
  margin-bottom: 2rem;
}

.error-icon svg {
  width: 80px;
  height: 80px;
  color: #fb7299;
  opacity: 0.9;
  filter: drop-shadow(0 2px 4px rgba(251, 114, 153, 0.3));
}

.error-code {
  font-size: 4.5rem;
  font-weight: 700;
  color: #fb7299;
  margin: 0 0 0.5rem 0;
  line-height: 1;
  letter-spacing: -2px;
  text-shadow: 0 2px 4px rgba(251, 114, 153, 0.3);
}

.error-title {
  font-size: 1.5rem;
  font-weight: 600;
  color: #2d3748;
  margin: 0 0 1rem 0;
}

.error-description {
  font-size: 1rem;
  color: #718096;
  margin: 0 0 2.5rem 0;
  line-height: 1.6;
}

.error-actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
}

.btn-back,
.btn-home {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.875rem 1.75rem;
  border: 2px solid #fb7299;
  border-radius: 8px;
  font-size: 0.9rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  text-decoration: none;
  background: #ffffff;
  color: #fb7299;
  position: relative;
  overflow: hidden;
}

.btn-back::before,
.btn-home::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.4), transparent);
  transition: left 0.5s;
}

.btn-back:hover::before,
.btn-home:hover::before {
  left: 100%;
}

.btn-back:hover {
  background: #fb7299;
  color: #ffffff;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(251, 114, 153, 0.4);
}

.btn-home {
  background: #fb7299;
  color: #ffffff;
}

.btn-home:hover {
  background: #e85a87;
  border-color: #e85a87;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(251, 114, 153, 0.4);
}

.btn-back svg,
.btn-home svg {
  width: 16px;
  height: 16px;
}

/* 响应式设计 */
@media (max-width: 480px) {
  .error-container {
    padding: 2rem;
    margin: 1rem;
  }
  
  .error-code {
    font-size: 3.5rem;
  }
  
  .error-title {
    font-size: 1.25rem;
  }
  
  .error-actions {
    flex-direction: column;
  }
  
  .btn-back,
  .btn-home {
    width: 100%;
    justify-content: center;
  }
}

/* 动画效果 */
.error-container {
  animation: fadeInUp 0.6s ease-out;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.error-icon {
  animation: bounce 2s infinite;
}

@keyframes bounce {
  0%, 20%, 50%, 80%, 100% {
    transform: translateY(0);
  }
  40% {
    transform: translateY(-8px);
  }
  60% {
    transform: translateY(-4px);
  }
}

/* Bilibili 风格的额外装饰 */
.error-container::after {
  content: '♪';
  position: absolute;
  top: 10px;
  right: 15px;
  font-size: 1.2rem;
  color: #fb7299;
  opacity: 0.6;
  animation: sparkle 2s ease-in-out infinite;
}

@keyframes sparkle {
  0%, 100% {
    opacity: 0.6;
    transform: scale(1);
  }
  50% {
    opacity: 1;
    transform: scale(1.1);
  }
}
</style>
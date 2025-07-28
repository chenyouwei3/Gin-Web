// utils/message.js
import { message } from 'ant-design-vue';

// 统一成功提示
export function msgSuccess(content, duration = 0.5) {
  message.success({
    content,
    duration,
    style: {
      marginTop: '20vh',
      fontSize: '16px',
      color: '#52c41a'
    }
  });
}

// 统一错误提示
export function msgError(content, duration = 2) {
  message.error({
    content,
    duration,
    style: {
      marginTop: '20vh',
      fontSize: '16px',
      color: '#ff4d4f'
    }
  });
}

// 统一警告提示
export function msgWarn(content, duration = 2) {
  message.warning({
    content,
    duration,
    style: {
      marginTop: '20vh',
      fontSize: '16px',
      color: '#faad14'
    }
  });
}

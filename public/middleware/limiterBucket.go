package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// LeakBucket 定义了漏桶的结构
type LeakBucket struct {
	capacity     int           // 漏桶的最大容量（即能存放的请求数）
	remaining    int           // 当前漏桶中剩余的容量
	fillInterval time.Duration // 请求漏出的速率间隔
	lastLeakTime time.Time     // 上次漏出请求的时间
	mutex        sync.Mutex    // 锁，确保并发情况下的线程安全
}

// NewLeakBucket 创建一个新的漏桶
func NewLeakBucket(capacity int, fillInterval time.Duration) *LeakBucket {
	return &LeakBucket{
		capacity:     capacity,
		remaining:    capacity,
		fillInterval: fillInterval,
		lastLeakTime: time.Now(),
	}
}

// Allow 方法用来判断是否允许请求进入（返回true），并更新漏桶状态
func (b *LeakBucket) Allow() bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	// 计算从上次漏水到当前时间，漏出的请求数
	now := time.Now()
	elapsed := now.Sub(b.lastLeakTime)
	leaked := int(elapsed / b.fillInterval) // 漏出的请求数

	// 更新漏桶中的剩余容量
	if leaked > 0 {
		b.remaining += leaked
		if b.remaining > b.capacity {
			b.remaining = b.capacity // 防止超过最大容量
		}
		b.lastLeakTime = now // 更新上次漏出的时间
	}

	// 如果漏桶还有剩余容量，允许请求进入
	if b.remaining > 0 {
		b.remaining-- // 减少剩余容量
		return true
	}

	// 否则拒绝请求
	return false
}

// LeakBucketMiddleware 创建一个 Gin 中间件，用来限流
func LeakBucketMiddleware(capacity int, fillInterval time.Duration) gin.HandlerFunc {
	bucket := NewLeakBucket(capacity, fillInterval)
	return func(c *gin.Context) {
		// 判断请求是否被允许
		if !bucket.Allow() {
			// 如果请求被拒绝，返回 429 状态码和错误信息
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"message": "Too many requests, please try again later.",
			})
			return
		}
		// 如果允许请求，继续处理下一个中间件或路由处理
		c.Next()
	}
}

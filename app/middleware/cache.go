package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"loopy-manager/initialize/messageQueue"
	"loopy-manager/pkg/redisUtils"
	"net/http"
	"sync"
	"time"
)

type bodyLogWriter struct {
	mutex sync.Mutex
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	n, err := w.body.Write(b)
	if err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(b)
}

func CacheTest() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		path := c.Request.URL.Path
		cache := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = cache
		//命中缓存
		if method == "GET" && c.Writer.Status() == 200 {
			cacheData, err := redisUtils.Redis{}.GetValue(path)
			if err != nil {
				logrus.Error("获取redis缓存失败:", err)
				c.Next()
				if cache.body != nil {
					errOver := redisUtils.Redis{}.SetValue(path, cache.body.String(), 5*time.Minute)
					if errOver != nil {
						logrus.Error("更新redis缓存失败:", err)
					}
				}
				return
			}
			if cacheData != "" { //命中
				var jsonData interface{}
				err := json.Unmarshal([]byte(cacheData), &jsonData)
				if err != nil {
					logrus.Error("缓存转化json失败:", err)
					c.Next()
					return
				}
				c.JSON(http.StatusOK, jsonData)
				c.Abort()
				return
			}
		}
		//延迟双删
		err := redisUtils.Redis{}.DeletedValue(path)
		if err != nil {
			logrus.Error("删除redis缓存错误:", err)
			return
		}
		c.Next()
		err = redisUtils.Redis{}.DeletedValue(path)
		if err != nil {
			logrus.Error("删除redis缓存错误:", err)
			return
		}
	}
}

func CacheTest2() gin.HandlerFunc {
	return func(c *gin.Context) {
		//启动缓存删除消息队列
		method := c.Request.Method
		path := c.Request.URL.Path
		cache := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = cache
		//命中缓存
		if method == "GET" && c.Writer.Status() == 200 {
			cacheData, err := redisUtils.Redis{}.GetValue(path)
			if err != nil {
				logrus.Error("获取redis缓存失败:", err)
				c.Next()
				if cache.body != nil {
					errOver := redisUtils.Redis{}.SetValue(path, cache.body.String(), 5*time.Minute)
					if errOver != nil {
						logrus.Error("更新redis缓存失败:", err)
					}
				}
				return
			}
			if cacheData != "" { //命中
				var jsonData interface{}
				err := json.Unmarshal([]byte(cacheData), &jsonData)
				if err != nil {
					logrus.Error("缓存转化json失败:", err)
					c.Next()
					return
				}
				c.JSON(http.StatusOK, jsonData)
				c.Abort()
				return
			}
		}
		c.Next()
		err := redisUtils.Redis{}.DeletedValue(path)
		if err != nil {
			logrus.Error("删除redis缓存错误:", err)
			messageQueue.RabbitCache.PublishSimple(path) //消息队列补偿删除
			return
		}
	}
}

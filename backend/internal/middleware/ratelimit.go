package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// client 承载了一个访问者的令牌桶和最后访问时间
type client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// ipRateLimiter 存储所有客户端的访问记录
type ipRateLimiter struct {
	mu      sync.Mutex
	clients map[string]*client
	rate    rate.Limit
	burst   int
}

// NewIPRateLimiter 返回一个新的按 IP 维度的限流器
func NewIPRateLimiter(r rate.Limit, b int) *ipRateLimiter {
	limiter := &ipRateLimiter{
		clients: make(map[string]*client),
		rate:    r,
		burst:   b,
	}

	go func() {
		for {
			time.Sleep(time.Minute)
			limiter.cleanupClients()
		}
	}()

	return limiter
}

// getLimiter 为特定 IP 获取或创建 Limiter
func (i *ipRateLimiter) getLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	v, exists := i.clients[ip]
	if !exists {
		limiter := rate.NewLimiter(i.rate, i.burst)
		i.clients[ip] = &client{limiter: limiter, lastSeen: time.Now()}
		return limiter
	}

	v.lastSeen = time.Now()
	return v.limiter
}

// cleanupClients 清除长时间没有产生新请求的记录
func (i *ipRateLimiter) cleanupClients() {
	i.mu.Lock()
	defer i.mu.Unlock()
	for ip, v := range i.clients {
		if time.Since(v.lastSeen) > 3*time.Minute {
			delete(i.clients, ip)
		}
	}
}

// RateLimitMiddleware Gin 中间件：强制执行请求速率限制
func RateLimitMiddleware(limiter *ipRateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取真实的客户端 IP
		ip := c.ClientIP()

		l := limiter.getLimiter(ip)

		// 消费一个令牌，如果没有可用的令牌，则拦截并返回 429
		if !l.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"code":    429,
				"error":   "Too Many Requests",
				"message": "您的请求太频繁，请稍后再试",
			})
			return
		}

		c.Next()
	}
}

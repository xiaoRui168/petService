package main

import (
	"log"

	"petService/backend/internal/config"
	"petService/backend/internal/controller"
	"petService/backend/internal/middleware"
	"petService/backend/internal/pkg"
	"petService/backend/internal/repository"
	"petService/backend/internal/service"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	// Added for rate limiting
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 初始化数据库
	db := pkg.InitDB(cfg)

	// 自动迁移 & 种子数据
	pkg.Seed(db)

	// 初始化各层
	repo := repository.New(db)
	svc := service.New(repo)
	ctrl := controller.New(svc)

	// 创建 Gin 引擎
	r := gin.Default()

	// 中间件
	r.Use(middleware.CORS())

	// 创建受速率限制的专属速率限制器
	// 设定每一个 IP 每秒允许生成 1 个访问令牌，最多累积（并发爆发）3个请求，超出的直接返回 429
	appointmentLimiter := middleware.NewIPRateLimiter(rate.Limit(1), 3)

	// 注册 API 路由
	api := r.Group("/api/v1")
	ctrl.RegisterRoutes(api)
	// 单独注册带有限流中间件的 POST 预约端点
	api.POST("/appointments", middleware.RateLimitMiddleware(appointmentLimiter), ctrl.CreateAppointment)

	// 启动服务
	addr := ":" + cfg.ServerPort
	log.Printf("服务启动在 %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}

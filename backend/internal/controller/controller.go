package controller

import (
	"net/http"
	"strconv"

	"petService/backend/internal/model"
	"petService/backend/internal/service"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	svc *service.Service
}

func New(svc *service.Service) *Controller {
	return &Controller{svc: svc}
}

func (c *Controller) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/stylings", c.GetStylings)
	r.GET("/stylings/:id", c.GetStylingByID)
	r.GET("/packages", c.GetPackages)
	r.GET("/packages/:id", c.GetPackageByID)
	r.GET("/photos", c.GetPhotos)
	r.GET("/reviews", c.GetReviews)

	// Default generic appointment without custom middleware.
	// The explicit route with rate limiter will be registered in main.go
}

// ok 返回成功响应
func ok(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, model.Response{Code: 200, Message: "success", Data: data})
}

// fail 返回错误响应
func fail(ctx *gin.Context, httpCode int, code int, msg string) {
	ctx.JSON(httpCode, model.Response{Code: code, Message: msg, Data: nil})
}

// GetStylings 获取所有造型
func (c *Controller) GetStylings(ctx *gin.Context) {
	items, err := c.svc.GetAllStylings()
	if err != nil {
		fail(ctx, http.StatusInternalServerError, 500, "获取造型列表失败")
		return
	}
	ok(ctx, items)
}

// GetStylingByID 根据 ID 获取造型
func (c *Controller) GetStylingByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fail(ctx, http.StatusBadRequest, 400, "无效的 ID")
		return
	}
	item, err := c.svc.GetStylingByID(uint(id))
	if err != nil {
		fail(ctx, http.StatusNotFound, 404, "造型不存在")
		return
	}
	ok(ctx, item)
}

// GetPackages 获取所有套餐
func (c *Controller) GetPackages(ctx *gin.Context) {
	items, err := c.svc.GetAllPackages()
	if err != nil {
		fail(ctx, http.StatusInternalServerError, 500, "获取套餐列表失败")
		return
	}
	ok(ctx, items)
}

// GetPackageByID 根据 ID 获取套餐
func (c *Controller) GetPackageByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fail(ctx, http.StatusBadRequest, 400, "无效的 ID")
		return
	}
	item, err := c.svc.GetPackageByID(uint(id))
	if err != nil {
		fail(ctx, http.StatusNotFound, 404, "套餐不存在")
		return
	}
	ok(ctx, item)
}

// GetPhotos 获取所有照片
func (c *Controller) GetPhotos(ctx *gin.Context) {
	items, err := c.svc.GetAllPhotos()
	if err != nil {
		fail(ctx, http.StatusInternalServerError, 500, "获取照片列表失败")
		return
	}
	ok(ctx, items)
}

// GetReviews 获取所有评价
func (c *Controller) GetReviews(ctx *gin.Context) {
	items, err := c.svc.GetAllReviews()
	if err != nil {
		fail(ctx, http.StatusInternalServerError, 500, "获取评价列表失败")
		return
	}
	ok(ctx, items)
}

// CreateAppointment 创建预约
func (c *Controller) CreateAppointment(ctx *gin.Context) {
	var req model.AppointmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fail(ctx, http.StatusBadRequest, 400, "请求参数错误: "+err.Error())
		return
	}
	// 简单手机号格式校验
	if len(req.Phone) < 11 {
		fail(ctx, http.StatusBadRequest, 400, "手机号格式不正确")
		return
	}
	if err := c.svc.CreateAppointment(&req); err != nil {
		fail(ctx, http.StatusInternalServerError, 500, "预约失败: "+err.Error())
		return
	}
	ok(ctx, gin.H{"message": "预约成功"})
}

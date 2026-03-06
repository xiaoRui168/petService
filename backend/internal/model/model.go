package model

import "time"

// Styling 宠物美容造型
type Styling struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"size:100;not null"`
	Description string    `json:"description" gorm:"size:500"`
	ImageURL    string    `json:"image_url" gorm:"size:255"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Package 服务套餐
type Package struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"size:100;not null"`
	Description string    `json:"description" gorm:"size:500"`
	Services    string    `json:"services" gorm:"size:1000"` // 逗号分隔的服务项
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Photo 宠物照片
type Photo struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	ImageURL    string    `json:"image_url" gorm:"size:255;not null"`
	PetName     string    `json:"pet_name" gorm:"size:50"`
	Description string    `json:"description" gorm:"size:200"`
	Groomer     string    `json:"groomer" gorm:"size:50"`
	CreatedAt   time.Time `json:"created_at"`
}

// Review 客户评价
type Review struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Nickname  string    `json:"nickname" gorm:"size:50;not null"`
	PetType   string    `json:"pet_type" gorm:"size:50"`
	Rating    int       `json:"rating" gorm:"default:5"`
	Content   string    `json:"content" gorm:"size:1000;not null"`
	CreatedAt time.Time `json:"created_at"`
}

// Appointment 预约
type Appointment struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	OwnerName   string    `json:"owner_name" gorm:"size:50;not null"`
	Phone       string    `json:"phone" gorm:"size:20;not null"`
	PetType     string    `json:"pet_type" gorm:"size:50;not null"`
	PetName     string    `json:"pet_name" gorm:"size:50"`
	Date        string    `json:"date" gorm:"size:20;not null"`
	Time        string    `json:"time" gorm:"size:20;not null"`
	ServiceID   uint      `json:"service_id"`
	ServiceType string    `json:"service_type" gorm:"size:20"` // styling 或 package
	StylingID   string    `json:"styling_id" gorm:"size:20"`   // 新增：预约设定的造型ID
	PackageID   string    `json:"package_id" gorm:"size:20"`   // 新增：预约设定的套餐ID
	Remark      string    `json:"remark" gorm:"size:500"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// AppointmentRequest 预约请求
type AppointmentRequest struct {
	OwnerName   string `json:"owner_name" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	PetType     string `json:"pet_type" binding:"required"`
	PetName     string `json:"pet_name"`
	Date        string `json:"date" binding:"required"`
	Time        string `json:"time" binding:"required"`
	ServiceID   uint   `json:"service_id"`
	ServiceType string `json:"service_type"`
	StylingID   string `json:"styling_id"`
	PackageID   string `json:"package_id"`
	Remark      string `json:"remark"`
}

// Response 统一响应格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

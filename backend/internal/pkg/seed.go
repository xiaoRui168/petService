package pkg

import (
	"log"
	"time"

	"petService/backend/internal/model"

	"gorm.io/gorm"
)

// Seed 初始化种子数据
func Seed(db *gorm.DB) {
	// AutoMigrate
	if err := db.AutoMigrate(
		&model.Styling{},
		&model.Package{},
		&model.Photo{},
		&model.Review{},
		&model.Appointment{},
	); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	// 检查是否已有数据
	var count int64
	db.Model(&model.Styling{}).Count(&count)
	if count > 0 {
		log.Println("种子数据已存在，跳过初始化")
		return
	}

	log.Println("正在插入种子数据...")

	// 造型数据
	stylings := []model.Styling{
		{Name: "泰迪装", Description: "经典泰迪造型，圆润可爱，适合贵宾犬和比熊等品种", ImageURL: "https://images.unsplash.com/photo-1587300003388-59208cc962cb?w=400&h=300&fit=crop", Price: 168},
		{Name: "日系萌宠装", Description: "日式风格修剪，清新自然，让宝贝更加萌萌哒", ImageURL: "https://images.unsplash.com/photo-1583511655857-d19b40a7a54e?w=400&h=300&fit=crop", Price: 198},
		{Name: "狮子装", Description: "帅气狮子造型，威风凛凛又不失可爱", ImageURL: "https://images.unsplash.com/photo-1588943211346-0908a1fb0b01?w=400&h=300&fit=crop", Price: 218},
		{Name: "韩系精灵装", Description: "韩式精灵风格，优雅时尚，适合各种小型犬", ImageURL: "https://images.unsplash.com/photo-1543466835-00a7907e9de1?w=400&h=300&fit=crop", Price: 238},
		{Name: "贵族装", Description: "高贵典雅的贵族风格造型，彰显宝贝独特气质", ImageURL: "https://images.unsplash.com/photo-1583512603805-3cc6b41f3edb?w=400&h=300&fit=crop", Price: 268},
		{Name: "清爽短毛装", Description: "夏日清爽短毛造型，凉快舒适又好打理", ImageURL: "https://images.unsplash.com/photo-1552053831-71594a27632d?w=400&h=300&fit=crop", Price: 128},
	}
	db.Create(&stylings)

	// 套餐数据
	packages := []model.Package{
		{Name: "基础洗护套餐", Description: "适合日常清洁护理", Services: "洗澡,吹干,梳毛,耳朵清洁,指甲修剪", Price: 98},
		{Name: "精致美容套餐", Description: "全面美容护理，让宝贝焕然一新", Services: "洗澡,吹干,造型修剪,耳朵清洁,指甲修剪,齿部清洁,眼部护理", Price: 198},
		{Name: "豪华 SPA 套餐", Description: "顶级 SPA 享受，全方位宠爱您的毛孩子", Services: "深层清洁浴,精油 SPA,造型修剪,耳朵清洁,指甲修剪,齿部清洁,眼部护理,皮毛护理,香薰按摩", Price: 368},
		{Name: "猫咪专属套餐", Description: "专为猫咪设计的温柔护理方案", Services: "温水洗浴,轻柔吹干,梳毛去结,耳朵清洁,指甲修剪", Price: 128},
	}
	db.Create(&packages)

	// 照片数据
	photos := []model.Photo{
		{ImageURL: "https://images.unsplash.com/photo-1596492784531-6e6eb5ea9993?w=400&h=400&fit=crop", PetName: "豆豆", Description: "美容后的泰迪装", Groomer: "小美"},
		{ImageURL: "https://images.unsplash.com/photo-1583337130417-3346a1be7dee?w=400&h=400&fit=crop", PetName: "旺财", Description: "帅气的狮子造型", Groomer: "阿杰"},
		{ImageURL: "https://images.unsplash.com/photo-1548199973-03cce0bbc87b?w=400&h=400&fit=crop", PetName: "球球", Description: "清爽夏日短毛装", Groomer: "小美"},
		{ImageURL: "https://images.unsplash.com/photo-1514888286974-6c03e2ca1dba?w=400&h=400&fit=crop", PetName: "咪咪", Description: "猫咪 SPA 后", Groomer: "小花"},
		{ImageURL: "https://images.unsplash.com/photo-1533738363-b7f9aef128ce?w=400&h=400&fit=crop", PetName: "乐乐", Description: "日系萌宠装超可爱", Groomer: "阿杰"},
		{ImageURL: "https://images.unsplash.com/photo-1530281700549-e82e7bf110d6?w=400&h=400&fit=crop", PetName: "大福", Description: "韩系精灵装出浴", Groomer: "小美"},
		{ImageURL: "https://images.unsplash.com/photo-1517423440428-a5a00ad493e8?w=400&h=400&fit=crop", PetName: "小白", Description: "贵族装出品", Groomer: "小花"},
		{ImageURL: "https://images.unsplash.com/photo-1518020382113-a7e8fc38eac9?w=400&h=400&fit=crop", PetName: "毛毛", Description: "洗完澡好开心", Groomer: "阿杰"},
	}
	db.Create(&photos)

	// 评价数据
	reviews := []model.Review{
		{Nickname: "萌宠妈妈", PetType: "贵宾犬", Rating: 5, Content: "第一次带豆豆来做美容，效果超棒！美容师很耐心，豆豆全程都很配合，做完泰迪装后简直萌翻了！", CreatedAt: time.Now().AddDate(0, 0, -30)},
		{Nickname: "汪星人爸爸", PetType: "金毛", Rating: 5, Content: "豪华 SPA 套餐物超所值！旺财洗完后毛发特别顺滑，回家后一直在沙发上撒娇，看得出来很享受！", CreatedAt: time.Now().AddDate(0, 0, -25)},
		{Nickname: "猫奴小王", PetType: "英短", Rating: 4, Content: "猫咪专属套餐很贴心，咪咪平时很怕水，但美容师的手法非常轻柔，全程没有太大的应激反应，下次还来！", CreatedAt: time.Now().AddDate(0, 0, -20)},
		{Nickname: "宠爱有加", PetType: "比熊", Rating: 5, Content: "朋友推荐来的，果然没让我失望！球球的造型做得特别精致，拍了好多照片发朋友圈，大家都在问是哪里做的！", CreatedAt: time.Now().AddDate(0, 0, -15)},
		{Nickname: "毛孩子的家", PetType: "柯基", Rating: 5, Content: "基础洗护套价格很实惠，洗得很干净，耳朵和指甲都处理得特别好。小短腿出来后像换了一只狗哈哈！", CreatedAt: time.Now().AddDate(0, 0, -10)},
		{Nickname: "爱猫人士", PetType: "布偶猫", Rating: 4, Content: "带我家小公主来做了精致美容套餐，梳毛很仔细，吹干后毛发蓬松柔软。就是猫咪有点认生，美容师安抚得很好。", CreatedAt: time.Now().AddDate(0, 0, -5)},
	}
	db.Create(&reviews)

	log.Println("种子数据插入完成")
}

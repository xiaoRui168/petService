# 萌宠美容工坊 (Pet Service)

## 1. 项目介绍
“萌宠美容工坊”是一个面向宠物主人的美容服务宣传与在线预约系统。项目提供了活泼温暖的 UI 界面，包含宠物造型展示、服务套餐说明、美容画廊、客户评价以及在线预约表单功能。

前后端分离架构设计，前端应用响应式布局及动态交互，后端提供稳定的 RESTful API 及数据持久化支持。

## 2. 技术栈
### 前端技术栈
- **框架**：Vue 3 (Composition API)
- **构建工具**：Vite 6
- **请求客户端**：Axios
- **样式预处理器**：SCSS / Sass
- **设计风格**：活泼温暖配色系统（原生 CSS 编写，未过度依赖第三方 UI 库）

### 后端技术栈
- **语言**：Go (1.25+)
- **Web 框架**：Gin (轻量级、高性能 RESTful 框架)
- **ORM 框架**：GORM
- **数据库**：MySQL 8
- **跨域中间件**：gin-contrib/cors

## 3. 项目结构说明

```text
petService/
├── frontend/               # 前端工程（Vue.js + Vite）
│   ├── public/             
│   ├── src/
│   │   ├── api/            # Axios 接口请求封装
│   │   ├── assets/         # 全局 SCSS 样式及静态资源
│   │   ├── components/     # 大量 Vue 拆分组件（NavBar, HeroBanner, StylingSection 等）
│   │   ├── App.vue         # 根前端组件
│   │   └── main.js         # 前端入口
│   ├── index.html          # HTML 模板
│   ├── package.json        # 前端依赖配置
│   └── vite.config.js      # Vite 配置及代理设置
├── backend/                # 后端工程（Go + Gin）
│   ├── internal/           # 私有业务代码
│   │   ├── config/         # 环境变量与配置读取
│   │   ├── controller/     # 路由控制与入参校验层
│   │   ├── middleware/     # 中间件（CORS 等）
│   │   ├── model/          # 数据库实体结构定义
│   │   ├── pkg/            # 数据库连接初始化及种子数据注入
│   │   ├── repository/     # 数据库访问层 (CRUD)
│   │   └── service/        # 核心业务逻辑层
│   ├── go.mod              # Go 依赖配置
│   └── main.go             # 后端应用启动入口
└── README.md               # 项目说明文档
```

## 4. API 接口规范说明

后端接口均遵循 RESTful 规范，基础路径为 `/api/v1`。
目前主要提供的无状态公共接口：

- **GET `/api/v1/stylings`**：获取所有门店推荐宠物造型及预估价格数据。
- **GET `/api/v1/packages`**：获取洗护套餐服务详情表。
- **GET `/api/v1/photos`**：获取宠物洗护美容后的展示画廊照片。
- **GET `/api/v1/reviews`**：获取顾客对商家和美容师的历史评价。
- **POST `/api/v1/appointments`**：提交宠物主人的在线预约请求（需包含主人口令、电话、爱宠信息及期望时间）。

所有接口成功时统一返回格式：
```json
{
  "code": 200,
  "message": "success",
  "data": { ... }
}
```

## 5. 测试（开发）环境的启动与关闭

项目开发环境中，前后端服务独立运行，前端通过 Vite 的 proxy 机制解决跨域问题。

### 5.1 启动前提
- 已安装 Node.js (v18+) 及 npm
- 已安装 Go (v1.25+)
- 运行中的 MySQL 数据库

### 5.2 启动后端 (Go)
1. 配置好本地 MySQL 服务并创建数据库 `pet_service`（如果尚未创建）。
2. 打开终端，进入后端目录：
   ```bash
   cd petService/backend
   ```
3. *(注意)* 请确保 `backend/internal/config/config.go` 中的数据库连接信息 (`DB_HOST`, `DB_USER`, `DB_PASSWORD`) 与你的本地环境一致。
4. 运行服务：
   ```bash
   go run main.go
   ```
   *服务成功启动后会监听在 `8080` 端口，并在首次启动时自动执行 AutoMigrate 数据库建表与注入示范种子数据。*

### 5.3 启动前端 (Vue)
1. 打开新的终端，进入前端目录：
   ```bash
   cd petService/frontend
   ```
2. 安装依赖包：
   ```bash
   npm install
   ```
3. 启动开发服务器：
   ```bash
   npm run dev
   ```
   *服务成功启动后，使用浏览器访问 `http://localhost:3000` (或局域网 IP `http://<YOUR_IP>:3000`) 即可预览。*

### 5.4 关闭测试环境服务
在运行前端或后端的终端窗口中，按下 `Ctrl + C` 即可停止对应的进程。
如果存在后台挂起、端口占用的情况，可以手动结束进程：
```bash
# MacOS / Linux
lsof -i :3000 -i :8080 | awk 'NR>1 {print $2}' | sort -u | xargs kill -9
```

## 6. 正式（生产）环境部署指南

在生产环境中，出于性能和安全性考虑，建议使用 Nginx 反向代理与静态资源托管，配合 Supervisor 守护后端进程 或 使用 Docker 容器化部署。

### 6.1 前端打包部署
前端由于是静态应用，仅需编译后放入 Web 服务器（如 Nginx）。
1. 编译打包：
   ```bash
   cd petService/frontend
   npm run build
   ```
2. 打包完成后，同级目录下将生成 `dist` 文件夹，将其拷贝到服务器指定的静态资源挂载目录（如 `/var/www/pet-service-frontend`）。

### 6.2 后端编译部署
1. 编译可执行文件：
   ```bash
   cd petService/backend
   CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pet-service-api main.go
   ```
   *(根据具体服务器系统架构修改 `GOOS` 和 `GOARCH`)*
2. 将生成的 `pet-service-api` 放到生产服务器。可以使用 `systemd` 守护进程后台运行它。

### 6.3 Nginx 配置示例
配置 Nginx 同屏代理前端的单页应用以及后端的 API。
```nginx
server {
    listen 80;
    server_name yourdomain.com;

    # 1. 代理前端静态资源
    location / {
        root /var/www/pet-service-frontend;
        index index.html;
        try_files $uri $uri/ /index.html;
    }

    # 2. 代理后端 API 接口
    location /api/ {
        proxy_pass http://127.0.0.1:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_addrs;
    }
}
```
配置完成后重启 Nginx 即可通过 `yourdomain.com` 访问站点并调用接口。

### 6.4 推荐：零成本全栈部署方案 (Docker + PaaS 托管平台)

如果你希望快速上线并且**不花费任何服务器费用**，非常建议使用现代化的容器云原生部署架构。以下是业界最佳实践的 0 成本组合方案：

#### 方案架构
- **前端静态托管**：[Vercel](https://vercel.com) / [Netlify](https://www.netlify.com) (全自动构建、免费全球 CDN 与 HTTPS 域名)
- **后端服务托管**：[Render](https://render.com) (首选，支持免费关联 GitHub 部署自动构建，无需绑卡验证)
- **生产级云数据库**：[Supabase](https://supabase.com) (提供免费额度的 Serverless PostgreSQL，被称为开源 Firebase)

#### 实操步骤
1. **获取免费云数据库**
   - 访问 [Supabase 官网](https://supabase.com) 注册并创建一个新项目，获取它分配给你的 PostgreSQL 数据库连接字符串（DSN）。
   - 修改本项目后端代码 `internal/config/config.go`，配置 GORM 使用 Postgres 驱动，并将账号密码绑定。

2. **容器化后端服务 (Go API)**
   - 在 `backend/` 目录下创建一个无后缀名为 `Dockerfile` 的文件：
   ```dockerfile
   # Stage 1: Build
   FROM golang:alpine AS builder
   WORKDIR /app
   COPY go.mod go.sum ./
   RUN go mod download
   COPY . .
   # CGO_ENABLED=0 编译脱离环境依赖的静态程序
   RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pet_server main.go

   # Stage 2: Run
   FROM alpine:latest
   WORKDIR /app
   COPY --from=builder /app/pet_server .
   EXPOSE 8080
   CMD ["./pet_server"]
   ```
   - 将整个 `backend` 目录推送到你个人的 GitHub 仓库。
   - 登录 [Render 控制台](https://dashboard.render.com)，选择新建 Web Service，连接你的 GitHub 仓库。填入环境变量（如数据库连接密码等），Render 会拉取代码、自动构建 Docker 镜像并运行，并送你一个免费带有 SSL 的接口域名（例如 `https://pet-api.onrender.com`）。

3. **前端代码修改与一键部署 (Vue 3)**
   - 在前端代码 `frontend/src/api/index.js` 中，把原本的本地测试 baseUrl `/api/v1` 修改为刚刚 Render 分配给你在线接口的绝对地址：`https://pet-api.onrender.com/api/v1`。
   - 将前端代码推送到另一个独立的 GitHub 仓库。
   - 登录 [Vercel](https://vercel.com) 控制台，导入这个前端仓库，框架预设选择 `Vite`。点击 Deploy，1 分钟内部署完成。

**至此，前端和后端的数据全链路全部打通，你获得了一套架构极其现代、自动集成 SSL 并且每月 0 成本费用的线上生产版本系统！**


## 7. 数据与安全性建议

目前为了方便展示，前后端接口是开放授权的。如果准备将其作为线上商用工具，强力建议做如下升级：
1. **防止恶意请求遍历**：为 POST 预约提交接口增加速率限制（Rate Limit）或者图形验证码/Cloudflare Turnstile 防治灌水机器人。
2. **鉴权升级**：若需实现店家后台管理端，需结合 JWT (JSON Web Token) 在 `middleware` 中封存 `authMiddleware()` 来保障后台数据安全。
3. **输入防注入**：当前表单基于 Gin Validator 进行基础参数格式化保障。若涉及富文本内容，应对输入源进行彻底的 XSS 清洗。

## 8. 未来拓展规划 (Roadmap)
由于工坊正处于快速成长期，以下是基于项目结构梳理出的未来扩展建议：
- [ ] **在线支付模块**：接入微信/支付宝小程序在线打通支付押金。
- [ ] **店家后台系统**：新增对应的 Vue 管理端应用，支持从图形界面实时控制“画廊图片”、“顾客评价”、“快速上下架套餐业务类型”。
- [ ] **美容师排班系统**：预约接口进一步绑定到美容师的可预约时段库存上，提供直观的日历组件供用户选择空当期。

# gin_jiaoshoujia

`gin_jiaoshoujia` 是一个基于 [Gin](https://github.com/gin-gonic/gin) 和 [GORM](https://gorm.io/) 的 Go 项目脚手架，支持 MySQL 和 SQLite 数据库，提供了模块化的路由管理和基础功能，方便快速构建 Web 应用。

## 功能特性

- **模块化路由**：支持分组管理路由。
- **多数据库支持**：可选择使用 MySQL 或 SQLite。
- **自动化迁移**：基于 GORM 的模型自动迁移功能。
- **示例模块**：包含用户模块示例（用户创建和查询）。

## 项目结构

```
gin_jiaoshoujia/
├── config/          # 配置文件
│   └── config.go    # 配置加载逻辑
├── controllers/     # 控制器
│   └── user_controller.go # 用户控制器
├── models/          # 数据模型
│   ├── db.go        # 数据库初始化
│   └── user.go      # 用户模型
├── routers/         # 路由管理
│   └── router.go    # 路由初始化
├── services/        # 业务逻辑 (暂未实现)
├── utils/           # 工具函数 (暂未实现)
├── main.go          # 项目入口
└── go.mod           # Go 模块配置
```

## 快速开始

### 1. 克隆项目
```bash
git clone https://github.com/KearChen/gin_jiaoshoujia.git
cd gin_jiaoshoujia
```

### 2. 安装依赖
```bash
go mod tidy
```

### 3. 配置文件
在 `config/config.yaml` 中配置数据库和应用信息：
```yaml
app:
  name: gin_jiaoshoujia
  port: 8080

database:
  type: mysql # 可选 sqlite 或 mysql
  mysql:
    host: 127.0.0.1
    port: 3306
    user: root
    password: root
    dbname: gin_db
  sqlite:
    path: gin_jiaoshoujia.db
```

### 4. 启动项目
```bash
go run main.go
```

访问以下接口测试功能：
- `GET /ping` - 健康检查
- `POST /users/` - 创建用户
- `GET /users/` - 获取用户列表

### 示例请求

#### 创建用户
```bash
curl -X POST http://localhost:8080/users/ \
-H "Content-Type: application/json" \
-d '{"name": "John Doe", "email": "john@example.com", "age": 30}'
```

#### 获取用户列表
```bash
curl http://localhost:8080/users/
```

## 扩展开发

1. **添加新模块**：在 `models/`、`controllers/` 和 `routers/` 中分别定义模型、控制器和路由。
2. **业务逻辑**：将复杂的逻辑拆分到 `services/` 中，保持代码清晰。
3. **中间件**：可在 `routers/router.go` 中注册自定义中间件（如鉴权、日志等）。

## 依赖

- [Gin](https://github.com/gin-gonic/gin) - 高性能 Web 框架
- [GORM](https://gorm.io/) - 强大的 ORM 工具
- [Viper](https://github.com/spf13/viper) - 配置管理工具


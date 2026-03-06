# Comic Go - 前后端分离版本

基于 Go + React 的漫画浏览器，前后端分离架构。

## 项目结构

```
comic-go/
├── backend/          # Go 后端
│   ├── main.go       # 入口文件
│   ├── config/       # 配置
│   ├── models/       # 数据模型
│   ├── handlers/     # API 处理器
│   └── middleware/   # 中间件
└── frontend/         # React 前端
    ├── src/
    │   ├── pages/    # 页面组件
    │   ├── components/ # 通用组件
    │   └── api.js    # API 客户端
    └── package.json
```

## 后端 API

### 已实现接口

#### 漫画相关
- `GET /api/comics` - 漫画列表（支持分页、排序、筛选）
  - 参数：`page`, `page_size`, `sort`, `tag_id`, `category_id`, `search`
- `GET /api/comics/:id` - 漫画详情
- `GET /api/comics/:id/images` - 漫画图片列表
- `GET /api/search` - 搜索漫画

#### 标签与分类
- `GET /api/tags` - 标签列表
- `GET /api/categories` - 分类列表

#### 图片服务
- `GET /api/images/:comic_id/:filename` - 图片代理服务

#### 用户相关（需认证）
- `POST /api/auth/register` - 用户注册
- `POST /api/auth/login` - 用户登录
- `POST /api/comics/:id/rate` - 评分
- `GET /api/user/favorites` - 收藏列表
- `POST /api/user/favorites/:id` - 添加收藏
- `DELETE /api/user/favorites/:id` - 取消收藏

### API 测试结果

✅ 所有接口测试通过
- 漫画列表：153758 条记录
- 漫画详情：包含标签、作者等完整信息
- 图片列表：正常返回图片 URL 和本地路径
- 搜索功能：支持标题、作者、标签搜索
- 标签列表：正常返回

## 快速开始

### 后端

```bash
cd backend
go run main.go
```

后端默认运行在 `http://localhost:8080`

### 前端

```bash
cd frontend
npm install
npm run dev
```

前端默认运行在 `http://localhost:5173`

## 数据库

使用现有的 SQLite 数据库：`/Users/huyaaaaaa/project/spider/data/comics.db`

### 表结构
- `comics` - 漫画主表（153758 条）
- `comic_images` - 漫画图片
- `comic_tags` - 漫画标签关联
- `comic_authors` - 漫画作者关联
- `tags` - 标签表

## 技术栈

### 后端
- **框架**：Gin
- **ORM**：GORM
- **数据库**：SQLite
- **认证**：JWT

### 前端
- **框架**：React 18
- **路由**：React Router
- **样式**：Tailwind CSS
- **构建**：Vite

## 开发计划

- [ ] 完善前端页面
- [ ] 添加图片懒加载
- [ ] 实现阅读器功能
- [ ] 添加下载管理
- [ ] 部署到生产环境

## License

MIT License

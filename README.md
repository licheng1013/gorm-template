# GormTemplate

## 介绍

- 本项目作为后端项目模板
- 致力于提供一个简单的后台项目模板，方便快速开发

## 支持技术

- [x] Go 
- [ ] Java
- [ ] Rust
- 后续可能会支持更多语言

### Go
- 配套Goland插件: https://plugins.jetbrains.com/plugin/20411-gorm/
- 用于生成后端/前端模板代码

## 前端项目

- 定制于此前端开源项目 https://github.com/vbenjs/vue-vben-admin (支持功能请通过链接查看)
- 细节改动:
- view/src/store/modules/user.ts
- view/src/utils/http/axios/index.ts
- view/types/axios.d.ts
- 常用配置:
- view/.env.development 配置开发环境
- view/.env.production 配置生成环境
- 登入请求:
- view/src/api/custom/admin.ts
- 增加页面和接口
- 接口 -> view/src/api/custom
- 页面 -> view/src/views/custom
- 路由 -> view/src/router/routes/modules


## 后端项目
- 基于gin+gorm打造的
- 使用库:
- 工具: https://github.com/duke-git/lancet 
- Jwt: https://github.com/golang-jwt/jwt/


## 快速上手

- 1.克隆本项目
- 2.创建数据库并导入sql -> t_gorm.sql, 修改 admin/app.yml 配置文件 mysql 链接
- 3.安装go依赖并运行admin/main.go
- 4.安装pnpm
- 5.进入view目录 运行pnpm install
- 6.进入view目录 pnpm run dev

### 后端关键部分

- common/middleware 中间件
- common/app 应用配置和功能
- common/component 一些简单组件
- common/model 数据库模型
- common/service 业务逻辑
- common/tool 工具类

### 扩展
- 暂无

## 帮助

- QQ群: 289132257
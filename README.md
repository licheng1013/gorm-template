# GormTemplate

## 介绍

- 本项目作为Gorm插件整合后端项目模板
- 致力于提供一个简单的后台项目模板，方便快速开发
- 配套Goland插件: https://plugins.jetbrains.com/plugin/20411-gorm/
- 用于生成后端/前端模板代码

## 前端项目

- 定制于此前端开源项目 https://github.com/vbenjs/vue-vben-admin
- 细节改动:
- view/src/store/modules/user.ts 
- view/src/utils/http/axios/index.ts
- view/types/axios.d.ts
- 常用配置:
- view/.env.development 配置开发环境
- view/.env.production  配置生成环境
- 登入请求:
- view/src/api/custom/admin.ts
- 接口页面存放目录
- 接口 -> view/src/api/custom
- 页面 -> view/src/views/custom


## 快速上手

- 1.克隆本项目
- 2.创建数据库并导入sql -> t_gorm.sql, 修改 admin/app.yml 配置文件 mysql 链接
- 3.运行admin/main.go

### 后端关键部分

- common/middleware 中间件
- common/app 应用配置和功能
- common/component 一些简单组件
- common/model 数据库模型
- common/service 业务逻辑
- common/tool 工具类

## 帮助

- QQ群: 289132257
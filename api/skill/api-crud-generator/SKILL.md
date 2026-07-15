---
name: api-crud-generator
description: 根据Prisma模型生成标准的go API route + 前端管理页面.
version: 1.0.0
category: Backend Development
platforms: [Cursor, Claude Code, WorkBuddy]
license: MIT
---

## 功能说明
根据制定的Prismatic模型，自动生成标准的管理后台CRUD代码
- API Routes（5个），列表、详情、创建、更新、删除，都使用POST
- 前段页面，数据列表页、创建、编辑、删除

## 执行步骤

### 第1步：确认模型信息
询问用户：
- 要生成的模型名称（如：student、teacher、course等）
- API路径（如：/course/personnel/student/list"）
- 页面路由（如 admin/page/index）
- 
### 第2步：生表结构
- 按照模块进行创建文明名
- 生成对于的sql语句，放在/Users/xiaohei/work/gowork/xiaohei/model/mysql下方
- 生成的sql语句可以执行，满足3范式
- 表结构需要确认后才可以执行下方的步骤。

### 第3步：生成API后段代码 Route Handlers logic model modelDao reqs resp
按照标准模版生成一下文件：
- routesCourse.go· - 列表、详情、创建、更新、删除。
- 根据现有的结构生成对于heandler logic model modelDao reqs resp的对应的代码

### 第4部： 生成前段的管理页面
生成一个包含以下功能的管理页面：
- 数据表格（列出主要的字段信息）可以根据传入的要求进行展示
- 新增按钮+表单
- 每行的编辑和删除按钮
- 样式保持当前的样式

## 注意事项
- 所有UI文案都使用中文
- 密码字段用管通过API进行返回

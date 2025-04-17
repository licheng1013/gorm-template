# Gorm

## 介绍

- 本项目是 `Gorm` 插件的示例项目
- 有问题可以加QQ群: `289132257`

### 生成代码

- 2024.1.1 +插件版本
- 1.打开Database插件进行连接数据库
- ![](image/img_3.png)
- 2.在右侧(一般情况)工具栏选择
- ![](image/img_4.png)
- 3.选择要生成的表进行右击
- ![](image/img_5.png)
- 4.选择生成的目录
- ![](image/img_6.png)
- 5.`支持模板生成可定制` 

### 接口导航

- 1.打开`Gorm`插件
- ![](image/router/img.png)
- 2.在右侧(一般情况)工具栏选择
- ![](image/router/img_1.png)
- 3.可以根据路由进行接口搜索(`注意需要按指定方式配置接口`)

### Goland控制台导航

- 2024.1.7+版本

- 1.运行效果
- ![](image/img_9.png)


### 结构体扩展

- 1.扩展功能,可将Go结构体转换为其他语言类
- ![](image/router/img_3.png)


## 设置

## Mvc.Json
- 文件配置一般在 .idea/gorm 或者 gorm 目录下
- 此文件有些设置需要注意: [Mvc.Json](gorm/Mvc.Json)
- 如果字段不存在可以手动添加进去即可
 
| 参数                 | 描述                                     | 插件版本要求    |
|--------------------|----------------------------------------|-----------|
| enableSqlNull      | 为 `true` 则开启sqlNull模式。 默认为`false`      |           |
| enableCommonResult | 为 `true` 则生成一个返回结果类(在模型目录下)。 默认为`true` | 2024.1.2+ |
| pathPrefix         | 控制台路径导航到接口设置,默认为""空, 解决全局前缀路径没法识别      | 2024.1.3+ |
| enableGormTag      | 为 `true` 则开启gorm tag的信息 。 默认为`false`   | 2024.1.5+ |


### enableGormTag

- 示例
- ![](image/img_7.png)

### sqlNull模式

- 字段改成
- ![1](image/img.png)

### 接口扩展

```json
{
  "extendTag": [
    "-dart",
    "-kt"
  ]
}
```
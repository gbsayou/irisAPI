# IrisAPI

---

## 项目介绍

+ 基于 [Iris](https://github.com/kataras/iris)
+ 使用 MySQL
+ orm 使用 [gorm](https://github.com/jinzhu/gorm)

## 实现的功能

+ 基本的数据库操作
+ 鉴权
+ 日志记录

## 目录结构

本项目的目录结构参考 Node.JS 项目的常用结构：

+ config：项目配置
+ controllers：路由处理controller
+ database： 提供数据库相关功能
+ logger：日志处理方法与日志
+ middleware：中间件
+ models：orm model
+ routers：路由
+ services：具体的业务处理方法，包括操作数据库的部分
+ main.go 入口

## 运行：

`go run main.go`

本项目涉及到 iris 源码的修改，具体请看我的另一个项目：[拦截返回信息](https://github.com/gbsayou/learnIris/blob/master/3.%E4%B8%AD%E9%97%B4%E4%BB%B6/3.2%E6%8B%A6%E6%88%AA%E8%BF%94%E5%9B%9E%E4%BF%A1%E6%81%AF/%E6%8B%A6%E6%88%AA%E8%BF%94%E5%9B%9E%E4%BF%A1%E6%81%AF.md)

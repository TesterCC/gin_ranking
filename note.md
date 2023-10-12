## Environment

go + gin + mysql5.7/8.0

## Notes

3-7 引入ORM框架Gorm并实现数据库操作 https://www.imooc.com/video/24561

3-6 如何自定义logger中间件并实现日志收集

主要3类日志需要收集：
1. 项目的请求日志（类似Nginx日志，每次请求都保留下来）
2. 错误日志
3. 开发业务时需要的自定义日志

安装第三方日志库:
```
go get github.com/sirupsen/logrus
```

3-7 引入ORM框架Gorm并实现数据库操作

gorm.io官网: `https://gorm.io/zh_CN/docs/connecting_to_the_database.html`

```
Gorm是Go的一个ORM（对象关系映射）库。支持包括：MySQL、PostgreSQL、SQLite等。
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

前三章基础demo在basic_learn分支中。

4-2 实现用户登录和注册功能以及session的使用

先清理架构旧代码，然后基于此框架做后续开发。

```
项目功能：
1.用户注册
2.用户登录
3.活动规则
4.参赛选手列表
5.参赛选手详情
6.排行榜
7.投票
```

## 基本开发流程

1.定义model
2.定义controller
3.定义router

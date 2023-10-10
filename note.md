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
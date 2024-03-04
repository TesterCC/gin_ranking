## Environment

go + gin + mysql5.7/8.0

## Build

交叉编译命令

```shell
GOOS=linux GOARCH=amd64 go build
```

会编译出linux可执行文件 gin_ranking
上传到服务器指定目录，执行：`chmod +x gin_ranking`
启动程序: `./gin_ranking`

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

启动session需要安装的三方包：

```
go get github.com/gin-contrib/sessions
```

另外注意，这里要sessions生效，需要先在服务器上安装配置redis

## CentOS7安装Redis6

```
wget https://download.redis.io/releases/redis-6.2.12.tar.gz
tar -zxvf redis-6.2.12.tar.gz
cd redis-6.2.12
make
make install
```

若无报错，则会在src目录里生成编译好的可执行文件若干，单还不能直接通过redis命令启动，可以使用ln -s软链接到CentOS7的/bin目录下

```shell
sudo ln -s /opt/source/redis-6.2.12/bin/redis-server /bin/redis-server
sudo ln -s /opt/source/redis-6.2.12/bin/redis-cli /bin/redis-cli
sudo ln -s /opt/source/redis-6.2.12/bin/redis-benchmark /bin/redis-benchmark
sudo ln -s /opt/source/redis-6.2.12/bin/redis-sentinel /bin/redis-sentinel
```

redis.conf中需要修改的内容

```
bind 0.0.0.0   # 方便调试，注意，对外开放后一定要设置密码
requirepass foobared    # 设置密码
daemonize no   # supervisor启动保持no（因为supervisor都是非daemon启动的），如果是手动启动可以设置为yes
```

redis-server通过supervisor启动的配置

```
[program:redis]
command=/opt/source/redis-6.2.12/bin/redis-server /opt/source/redis-6.2.12/bin/redis.conf
autostart=true
autorestart=true
startsecs=5
startretries=2000
```

## 基本开发流程

1.定义model 2.定义controller 3.定义router

## Go使用Redis

第三方包go-redis

```
https://redis.uptrace.dev/zh/
支持Redis Server和Redis Cluster的Golang客户端

go get "github.com/redis/go-redis/v9"
```

Redis GUI工具
支持windows环境
```
RedisInsight
https://redis.io/docs/connect/insight/
```

官方使用方式example

```
package main

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "", // no password set
		DB:		  0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

```
## 学习建议
1. 站在架构的层面思考问题
2. 学习设计模式 - [类的设计原则](https://www.imooc.com/article/335650)
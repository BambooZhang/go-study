# go-study


## go-base：基础学习

- o-1：基础类型
- go-2:信道:channel
- go-3:mysql的使用
- go-4:mysql的使用
- go-5：io使用
- go-6：模板template使用
- go-7：package和可见性
- go-8：反射
- go-9：excle读写操作(用到了一点反射根据属性获取实例的值)
- go-10: tcp服务器，简单聊天室
- go-11: JSON字符串解析和生成JSON字符串
- go-12: 国际化处理
- go-13: 国际化处理
- go-14:  web反射找到对应的方法


## go web
go-web:gobee



## 说明
author:zjcjava@163.com
time:2018-9-8




# GO项目说明



项目放置所在文件夹路径如下
CODE_PATH/src/当前项目

eg:D:/workerspaace/go/src




## 环境配置
0.   GOROOT,PATH配置好

1 . wind环境变量设置GOPATH,建议非系统盘,第三方依赖都会下载到该路径下,会耗费磁盘空间

GOPATH:D:\go

2.IDEA设置扩展PATH

IDEA->file(文件)->settings->Languages & Frameworks -> Go ->GORATH

加入自己的本地项目目录的CODE_PATH,它会自动CODE_PATH/src引入本地依赖的项目package


## 此项目依赖第三方包和项目


ssh://git@120.78.183.122:59188/srv/git/commongo.git
ssh://git@120.78.183.122:59188/srv/git/kfgo.git
ssh://git@120.78.183.122:59188/srv/git/rdgo.git


```
go get gopkg.in/yaml.v2
go get github.com/Shopify/sarama
go get github.com/beego/goyaml2
go get -u github.com/kataras/iris
go get -u github.com/jmoiron/sqlx
go get github.com/pkg/errors
go get github.com/sirupsen/logrus
go get github.com/robfig/cron
go get github.com/rifflock/lfshook
go get github.com/patrickmn/go-cache
go get github.com/lestrrat-go/file-rotatelogs

go get github.com/beego/goyaml2
go get github.com/bsm/sarama-cluster
go get github.com/garyburd/redigo/redis
go get github.com/smallnest/rpcx/client
go get github.com/smallnest/rpcx/server
go get github.com/smallnest/rpcx/serverplugin
go get github.com/docker/libkv
go get github.com/samuel/go-zookeeper/zk
go get github.com/shopspring/decimal

go  get golang.org/x/net 替代安装方法
在github.com路径下创建x/net目录后进入执行如下操作
git clone https://github.com/golang/net net
go install net
git clone https://github.com/golang/text text


rpc错误处理
# common/rpc
..\common\rpc\client.go:27:9: undefined: client.NewZookeeperDiscovery
..\common\rpc\server.go:21:8: undefined: serverplugin.ZooKeeperRegisterPlugin
打开如下文件，把第一行的// +build zookeeper删除保存即可
NewZookeeperDiscovery.go
ZooKeeperRegisterPlugin.go

```


exec: "gcc": executable file not found in %PATH%
错误处理
解压如下文件，到本机，并把解压后的路径 WGW-PATH/mingw64/bin
x86_64-8.1.0-release-posix-sjlj-rt_v6-rev0
1.系统环境变量PATH中加入WGW-PATH/mingw64/bin
2.IDEA设置扩展PATH:WGW-PATH/mingw64/bin添加扩展path中（参考上文中的->idea设置扩展PATH）




## windows下编译coin,部署至CENTOS测试环境方式
1、通过cmd进入项目代码存放的工作空间
2、执行如下命令：
    set GOOS=linux
    set GOARCH=amd64
3、编译打包coin项目
    go build match.engine
4、执行完编译后,再coin项目工作目录下会生成coin文件，大小月20M左右(根据实际情况来...)
5、登陆linux测试服务器, 进入路径： /data/webdata/apps/match/， 然后通过ftp工具上传match文件，并替换测试环境服目录下的match文件,
   注意：因权限问题此处不能直接上传match文件，需将要上传的文件先改个名再上传, 备份好原来的文件， 例如：coin_1,
        然后将新上传的coin_1文件重新命名为coin  
6、进入linux命令行, 查找进程，并执行kill命令, 注意此处不能带 -9，直接使用kill 进程号   
    查找match进程： ps -ef|grep match.engine
    kill进程: kill 92368
7、因权限问题无法直接启动， 需先授权： chmod 777 coin
8、进入目录: cd /data/webdata/apps/match 执行启动命令： startmatch
9、进入日志目录cd apps/logs  观察是否启动异常    



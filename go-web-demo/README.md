


```



go get github.com/btcsuite/btcd/btcjson

go get github.com/btcsuite/btcutil
go get github.com/tealeg/xlsx


go get github.com/btcsuite/btclog
go get github.com/btcsuite/go-socks/socks
go get github.com/btcsuite/websocket
github.com/btcsuite/go-socks/socks


```

github.com/kavu/go_reuseport
golang.org/x/sys/unix

go get golang.org/x/crypto/ripemd160 替代安装方法
在github.com路径下创建x/net目录后进入执行如下操作
git clone https://github.com/golang/crypto crypto
go install crypto

windows下编译coin.merchant,部署至测试环境方式
1、通过cmd进入项目代码存放的工作空间
2、执行如下命令：
    set GOOS=linux
    set GOARCH=amd64
3、编译打包coin.merchant项目
    go build coin.merchant
4、执行完编译后,在coin.merchant项目工作目录下会生成coin.merchant文件，大小月20M左右(根据实际情况来...)
5、登陆linux测试服务器, 进入路径： /data/webdata/apps/coin.merchant/， 然后通过ftp工具上传coin.merchant文件，并替换测试环境服目录下的coin.merchant文件,
   注意：因权限问题此处不能直接上传coin.merchant文件，需将要上传的文件先改个名再上传, 备份好原来的coin.merchant文件， 例如：coin.merchant_1,
        然后将新上传的coin.merchant_1文件重新命名为coin.merchant  
6、进入linux命令行, 查找coin.merchant进程，并执行kill命令, 注意此处不能带 -9，直接使用kill 进程号   
    查找coin.merchant进程： ps -ef|grep coin.merchant
    kill进程: kill 92368
7、因权限问题无法直接启动， 需先授权： chmod 777 coin.merchant
8、进入目录: cd /data/webdata/apps/coin.merchant 执行启动命令： startmerchant,  
9、进入日志目录cd apps/logs  观察是否启动异常    
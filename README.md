## go 

### 环境安装

### 打包部署

1.

``` 
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o auth_server authenticate_server.go 
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -o libmqtt github.com/goiiot/libmqtt/cmd/libmqtt

go build
```
### 启动服务
``` 
nohup $COMMAND >nohup.out 2>&1 &
```
### workspace
1. 创建工作区
go work init [localModules]
2. 添加go module
go work use ./utils
go work use ../service-common
#### 参考
1. [workspace mode](https://juejin.cn/post/7084584958307598344)
2. [workspace tutorial](https://go.dev/doc/tutorial/workspaces)


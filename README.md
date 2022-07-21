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

### 关键字

``` 
package import 
var constant int/uint bool string true false
struct func return struct type interface defer
switch case fallthrough default if else
for continue break
[] map chan make new 

```
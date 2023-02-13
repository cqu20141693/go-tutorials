### go 工作区
1. 问题
``` 
在Go 1.18以前，如果遇到以下场景：Module A新增了一个feature，Module B需要使用Module A的这个新feature，你有2种方案：
a. 发布Module A的修改到代码仓库，Module B更新依赖的Module A的版本即可
b. 修改Module B的go.mod，使用replace指令把对Module A的依赖指向你本地未发布的Module A所在目录。等Module A发布后，
在发布Module B的时候，再删除Module B的go.mod文件里的replace指令。
```
2. 工作区
```  
你可以在工作区目录维护一个go.work文件来管理你的所有依赖。go.work里的use和replace指令会覆盖工作区目录下的每个Go Module的go.mod文件，
因此没有必要去修改Go Module的go.mod文件了。 
```
#### 使用
使用见分支demo
# go命令

- run
- build
- clean
- env
- fix
- fmt
- generate
- get
- install
- list
- mod
- test
- tool
- version
- vet

go version： 获取Go版本 
go env: 查看Go环境变量 
go help： 查看Go帮助命令 
go get： 获取远程包（需提前安装git或hg） 
go build： 编译并生成可执行程序 
go run： 直接运行程序
 go fmt： 格式化源码 
 go install： 编译包文件以及整个程序 
 go test： go原生提供的单元测试命令 
 go clean： 移除当前源码包和关联源码包里编译生成的文件 
 go tool： 升级Go版本时，修复旧版代码 
 godoc -http:80：开启一个本地80端口的web文档 
 gdb 可执行程序名：调试Go编译出来的文件


./main.exe completion bash > cobra_completion && sudo mv cobra_completion /etc/bash_completion.d/
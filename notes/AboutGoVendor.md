## Go vendor
1. 安装：
  go get -u -v github.com/kardianos/govendor
2. 使用：
    1. 进入目录初始化： govendor init
    2. 添加依赖包： govendor fetch 包名
    3. 更新vendor.json中的包： govendor sync
    
## Go Modules
1. 从 Go 1.11 版本开始，就可以通过设置 GO111MODULE=on 环境变量来启用 Go Modules 管理第三方依赖包，而从 Go 1.13 版本开始，默认已经启用了 Go Modules
2. 1.13之后执行 
    >go mod init
3. 再执行 go run 或 go build 命令时，Go Modules 会自动帮你查找依赖并下载到本地，非常方便！
4. 对于golang.org的包下载不了的问题的处理：
    1. 设置成国内能访问的代理服务器 
        > go env -w GOPROXY=https://goproxy.cn,direct
    2. 查看修改后的值
        >go env
5. 注意： 如果你使用 Goland IDE 进行代码编辑的话，还需要在 File | Settings | Go | Go Modules (vgo) 中设置 proxy 为  `https://goproxy.cn,direct`
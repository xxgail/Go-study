# 关于Go的版本的更新

~~更新的坑~~

#### 1. 关于软件的更新
- 本次更新go1.13
- 目前只更新了公司的电脑（Windows系统）。
- 从go的官网直接下载更新就行，安装途径选择原来的1.12的路径即可。
- 安装之后执行`go version` : `go version go1.13.6 windows/amd64`

#### 2. go1.13.6
- 更新之后的go默认set GOPROXY=https://proxy.golang.org,direct
- 关于go module的更新详见[AboutGoVendor](AboutGoVendor.md)
                       
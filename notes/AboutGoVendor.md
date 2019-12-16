 ## Go vendor
 1. 安装：
  go get -u -v github.com/kardianos/govendor
 2. 使用：
    1. 进入目录初始化： govendor init
    2. 添加依赖包： govendor fetch 包名
    3. 更新vendor.json中的包： govendor sync
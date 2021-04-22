设置代理和jdk
vim /etc/profile
----------------------------------------
export GOPROXY=https://goproxy.io
export GOLAND_JDK=/disk/jdk-11
----------------------------------------
下载并安装Gin和Gorm:
右键点击项目-->open in terminal
go get -u github.com/gin-gonic/gin
go get -u github.com/gin-contrib/multitemplate
go get -u github.com/gin-contrib/sessions
go get -u golang.org/x/sync/errgroup
go get -u github.com/jinzhu/gorm
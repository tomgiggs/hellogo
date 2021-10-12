这是一个golang的学习工程，非常粗糙简陋
之前学过一些go的编程，最近看了无闻老师的一些视频，加快了go入门的速度，在此表示感谢，资源GitHub地址：https://github.com/Unknwon/go-fundamental-programming
包依赖之前使用godep管理，现在用govendor管理，刚入门挺乱的


#依赖管理

## dep/godep
Dep 通过两个 metadata 文件来管理依赖：manifest 文件 Gopkg.toml 和 lock 文件 Gopkg.lock。Gopkg.toml 可以灵活地描述用户的意图，包括依赖的 source、branch、version 等。
Gopkg.lock 仅仅描述依赖的具体状态，例如各依赖的 revision。Gopkg.toml 可以通过命令生产，也可以被用户根据 需要手动修改，Gopkg.lock 是自动生成的，不可以修改。
gopkg.lock和gopkg.toml是godep生成的用来管理依赖

## govendor
安装:go get -u -v github.com/kardianos/govendor
而vendor/vendor.json是govendor生成的，好像两个包管理器都公用vendor文件夹，而且都可以把第三方包放在vendor文件夹底下管理
添加新的依赖：govendor add  github.com/Shopify/sarama 

## go get
有些使用：go get gopkg.in/jcmturner/aescts.v1导入
golang.org 中x目录可以从https://github.com/golang下载再导入进去
有些时候github访问特别缓慢或者是在Google仓库里面的，就到码云（gitee.com）上搜相同的仓库名称，一般情况下用的人多的仓库会在gitee上也有。

## go mod
go mod是官方支持的功能，使用起来很方便，
go init 初始化项目依赖管理
go mod tidy 自动下载依赖
go mod vendor 将依赖打包到项目下vendor目录
启用gomod：
export GO111MODULE=on
配置代理加快代码拉取：
export GOPROXY=https://goproxy.cn

# 语法
golang的包导入方式是导入包名，然后就可以使用packageName.funcName使用包下面的函数了，这一个跟其他语言有点区别，其他语言是导入文件，然后使用FileName.funcName调用，这样说来就是在同一个包底下不能有相同的公开函数
golang的入口只能是package main下面的main函数，如果一个包下面有多个main函数会报错

golang不允许导入了包但是没实际使用，也不允许定义了变量但是没有用到，不然都会报错，这个是跟Python缩进一样严格的要求
golang 不用new来创建对象，直接使用structName{}这样的形式就行

golang 的map不像其他语言一样可以存value不同的键值对，map声明的时候就确定了只能存什么样的值，如果想要存像json那样的数据就需要自定义一个结构体，然后使用结构体来存，结构体里面的键名需要以大写字母开头才能顺利转json，如果需要指定转换后的json键名需要额外指定
golang中没有类这个概念，struct是与类最接近的数据结构，golang使用interface来实现对数据结构声明与传承

make和new的区别：make 只能为 slice、map或 channel 类型分配内存并初始化，同时返回一个有初始值的 slice、map 或 channel 类型引用，不是指针。内建函数 new 用来分配内存，它的第一个参数是一个类型，不是一个值，它的返回值是一个指向新分配类型零值的指针。
golang中数组长度是数组类型的一部分，也就是两个长度不同的string数组是无法比较的，但是如果两个string数组长度是一样的话就可以进行比较

Go语言实现一个接口并不需要显示声明，而是只要你实现了接口中的所有方法就认为你实现了这个接口。这称之为Duck typing
使用go funName()创建一个新的协程可以参考创建一个新的线程，创建完就继续执行后面的语句了并不会导致阻塞。

#go grpc使用
### 工具安装
choco install protoc
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

编译生成.pb.go文件命令：
protoc --proto_path=./proto --plugin=protoc-gen-go.exe --go_out=proto/ --go_opt=paths=source_relative proto/hellogrpc.proto
生成grpc服务代码：
protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  service.proto


使用 go mod添加依赖只需要在项目下执行go get xxx就可以自动添加修改go.mod文件
egg: go get github.com/lxn/walk

#go gui样例
使用go walk制作gui，walk是直接封装的Windows-gui接口
打包需要一个manifest文件
代码写好后，我们直接go build打包是不行的，golang的图形exe需要依赖于manifest才能正常运行。
而go却没有提供资源打包的所有功能，所以要把manifest嵌入exe文件中，还需要一个工具：rsrc。
go get github.com/akavel/rsrc
rsrc -manifest test.manifest -o rsrc.syso
go build -ldflags="-H windowsgui"
参考链接：
https://www.cnblogs.com/sbman/p/9366530.html 
https://github.com/lxn/walk

## thrift 教程
生成代码：thrift -r --gen go .\manager.thrift
https://thrift.apache.org/tutorial/go.html


---------

标准项目目录结构：
https://github.com/golang-standards/project-layout

使用go install可以生成.a文件，并安装到gopath/pkg/windows_amd64/proj_name/pkg_name/xxx.a，结果是以包名命名的，不是一个文件一个.a文件
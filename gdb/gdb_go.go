go env  //看GOROOT路径
gdb -tui ./bin/mt-data-train -d /usr/lib/golang   //-d后接GOROOT
其它参考c++的gdb方法

goland的preference的file watchers里面配置goimports和golangci-lint能自动识别语法错误，前提是：
go get -u github.com/fsgo/go_fmt@v0.1.20200404

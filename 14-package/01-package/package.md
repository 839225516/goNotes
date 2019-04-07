GOPATH 表示当前工作目录

    代码总是保存在$GOPATH/src目录下
    go build,go install,go get 等指令，产生的二进制可执行文件放在$GOPATH/bin目录下
    生成的中间缓存文件会被保存在$GOPATH/pkg下

package 特性

    包要求在同一目录下的所有文件的第一行添加 package 包名
    一个目录下的同级文件归属一个包
    包名可以与其目录不同名
    包名为main的包为应用程序的入口包，编译没有main包时，将无法编译出可执行的文件





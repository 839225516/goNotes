package main

import (
	// $GOPATH/src后面的路径
	"14-package/02-mypkg/mypkg"
	"fmt"
)

func main() {
	fmt.Print(mypkg.Add(int(mypkg.MyConst), 100))
}

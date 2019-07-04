package main

import (
	"blog/pkg"
	"fmt"
)

func main() {
	fmt.Print(pkg.Config.App.PageSize)
}

// var xx *yy 声明后的值为nil，所以无法取赋值给实体*xx，所以一般使用 a := new(xx)，new会分配内存地址

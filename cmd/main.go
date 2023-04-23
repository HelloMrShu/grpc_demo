package main

import (
	"fmt"
	. "github.com/HelloMrShu/grpc_demo/global"
	"os"
	"os/signal"
)

func main() {
	// 初始化
	Initialize()
	//注册服务
	RegisterConsul()

	// 关闭服务器
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("\n关闭服务器 ...")
}

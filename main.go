// TestGOWeb project main.go
package main

import (
	"TestGOWeb/function"
	"TestGOWeb/recordlog"
	"fmt"
	"net/http" //导入go语言内置的http服务

	"TestGOWeb/config"

	"github.com/julienschmidt/httprouter" //导入路由设置的包
)

var (
	routerA *httprouter.Router
	Host    string
	Port    string
)

func init() {
	recordlog.SetLogLevel(recordlog.LevelDebug) //设置日志级别
	setRouter()
}

//设置接口地址
func setRouter() {
	routerA = httprouter.New()
	routerA.GET("/sz.com/getMassage", function.TestFuncObj.GetMessage) //设置接口地址和都映射方法 后面会分析映射的方法
}

var err error

func main() {
	Host = "127.0.0.1"
	Port = "8888"
	recordlog.Debug("项目启动") //设置日志组件
	var port string
	port, err = conf.Conf.String("DB", "DbPort")
	fmt.Println("端口号：", port)
	fmt.Println("Hello World!")
	addr := Host + ":" + Port
	recordlog.Debug("Start HTTP Server Listen In", addr, "...")
	http.ListenAndServe(addr, routerA) //开启http服务
}

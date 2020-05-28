package main

import (
	"acho/config"
	"acho/controller"
	"acho/service"
	"fmt"
	"net/http"
)

/*
程序入口
*/
func main() {
	//开启一个协程，用于定时剔除失效服务或client
	go service.EvictionTicker()
	//用于注册的controller
	controller.RegistryGo()
	//用于心跳的controller
	controller.BeatingGo()
	//从yaml配置文件中读取配置监听的端口
	c := new(config.Conf)
	c.GetConf()
	//如果未配置，则默认监听8080
	if c.Server.Port == "" {
		c.Server.Port = "8080"
	}
	fmt.Printf("已开始监听%s端口\n", c.Server.Port)
	//开启监听端口
	http.ListenAndServe(":"+c.Server.Port, nil)
}

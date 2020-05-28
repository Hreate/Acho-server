package service

import "sync"

//存储注册服务信息的容器，线程安全的map
type infoMap struct {
	Map sync.Map
}

//声明单例容器
var single infoMap

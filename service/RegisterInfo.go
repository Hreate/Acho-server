package service

//需要注册的服务的信息 pojo
type registerInfo struct {
	ApplicationName string `json:"name"`
	Host            string `json:"host"`
	Port            string `json:"port"`
	Date            int64  `json:"date"`
	RenewTime       int64  `json:"renewTime"`
	AliveTime       int64  `json:"aliveTime"`
}

//注册是否成功的响应 结构体
type result struct {
	Status bool `json:"status"`
}

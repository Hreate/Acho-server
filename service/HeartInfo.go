package service

/*
心跳信息的pojo
*/
type heartInfo struct {
	Status string `json:"status"`
	Name   string `json:"name"`
	Host   string `json:"host"`
	Port   string `json:"port"`
}

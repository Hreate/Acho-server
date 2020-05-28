package controller

import (
	"acho/service"
	"net/http"
)

/*
定义心跳的url
*/
func BeatingGo() {
	http.HandleFunc("/acho/heart", service.Renew)
}

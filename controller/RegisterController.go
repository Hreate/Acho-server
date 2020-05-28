package controller

import (
	"acho/service"
	"net/http"
)

/*
定义注册和拉取的url
*/
func RegistryGo() {
	http.HandleFunc("/acho/registry", service.Registry)
	http.HandleFunc("/acho/fetch", service.Fetch)
}

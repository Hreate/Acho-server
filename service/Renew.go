package service

import (
	"encoding/json"
	"net/http"
	"time"
)

/*
处理心跳请求
*/
func Renew(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		len := r.ContentLength
		body := make([]byte, len)
		defer r.Body.Close()
		r.Body.Read(body)
		var heartInfo heartInfo
		json.Unmarshal(body, &heartInfo)
		if flag := single.update(&heartInfo); flag {
			w.WriteHeader(200)
		}
	}
}

/*
更新续约时间
*/
func (single *infoMap) update(h *heartInfo) bool {
	if h.Status == "up" {
		if rInfo, ok := single.Map.Load(h.Name); ok {
			for _, value := range rInfo.([]*registerInfo) {
				if value.Host == h.Host && value.Port == h.Port {
					value.Date = time.Now().Unix()
				}
			}
			return true
		}

		//c.Map.Range(func(key, value interface{}) bool {
		//	for _, v := range value.([]*registerInfo) {
		//		fmt.Println(v)
		//	}
		//	return true
		//})
		//return true
	}
	return false
}

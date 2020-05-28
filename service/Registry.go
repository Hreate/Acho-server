package service

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

/*
处理注册请求
*/
func Registry(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		len := r.ContentLength
		body := make([]byte, len)
		defer r.Body.Close()
		r.Body.Read(body)
		var serviceInfo registerInfo
		json.Unmarshal(body, &serviceInfo)
		serviceInfo.Date = time.Now().Unix()
		single.save(&serviceInfo)
		r := result{false}
		rJson, err := json.Marshal(r)
		if err != nil {
			log.Printf("json marshal failed: %v", err)
		}
		//注册成功响应
		w.Header().Set("Content-type", "application/json")
		w.Write(rJson)
		w.WriteHeader(200)
		//single.Map.Range(print)
	}
}

/*
将注册信息存储到容器中
*/
func (single *infoMap) save(s *registerInfo) {
	if serviceList, ok := single.Map.Load(s.ApplicationName); ok {
		for _, s := range serviceList.([]*registerInfo) {
			if s.Host != s.Host || s.Port != s.Port {
				serviceList = append(serviceList.([]*registerInfo), s)
				single.Map.Store(s.ApplicationName, serviceList)
			}
		}
	} else {
		newServiceList := make([]*registerInfo, 0, 10)
		newServiceList = append(newServiceList, s)
		single.Map.Store(s.ApplicationName, newServiceList)
	}
}

/*func print(s , i interface{}) bool {
	fmt.Printf("%s, %v\n", s, i)
	return true
}*/

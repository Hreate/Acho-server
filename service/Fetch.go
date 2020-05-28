package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
处理拉取请求
*/
func Fetch(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		m := single.toMap()

		for _, value := range m {
			for _, v := range value {
				fmt.Println(v)
			}
		}

		infoBytes, err := json.Marshal(m)
		if err != nil {
			log.Printf("json marshal failed: %v\n", err)
		}
		//fmt.Println(string(infoBytes))
		w.Header().Set("Content-type", "application/json")
		w.Write(infoBytes)
		w.WriteHeader(200)
	}
}

/*
将sync.Map转成能json序列化的map
*/
func (single *infoMap) toMap() map[string][]*registerInfo {
	var m = make(map[string][]*registerInfo)
	single.Map.Range(func(key, value interface{}) bool {
		m[key.(string)] = value.([]*registerInfo)
		return true
	})
	return m
}

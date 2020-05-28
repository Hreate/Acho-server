package service

import (
	"acho/config"
	"fmt"
	"log"
	"strconv"
	"time"
)

/*
定时清理过期服务或client任务
*/
func EvictionTicker() {
	c := new(config.Conf)
	c.GetConf()
	if c.Server.EvictionS == 0 {
		c.Server.EvictionS = 60
	}
	s := strconv.FormatInt(c.Server.EvictionS, 10) + "s"
	d, err := time.ParseDuration(s)
	if err != nil {
		log.Printf("parse duration failed: %v", err)
	}
	var ticker = time.NewTicker(d)
	for range ticker.C {
		fmt.Println("定时执行剔除任务")
		single.eviction()
		//测试用，打印到控制台目前存活的服务信息
		single.Map.Range(func(key, value interface{}) bool {
			fmt.Printf("%s, %v\n", key, value)
			return true
		})
	}
}

/*
剔除失效的服务或主机
*/
func (single *infoMap) eviction() {

	single.Map.Range(func(key, value interface{}) bool {
		//获取服务对应的client切片
		v := value.([]*registerInfo)
		//每次循环初始化切片
		l := make([]*registerInfo, 0, len(v))
		//获取每台client。如果没过期，则存入l
		for _, r := range v {
			if (*r).AliveTime+(*r).Date > time.Now().Unix() {
				l = append(l, r)
			}
		}
		//如果最后l的长度为0，说明该服务没有一台存活的client，直接剔除该服务
		if len(l) == 0 {
			single.Map.Delete(key)
		} else {
			single.Map.Store(key, l)
		}
		return true
	})
}

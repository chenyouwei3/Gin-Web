package main

import (
	"loopy-manager/initialize"
	"loopy-manager/internal/router"
)

func init() {
	initialize.InitConfig()
}

type String string

func (d String) Len() int {
	return len(d)
}

func main() {
	engine := router.GetEngine()
	if err := engine.Run(":8099"); err != nil {
		panic(err)
	}
	////service.GetRequestLog0()
	//lru := lru.New(int64(0), nil)
	//fmt.Println("lru start")
	//lru.Add("/user.test", String("1234"))
	//if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "1234" {
	//	fmt.Println("cache hit key1=1234 failed")
	//}
	//if _, ok := lru.Get("key2"); ok {
	//	fmt.Println("cache miss key2 failed")
	//}
}

package main

import (
	"loopy-manager/app/router"
	"loopy-manager/initialize"
)

func init() {
	initialize.InitConfig()
}

func main() {
	//go ticker.CornTicker()
	engine := router.GetEngine()
	if err := engine.Run(":8098"); err != nil {
		panic(err)
	}
}

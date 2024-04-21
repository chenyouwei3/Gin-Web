package main

import (
	"loopy-manager/initialize"
	"loopy-manager/internal/controller/ticker"
	"loopy-manager/internal/router"
)

func init() {
	initialize.InitConfig()
}

func main() {
	go ticker.CornTicker()
	engine := router.GetEngine()
	if err := engine.Run(":8098"); err != nil {
		panic(err)
	}
}

package main

import (
	"loopy-manager/initialize"
	"loopy-manager/internal/router"
)

func init() {
	initialize.InitConfig()
}

func main() {
	engine := router.GetEngine()
	if err := engine.Run(":8099"); err != nil {
		panic(err)
	}
}

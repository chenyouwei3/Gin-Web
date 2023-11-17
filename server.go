package main

import (
	"loopy-manager/initialize"
	"loopy-manager/router"
)

func main() {
	initialize.Init()
	defer initialize.DataBaseClose()
	engine := router.GetEngine()
	if err := engine.Run(":8091"); err != nil {
		panic(err)
	}

}

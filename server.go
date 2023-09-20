package main

import (
	"LoopyTicker/initialize"
	"LoopyTicker/router"
)

func main() {
	engine := router.GetEngine()
	initialize.Init()
	defer initialize.DataBaseClose()

	if err := engine.Run(":8091"); err != nil {
		panic(err)
	}

}

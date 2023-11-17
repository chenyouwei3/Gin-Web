package initialize

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"loopy-manager/global"
)

// 初始化雪花算法
func SnowFlakeInit() {
	if global.UserSnowFlake == nil {
		node, err := snowflake.NewNode(1)
		if err != nil {
			fmt.Println("snowflake init:", err)
		}
		global.UserSnowFlake = node
	}
	if global.RoleSnowFlake == nil {
		node, err := snowflake.NewNode(2)
		if err != nil {
			fmt.Println("snowflake init:", err)
		}
		global.RoleSnowFlake = node
	}
	if global.ApiSnowFlake == nil {
		node, err := snowflake.NewNode(3)
		if err != nil {
			fmt.Println("snowflake init:", err)
		}
		global.ApiSnowFlake = node
	}
	if global.LogSnowFlake == nil {
		node, err := snowflake.NewNode(4)
		if err != nil {
			fmt.Println("snowflake init:", err)
		}
		global.LogSnowFlake = node
	}

}

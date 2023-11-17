package global

import "github.com/bwmarrin/snowflake"

// *snowflake.Node可能是指Snowflake算法中的节点。
// 在Snowflake算法中，节点用于标识不同的分布式节点。每个节点都有自己的唯一ID，该ID在生成ID时帮助区分不同的节点，以避免ID冲突
// *snowflake.ID可能是指Snowflake算法生成的唯一ID。Snowflake算法生成的ID通常包含一个时间戳部分、
// 一个节点ID部分和一个序列号部分，用于保证全局唯一性并具有时间有序性
var (
	UserSnowFlake *snowflake.Node
	RoleSnowFlake *snowflake.Node
	ApiSnowFlake  *snowflake.Node
	LogSnowFlake  *snowflake.Node
)

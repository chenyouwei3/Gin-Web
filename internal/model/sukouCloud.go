package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type SuKonToken struct { //速控云Token
	Code           int                       `form:"code" json:"code"` //状态码
	SukonTokenData `form:"data" json:"data"` //数据集，状态success为true时返回，否则为空
	Msg            string                    `form:"msg" json:"msg"`         //返回消息
	Success        bool                      `form:"success" json:"success"` //状态
}

type SukonTokenData struct { //速控云Token(子结构体)
	Expire int    `form:"expire" json:"expire"`
	Token  string `form:"token" json:"token"`
	Type   string `form:"type" json:"type"`
}

type ProjectBox struct { // 接收获取项目Box时返回的数据
	Code    int           `form:"code" json:"code"`
	Data    []BoxDataType `form:"data" json:"data"`
	Msg     string        `form:"msg" json:"msg"`
	Success bool          `form:"success" json:"success"`
}
type BoxDataType struct { //接收获取项目Box时返回的数据(子结构体)
	BoxId       string `form:"boxId" json:"boxId"`             //BOXID
	Name        string `form:"name" json:"name"`               //BOX名称
	ProjectType string `form:"projectType" json:"projectType"` //项目类型：0：自由项目，1：模板项目
	Serlnum     string `form:"serlnum" json:"serlnum"`         //	序列号
	Status      string `form:"status" json:"status"`           //状态
}

type SuKonProject struct { //数控云项目
	Code int `form:"code" json:"code"`
	Data []struct {
		Id          string `form:"id" json:"id"`
		Name        string `form:"name" json:"name"`
		ProjectType string `form:"projectType" json:"projectType"`
		Status      string `form:"status" json:"status"` //状态 0：离线，1：在线
	} `form:"data" json:"data"`
	Msg     string `form:"msg" json:"msg"`
	Success bool   `form:"success" json:"success"`
}

type BoxPlc struct { // 接收获取项目Box Plc时返回的数据
	Code    int       `form:"code" json:"code"`
	Data    []PlcType `form:"data" json:"data"`
	Msg     string    `form:"msg" json:"msg"`
	Success bool      `form:"success" json:"success"`
}
type PlcType struct { //接收获取项目Box Plc时返回的数据(子结构体)
	PlcId  string `form:"plcId" json:"plcId"`
	Name   string `form:"name" json:"name"`
	Status string `form:"status" json:"status"`
}

type BoxVariant struct { // 接收获取项目变量时返回的数据
	Code    int           `form:"code" json:"code"`
	Data    []VariantType `form:"data" json:"data"`
	Msg     string        `form:"msg" json:"msg"`
	Success bool          `form:"success" json:"success"`
}
type VariantType struct { //接收获取项目变量时返回的数据(子结构体)
	Addr      string `form:"addr" json:"addr"`
	Name      string `form:"name" json:"name"`
	Type      string `form:"type" json:"type"`
	VariantId string `form:"variantId" json:"variantId"`
}

type Box struct { //box数据
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	DeviceTypeId primitive.ObjectID `bson:"deviceTypeId" json:"deviceTypeId"` //设备类型id 对应设备类型表
	BoxId        string             `bson:"boxId" json:"boxId"`               //Box上唯一表示
	Data         []BoxData          `bson:"data" json:"data"`
	CreateTime   string             `bson:"createTime,omitempty" json:"createTime,omitempty"` //创建时间
	UpdateTime   string             `bson:"updateTime,omitempty" json:"updateTime,omitempty"` //更新时间
}
type BoxData struct { //box数据(子结构体)
	SensorId    string          `bson:"sensorId" json:"sensorId"`                           //plc对应变量的地址位
	SensorName  string          `bson:"sensorName" json:"sensorName"`                       //Box设备名称
	StartTime   string          `bson:"startTime,omitempty" json:"startTime,omitempty"`     //石墨化送电时刻 其他工艺不用
	StoveNumber string          `bson:"stoveNumber,omitempty" json:"stoveNumber,omitempty"` //石墨化炉号 其他工艺不用
	Detail      []BoxDataDetail `bson:"detail" json:"detail"`                               //详细数据信息
}

type BoxDataDetail struct { //box数据(子结构体)(子结构体)
	Key   string `bson:"key" json:"key"`     //Box检测的key值
	Value string `bson:"value" json:"value"` //value值
	Unit  string `bson:"unit" json:"unit"`   //单位
}

type RealtimeData struct { // 接收实时数据变量时返回的数据
	Code    int              `form:"code" json:"code"`
	Data    []RealtimeDataVo `form:"data" json:"data"`
	Msg     string           `form:"msg" json:"msg"`
	Success bool             `form:"success" json:"success"`
}

type RealtimeDataVo struct { // 接收实时数据变量时返回的数据(子结构体)
	Id     string `form:"id" json:"id"`
	Status string `form:"status" json:"status"`
	Time   int    `form:"time" json:"time"` //单位s
	Value  string `form:"value" json:"value"`
}

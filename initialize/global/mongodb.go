package global

import "go.mongodb.org/mongo-driver/mongo"

var (
	MongodbClient                    *mongo.Client
	DeviceColl                       *mongo.Collection
	ImmersionHisData                 *mongo.Collection //浸渍
	WestCraneCarHisData              *mongo.Collection //西跨吸料天车
	GraphitingHisData                *mongo.Collection //石墨化
	TunnelWetElectricHisDataColl     *mongo.Collection //隧道窑湿电
	RoastWetElectricHisDataColl      *mongo.Collection //隧道窑湿电
	GraphitingWetElectricHisDataColl *mongo.Collection //石墨化湿电
	EastCraneCarHisData              *mongo.Collection //东跨跨吸料天车
	TunnelHisDataColl                *mongo.Collection //隧道窑
	CrucibleHisDataColl              *mongo.Collection //坩埚
	CalcinationHisDataColl           *mongo.Collection //煅烧脱销
	FormPlcHisDataColl               *mongo.Collection //压型
	RoastDenitrificationHisColl      *mongo.Collection //存储焙烧脱硝
	FourSeaStoreFormHisColl          *mongo.Collection //四海成型plc
)

package config

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"loopy-manager/initialize/global"
)

func MongodbInit(config MongodbConfig) {
	if global.MongodbClient == nil {
		global.MongodbClient = getMongoClient(config.Address)
	}
	smartGraphiteHBClone := global.MongodbClient.Database("smartGraphiteHB-Clone")
	{
		global.DeviceColl = smartGraphiteHBClone.Collection("device")
		//速控云

	}
	sukonCloud := global.MongodbClient.Database("sukouCloud")
	{
		//速控云
		global.ImmersionHisData = sukonCloud.Collection("ImmersionHisData")
		global.WestCraneCarHisData = sukonCloud.Collection("WestCraneCarHisData")
		global.GraphitingHisData = sukonCloud.Collection("GraphitingHisData")
		global.TunnelWetElectricHisDataColl = sukonCloud.Collection("TunnelWetElectricHisDataColl")
		global.RoastWetElectricHisDataColl = sukonCloud.Collection("RoastWetElectricHisDataColl")
		global.GraphitingWetElectricHisDataColl = sukonCloud.Collection("GraphitingWetElectricHisDataColl")
		global.EastCraneCarHisData = sukonCloud.Collection("EastCraneCarHisData")
		global.TunnelHisDataColl = sukonCloud.Collection("TunnelHisDataColl")
		global.CrucibleHisDataColl = sukonCloud.Collection("CrucibleHisDataColl")
		global.CalcinationHisDataColl = sukonCloud.Collection("CalcinationHisDataColl")
		global.FormPlcHisDataColl = sukonCloud.Collection("FormPlcHisDataColl")
		global.RoastDenitrificationHisColl = sukonCloud.Collection("RoastDenitrificationHisColl *mongo.Collection")
		global.FourSeaStoreFormHisColl = sukonCloud.Collection("FourSeaStoreFormHisColl")
	}
}

func getMongoClient(uri string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(uri)
	MongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logrus.Fatalln(err)
	}
	if err = MongoClient.Ping(context.TODO(), nil); err != nil {
		log.Fatalln(err)
	}
	return MongoClient
}

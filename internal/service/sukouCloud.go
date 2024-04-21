package service

import (
	"context"
	"loopy-manager/initialize/global"
	"loopy-manager/internal/model"
	"loopy-manager/pkg/utils"

	"fmt"
	"github.com/sirupsen/logrus"
)

func StoreDippingData(box model.Box, data model.RealtimeData) { // 存储浸渍分钟数据x
	newBox := utils.GetSKCloudHisData(box, data, "2da580adb26b4a12accd4aec80e04656:")
	_, err := global.ImmersionHisData.InsertOne(context.Background(), newBox)
	if err != nil {
		logrus.Error(box, "存储分钟历史数据失败:", err.Error())
	}
}

func StoreWestAirCarData(box model.Box, data model.RealtimeData) { // 存储西跨吸料天车分钟数据x
	newBox := utils.GetSKCloudHisData(box, data, "b46a0faf11cc4000a4c290eba5cc949a:")
	_, err := global.WestCraneCarHisData.InsertOne(context.Background(), newBox)
	if err != nil {
		logrus.Error(box, "存储分钟历史数据失败:", err.Error())
	}

}

func StoreGraphitingData(box model.Box, data model.RealtimeData) { //存储石墨化分钟数据
	newBox := utils.GetSKCloudHisData(box, data, "be67c2b8216e49e8981a95663413f115:")
	_, err := global.GraphitingHisData.InsertOne(context.Background(), newBox)
	if err != nil {
		logrus.Error(box, "存储分钟历史数据失败:", err.Error())
	}

}

func StoreTunnelWetElectricData(box model.Box, data model.RealtimeData) { // 存储隧道窑湿电数据x
	newBox := utils.GetSKCloudHisData(box, data, "5cba298477bc456ab1a2bd06e35cb0d8:")
	_, err := global.TunnelWetElectricHisDataColl.InsertOne(context.Background(), newBox)
	if err != nil {
		logrus.Error(box, "存储分钟历史数据失败:", err.Error())
	}

}

func StoreRoastingWetElectricData(box model.Box, data model.RealtimeData) { // 存储焙烧湿电数据x
	newBox := utils.GetSKCloudHisData(box, data, "01e844f884844aa2bb5d1cab87316c17:")
	_, err := global.RoastWetElectricHisDataColl.InsertOne(context.Background(), newBox)
	if err != nil {
		logrus.Error(box, "存储分钟历史数据失败:", err.Error())
	}
}

func StoreGraphiteWetElectricData(box model.Box, data model.RealtimeData) { // 存储石墨化湿电数据x
	newBox := utils.GetSKCloudHisData(box, data, "69fb82a9cba744188cab9da766787f25:")
	_, err := global.GraphitingWetElectricHisDataColl.InsertOne(context.Background(), newBox)
	if err != nil {
		logrus.Error(box, "存储分钟历史数据失败:", err.Error())
	}
}

func StoreEarthAirCarData(box model.Box, data model.RealtimeData) { // 存储东跨跨吸料天车分钟数据x
	newBox := utils.GetSKCloudHisData(box, data, "f73fe0d8688046e088bb073849aa0c3f:")
	_, err := global.EastCraneCarHisData.InsertOne(context.Background(), newBox)
	if err != nil {
		logrus.Error(box, "存储分钟历史数据失败:", err.Error())
	}
}

func StoreTunnelData(box model.Box, data model.RealtimeData) { // 存储隧道窑分钟数据x
	fmt.Println(data, "+++++++++++++++", box.BoxId)

	//newBox := utils.GetSKCloudHisData(box, data, "9f62bc0edbd542b2bec159ac8f023509:")
	//_, err := global.TunnelHisDataColl.InsertOne(context.Background(), newBox)
	//if err != nil {
	//	logrus.Error(box, "存储分钟历史数据失败:", err.Error())
	//}
}

func StoreCrucibleData(box model.Box, data model.RealtimeData) { // 存储坩埚分钟数据x
	newBox := utils.GetSKCloudHisData(box, data, "9bd62f734af94dc0b0641817ac2807e9:")
	_, err := global.CrucibleHisDataColl.InsertOne(context.Background(), newBox)
	if err != nil {
		logrus.Error(box, "存储分钟历史数据失败:", err.Error())
	}
}

func StoreCalcinationData(box model.Box, data model.RealtimeData) { //煅烧脱销
	newBox := utils.GetSKCloudHisData(box, data, "ef62aa2e44204b5d82463b72a86f9621:")
	_, err := global.CalcinationHisDataColl.InsertOne(context.Background(), newBox)
	if err != nil {
		logrus.Error(box, "存储分钟历史数据失败:", err.Error())
	}
}

func StoreFormPlcData(box model.Box, data model.RealtimeData) { //压型
	newBox := utils.GetSKCloudHisData(box, data, "65d27a491d744a0e91b4d8e6db628887:")
	_, err := global.FormPlcHisDataColl.InsertOne(context.Background(), newBox)
	if err != nil {
		logrus.Error(box, "存储分钟历史数据失败:", err.Error())
	}
}

func StoreRoastDenitrificationData(box model.Box, data model.RealtimeData) { //焙烧脱硝
	newBox := utils.GetSKCloudHisData(box, data, "52980204e2dc4ce9907196441c6f9a32:")
	_, err := global.RoastDenitrificationHisColl.InsertOne(context.Background(), newBox)
	if err != nil {
		logrus.Error(box, "存储分钟历史数据失败:", err.Error())
	}
}

func FourSeaStoreFormPlcData(box model.Box, data model.RealtimeData) {
	newBox := utils.GetSKCloudHisData(box, data, "97509d2212bf4b1cb5bc3ea8dd8649d7:")
	//历史表插入
	_, err := global.FourSeaStoreFormHisColl.InsertOne(context.Background(), newBox)
	if err != nil {
		logrus.Error(box, "存储分钟历史数据失败:", err.Error())
	}
}

package sukonCloud

import (
	"github.com/sirupsen/logrus"
	"loopy-manager/initialize/global"
	"loopy-manager/internal/model"
	"loopy-manager/internal/service"
	"loopy-manager/pkg/utils"
	"net/url"
)

// 获取plc
func BoxPlc(boxId string) {
	URL := "http://sukon-cloud.com/api/v1/base/boxPlcs"
	urlValues := url.Values{}
	urlValues.Add("token", global.SukonCloudToken)
	urlValues.Add("boxId", boxId)
	var data model.BoxPlc
	data, err := utils.SukouCloudHTTP(data, URL, urlValues)
	if err != nil {
		logrus.Error("http请求失败:", err)
	}
	urlValues.Del("boxId")
	if data.Data == nil {
		logrus.Error("plcId为空", data)
		return
	}
	for _, a := range data.Data {
		FindVariant(boxId, a.PlcId)
	}
}

// 获取变量
func FindVariant(boxId, plcId string) {
	URL := "http://sukon-cloud.com/api/v1/base/boxVariants"
	urlValues := url.Values{}
	urlValues.Add("token", global.SukonCloudToken)
	urlValues.Add("boxId", boxId)
	urlValues.Add("plcId", plcId)
	//获取每个sid下变量
	var data model.BoxVariant
	data, err := utils.SukouCloudHTTP(data, URL, urlValues)
	if err != nil {
		logrus.Error("http请求失败:", err)
	}
	urlValues.Del("boxId")
	urlValues.Del("plcId")
	if data.Success == false {
		logrus.Error("获取变量失败")
		return
	}
	//得到用于获取实时数据的变量字符串
	var variantIds string
	for i, variant := range data.Data {
		if len(data.Data) == 1 {
			variantIds = boxId + "(" + variant.VariantId + ")"
		} else {
			if i == len(data.Data)-1 {
				variantIds = variantIds + variant.VariantId + ")"
			} else {
				if i == 0 {
					variantIds = variantIds + boxId + "("
				}
				variantIds = variantIds + variant.VariantId + ":"
			}
		}
	}
	//box数据
	var box model.Box
	box.BoxId = boxId
	var detail []model.BoxDataDetail
	for _, a := range data.Data {
		detail = append(detail, model.BoxDataDetail{
			Key:   a.Name,
			Value: "",
			Unit:  "",
		})
	}
	box.Data = append(box.Data, model.BoxData{
		SensorId:   "",
		SensorName: "",
		Detail:     detail,
	})
	GetBoxRealTimeData(variantIds, box) //获取实时数据;
}

// 获取实时数据
func GetBoxRealTimeData(variantIds string, box model.Box) {
	URL := "http://sukon-cloud.com/api/v1/data/realtimeDatas"
	urlValues := url.Values{}
	urlValues.Add("token", global.SukonCloudToken)
	urlValues.Add("variantIds", variantIds)
	var data model.RealtimeData
	data, err := utils.SukouCloudHTTP(data, URL, urlValues)
	if err != nil {
		logrus.Error("http请求失败:", err)
	}
	urlValues.Del("variantIds")
	switch box.BoxId {
	//case "2da580adb26b4a12accd4aec80e04656":
	//	service.StoreDippingData(box, data) //浸渍x
	//case "b46a0faf11cc4000a4c290eba5cc949a":
	//	service.StoreWestAirCarData(box, data) //西跨吸料天车x
	//case "be67c2b8216e49e8981a95663413f115":
	//	service.StoreGraphitingData(box, data) //石墨化x
	//case "5cba298477bc456ab1a2bd06e35cb0d8":
	//	service.StoreTunnelWetElectricData(box, data) //隧道窑湿电x
	//case "01e844f884844aa2bb5d1cab87316c17":
	//	service.StoreRoastingWetElectricData(box, data) //焙烧湿电x
	//case "69fb82a9cba744188cab9da766787f25":
	//	service.StoreGraphiteWetElectricData(box, data) //石墨化湿电x
	//case "f73fe0d8688046e088bb073849aa0c3f":
	//	service.StoreEarthAirCarData(box, data) //东跨吸料天车x
	case "9f62bc0edbd542b2bec159ac8f023509":
		service.StoreTunnelData(box, data) //隧道窑x
	//case "9bd62f734af94dc0b0641817ac2807e9":
	//	service.StoreCrucibleData(box, data) //坩埚x
	//case "ef62aa2e44204b5d82463b72a86f9621":
	//	service.StoreCalcinationData(box, data) //煅烧脱硝x
	//case "65d27a491d744a0e91b4d8e6db628887":
	//	service.StoreFormPlcData(box, data) //压型x
	//case "52980204e2dc4ce9907196441c6f9a32":
	//	service.StoreRoastDenitrificationData(box, data) //焙烧脱硝x
	//case "97509d2212bf4b1cb5bc3ea8dd8649d7":
	//	service.FourSeaStoreFormPlcData(box, data) //四海成型x
	default:
		logrus.Println("没有这个设备")
	}
}

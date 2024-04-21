package utils

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"log"
	"loopy-manager/initialize/global"
	"loopy-manager/internal/model"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func IsProd() bool {
	return runtime.GOOS != "linux"
}

func SukonToken() {
	timeUnix := time.Now().UnixNano() / 1e6
	random := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000)) //3.随机6位字母与数字的字符串
	timestamp := fmt.Sprintf("%v", timeUnix)                                                    //4.当前时间戳（13位）
	ctx := md5.New()                                                                            //md5加密
	ctx.Write([]byte(global.SuKonUid + global.SuKonSid + random + timestamp))
	signature := strings.ToUpper(hex.EncodeToString(ctx.Sum(nil))) //签名转换成字符串和大写32位
	tokenUrl := "http://sukon-cloud.com/api/v1/token/initToken"
	body := strings.NewReader(fmt.Sprintf("uid=%s&sid=%s&random=%s&timestamp=%s&signature=%s", global.SuKonUid, global.SuKonSid, random, timestamp, signature))
	var data model.SuKonToken
	res, err := http.Post(tokenUrl, "application/x-www-form-urlencoded", body)
	if err != nil {
		log.Println("请求错误:", err.Error())
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Println("关闭获取速控云token失败:", err)
		}
	}()
	body0, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("响应错误:", err.Error())
	}
	err = json.Unmarshal(body0, &data) //非流式传输
	if err != nil {
		log.Println("解析错误:", err.Error())
	}
	global.SukonCloudToken = data.SukonTokenData.Token //定义全局变量
	hour := int(math.Floor(float64(data.SukonTokenData.Expire / 3600)))
	//token时效等于0,重新获取token
	if hour <= 0 {
		time.Sleep(time.Second * 10)
		SukonToken()
		return
	} else {
		if hour == 24 {
			hour = hour - 1
		}
		t := strconv.Itoa(hour)
		global.Spec = "0 0 */" + t + " * * *"
	}
	return
}

func GetSKCloudHisData(box model.Box, data model.RealtimeData, str string) model.Box { //数据标准格式化
	//查找设备
	var BoxDevice model.Device
	if err := global.DeviceColl.FindOne(context.TODO(), bson.M{"code": box.BoxId}).Decode(&BoxDevice); err != nil {
		log.Println(box.BoxId+"设备不存在", err)
	}
	box.DeviceTypeId = BoxDevice.DeviceTypeId
	box.CreateTime = TimeFormat(time.Now())
	//处理单位
	if len(BoxDevice.Sensors) == 1 {
		box.Data[0].SensorId = BoxDevice.Sensors[0].Code
		box.Data[0].SensorName = BoxDevice.Sensors[0].Name
		for i, a := range box.Data[0].Detail {
			for _, b := range BoxDevice.Sensors[0].DetectionValue {
				if a.Key == b.Key {
					box.Data[0].Detail[i].Unit = b.Unit
					continue
				}
			}
		}
	}
	for i, _ := range box.Data[0].Detail { //处理值
		for _, b := range data.Data {
			array := strings.Split(b.Id, str)
			id := strconv.Itoa(i)
			if array[1] == id {
				box.Data[0].Detail[i].Value = b.Value
				continue
			}
		}
	}
	return box
}

func SukouCloudHTTP[T model.SuKonProject | model.ProjectBox | model.BoxPlc | model.BoxVariant | model.RealtimeData](data T, url string, urlValues url.Values) (T, error) {
	var t T
	res, err := http.PostForm(url, urlValues)
	if err != nil {
		return t, fmt.Errorf("请求错误:%w", err)
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			logrus.Error("获取box错误:", err)
		}
	}()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return t, fmt.Errorf("响应错误:%w", err)
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return t, fmt.Errorf("解析错误:%w", err)
	}
	return data, nil
}

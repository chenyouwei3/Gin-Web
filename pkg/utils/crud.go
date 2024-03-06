package utils

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"time"
)

func GetPage(currPage, pageSize string) (int, int, error) {
	curr, err := strconv.Atoi(currPage)
	if err != nil {
		return 0, 0, err
	}
	size, err := strconv.Atoi(pageSize)
	if err != nil {
		return 0, 0, err
	}
	skip := (curr - 1) * size
	return skip, size, nil
}

func GetTime(startTime, endTime string) (filter bson.M) {
	if startTime != "" && endTime != "" {
		filter = bson.M{
			"createTime": bson.M{
				"$gte": startTime,
				"$lte": endTime,
			},
		}
		return filter
	}
	filter = bson.M{}
	return filter
}

func TimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func GetNowTime() time.Time {
	parsedTime, err := time.Parse("2006-01-02 15:04:05", TimeFormat(time.Now()))
	if err != nil {
		fmt.Println("解析时间字符串失败：", err)
	}
	return parsedTime
}

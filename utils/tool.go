package utils

import (
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

func TimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

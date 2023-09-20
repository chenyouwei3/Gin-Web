package utils

import "strconv"

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

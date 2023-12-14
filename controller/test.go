package controller

import (
	"net/http"
	"os"
)

func Test(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://zouzh.cn/", http.StatusNotFound)
	content := []byte("重定向")
	err := os.WriteFile("test.txt", content, 0644)
	if err != nil {
		panic(err)
	}
}

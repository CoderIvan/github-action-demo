package main

import (
	"fmt"
	"net/http"
)

const PORT = 80
const VERSION = "1.0.0"

func versionHandler(w http.ResponseWriter, r *http.Request) {
	// 设置返回的内容为JSON格式
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, VERSION)
}

func main() {
	// 注册路由
	http.HandleFunc("/api/version", versionHandler)

	// 启动HTTP服务器
	fmt.Printf("Server started at :%d\n", PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

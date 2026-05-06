package main

import (
	"fmt"
	"log"
	"os"

	"eit-aps-api/api" // 这里的路径需与 go.mod 中的模块名一致
)

func main() {
	// 从环境变量读取端口，默认为 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 调用共享的路由初始化函数
	r := handler.SetupRouter()

	fmt.Printf("Server is running on http://localhost:%s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

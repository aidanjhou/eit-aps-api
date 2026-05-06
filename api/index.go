package handler

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	app  *gin.Engine
	once sync.Once
)

// SetupRouter 封装核心路由逻辑，供 Vercel 和 main.go 共同调用
func SetupRouter() *gin.Engine {
	// 生产环境建议设置为 ReleaseMode
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// 示例接口
	r.GET("/api/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Hello from Gin on Vercel/Docker!",
			"runtime": "Go 1.25.0+",
		})
	})

	// 健康检查接口
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	return r
}

// Handler 符合 Vercel Go Runtime 要求的导出函数
func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(func() {
		app = SetupRouter()
	})
	app.ServeHTTP(w, r)
}

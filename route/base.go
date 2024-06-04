package route

import (
	"fmt"
	"official_site/global"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var route *gin.Engine

// 单例
var once sync.Once

func GetRout() *gin.Engine {
	once.Do(func() {
		gin.SetMode(global.G_CONFIG.Server.Mode)
		route = gin.New()
		route.Use(gin.Logger(), gin.Recovery()).
			Use(cors.New(cors.Config{
				AllowOriginFunc:  func(origin string) bool { return true },
				AllowMethods:     []string{"OPTIONS", "GET", "POST", "PUT", "DELETE", "PATCH"},
				AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
				AllowCredentials: true,
				MaxAge:           12 * time.Hour,
			}))
		fmt.Println("====== Mode [Router-Gin] Loading success ======")
	})
	return route
}

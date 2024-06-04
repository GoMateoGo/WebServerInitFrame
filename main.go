package main

import (
	"flag"
	"fmt"
	"official_site/core"
	"official_site/global"
	"official_site/route"

	"go.uber.org/zap"
)

var cfgPath = ""

func main() {
	fmt.Println("======================================")
	fmt.Println("Server Starting...")
	flag.StringVar(&cfgPath, "w", "config.yaml", "选择配置文件.") // 命令行参数 -w 配置文件路径
	flag.Parse()
	global.G_VP = core.Viper(cfgPath) // 初始化Viper
	global.G_LOG = core.Zap()         // 初始化zap日志库
	zap.ReplaceGlobals(global.G_LOG)
	global.G_DB = core.InitDB(&global.G_CONFIG.Mysql) // 初始化Mysql数据库

	fmt.Println("======================================")
	address := fmt.Sprintf(":%s", global.G_CONFIG.Server.Port)

	route := route.GetRout()
	route.Run(address)
}

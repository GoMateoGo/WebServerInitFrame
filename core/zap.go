package core

import (
	"fmt"
	"official_site/core/internal"
	"official_site/global"
	"official_site/utils"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (logger *zap.Logger) {

	if ok, _ := utils.PathExists(global.G_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.G_CONFIG.Zap.Director)
		_ = os.Mkdir(global.G_CONFIG.Zap.Director, os.ModePerm)
	}
	levels := global.G_CONFIG.Zap.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := internal.NewZapCore(levels[i])
		cores = append(cores, core)
	}
	logger = zap.New(zapcore.NewTee(cores...))
	if global.G_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	fmt.Println("====== Mode [Logger] Loading success ======")
	return logger
}

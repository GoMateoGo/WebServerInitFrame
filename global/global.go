package global

import (
	"official_site/config"

	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	G_CONFIG config.Config //配置信息
	G_VP     *viper.Viper  //配置文件实例
	G_LOG    *zap.Logger   //日志实例
	G_DB     *xorm.Engine  //数据库实例
)

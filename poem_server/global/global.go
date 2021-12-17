package global

import (
	"github.com/qiniu/qmgo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"poem_server/config"
)

var (
	POEM_DB                  *gorm.DB
	POEM_CONFIG              config.Server
	POEM_VP                  *viper.Viper
	POEM_LOG                 *zap.Logger
	POEM_MONGO               *qmgo.Database
	POEM_Concurrency_Control = &singleflight.Group{}
)

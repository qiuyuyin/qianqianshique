package main

import (
	"poem_server/core"
	"poem_server/global"
	"poem_server/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title 千千诗阙API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.POEM_VP = core.Viper()          // 初始化配置库
	global.POEM_LOG = core.Zap()           // 初始化日志工具
	global.POEM_DB = initialize.Gorm()     // 初始化mysql连接
	global.POEM_MONGO = initialize.Mongo() // 初始化mongo连接
	core.Server()                          // 开启router服务
}

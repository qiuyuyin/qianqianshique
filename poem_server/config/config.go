package config

type Server struct {
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// 用户角色关系
	Casbin Casbin `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	// 日志
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
	// Mongo
	Mongo     Mongo     `mapstructure:"mongo" json:"mongo" yaml:"mongo"`
	System    System    `mapstructure:"system" json:"system" yaml:"system"`
	JWT       JWT       `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	AliyunOSS AliyunOSS `mapstructure:"aliyun-oss" json:"aliyunOSS" yaml:"aliyun-oss"`
}

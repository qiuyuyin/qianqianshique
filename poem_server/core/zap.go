package core

import (
    "fmt"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "os"
    "poem_server/global"
    "poem_server/utils"
    "time"
)

var level zapcore.Level

func Zap() (logger *zap.Logger) {
    // 判断是否有Director文件夹,如果没有则进行创建
    if ok, _ := utils.PathExists(global.POEM_CONFIG.Zap.Director); !ok {
        fmt.Printf("create %v directory\n", global.POEM_CONFIG.Zap.Director)
        _ = os.Mkdir(global.POEM_CONFIG.Zap.Director, os.ModePerm)
    }

    // 初始化配置文件的Level
    switch global.POEM_CONFIG.Zap.Level {
    case "debug":
        level = zap.DebugLevel
    case "info":
        level = zap.InfoLevel
    case "warn":
        level = zap.WarnLevel
    case "error":
        level = zap.ErrorLevel
    case "dpanic":
        level = zap.DPanicLevel
    case "panic":
        level = zap.PanicLevel
    case "fatal":
        level = zap.FatalLevel
    default:
        level = zap.InfoLevel
    }

    if level == zap.DebugLevel || level == zap.ErrorLevel {
        logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))
    } else {
        logger = zap.New(getEncoderCore())
    }
    if global.POEM_CONFIG.Zap.ShowLine {
        logger = logger.WithOptions(zap.AddCaller())
    }
    return logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
    config = zapcore.EncoderConfig{
        MessageKey:     "message",
        LevelKey:       "level",
        TimeKey:        "time",
        NameKey:        "logger",
        CallerKey:      "caller",
        StacktraceKey:  global.POEM_CONFIG.Zap.StacktraceKey,
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.LowercaseLevelEncoder,
        EncodeTime:     CustomTimeEncoder,
        EncodeDuration: zapcore.SecondsDurationEncoder,
        EncodeCaller:   zapcore.FullCallerEncoder,
    }
    switch {
    case global.POEM_CONFIG.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
        config.EncodeLevel = zapcore.LowercaseLevelEncoder
    case global.POEM_CONFIG.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
        config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
    case global.POEM_CONFIG.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
        config.EncodeLevel = zapcore.CapitalLevelEncoder
    case global.POEM_CONFIG.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
        config.EncodeLevel = zapcore.CapitalColorLevelEncoder
    default:
        config.EncodeLevel = zapcore.LowercaseLevelEncoder
    }
    return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
    if global.POEM_CONFIG.Zap.Format == "json" {
        return zapcore.NewJSONEncoder(getEncoderConfig())
    }
    return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore() (core zapcore.Core) {
    writer, err := utils.GetWriteSyncer() // 使用file-rotatelogs进行日志分割，一个小插件进行日志分割。。
    if err != nil {
        fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
        return
    }
    return zapcore.NewCore(getEncoder(), writer, level)
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
    enc.AppendString(t.Format(global.POEM_CONFIG.Zap.Prefix + "2006/01/02 - 15:04:05"))
}



package bootstrap

import (
	"7youo-wms/global"
	"7youo-wms/utils"
	"fmt"
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

var level zapcore.Level

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.G_CONFIG.Zap.Director); !ok {
		_ = os.Mkdir(global.G_CONFIG.Zap.Director, os.ModePerm)
	}

	level = zap.InfoLevel

	logger = zap.New(GetEncoderCore(), zap.AddStacktrace(level))
	if global.G_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return logger
}

func GetEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(GetEncoderConfig())
}

func GetEncoderConfig () (config zapcore.EncoderConfig)  {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.G_CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		//ConsoleSeparator: "  ",//分割符
	}
	return config
}

// 获取Encoder的zapcore.Core
func GetEncoderCore() (core zapcore.Core) {
	writer, err := GetWriteSyncer()
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(GetEncoder(), writer, level)
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.G_CONFIG.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}

//日志按日切割
func GetWriteSyncer() (zapcore.WriteSyncer, error)  {
	fileWriter, err := zaprotatelogs.New(
		path.Join(global.G_CONFIG.Zap.Director, "%Y-%m-%d.log"),
		zaprotatelogs.WithLinkName(global.G_CONFIG.Zap.LinkName),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
}
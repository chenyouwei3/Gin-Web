package runLog

import (
	conf "gin-web/init/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var ZapLog *zap.Logger

func InitRunLog(conf conf.Config) error {
	// 配置日志格式
	config := zapcore.EncoderConfig{
		TimeKey:      "time",
		LevelKey:     "level",
		MessageKey:   "msg",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	// 日志文件路径（每天一个）
	today := time.Now().Format("2006-01-02")
	logFilePath := "../logs/" + today + ".log"

	// 确保 logs 目录存在
	if err := os.MkdirAll("../logs", os.ModePerm); err != nil {
		return err
	}

	// 文件存在就追加，不存在就创建
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	writeSyncer := zapcore.AddSync(file)

	// 创建 core
	var core zapcore.Core
	if conf.APP.Mode == "debug" {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(config),
			writeSyncer,
			zapcore.DebugLevel, // 开发环境建议是 Debug
		)
	} else {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(config),
			writeSyncer,
			zapcore.InfoLevel,
		)
	}
	ZapLog = zap.New(core)
	return nil
}

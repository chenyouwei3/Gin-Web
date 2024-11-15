package runLog

import (
	conf "gin-web/initialize/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var ZapLog *zap.Logger

func InitRunLog() error {
	// 配置日志级别和编码
	config := zapcore.EncoderConfig{
		TimeKey:      "time",                      // 时间字段
		LevelKey:     "level",                     // 日志级别字段
		MessageKey:   "msg",                       // 日志消息字段
		EncodeTime:   zapcore.ISO8601TimeEncoder,  // 时间格式为 ISO8601
		EncodeLevel:  zapcore.CapitalLevelEncoder, // 级别为大写
		EncodeCaller: zapcore.ShortCallerEncoder,  // 简短调用者信息
	}
	// 创建文件写入器
	workDir, _ := os.Getwd()
	file, err := os.Create(workDir + "/logs/app.log") // 指定日志文件
	if err != nil {
		return err
	}
	writeSyncer := zapcore.AddSync(file)
	// 创建 logger
	var core zapcore.Core
	if conf.Conf.APP.Mode == "debug" {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(config), // JSON 格式日志
			writeSyncer,                    // 输出到文件
			zapcore.WarnLevel,              // 最低日志级别
		)
	} else {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(config), // JSON 格式日志
			writeSyncer,                    // 输出到文件
			zapcore.InfoLevel,              // 最低日志级别
		)
	}
	ZapLog = zap.New(core)
	defer ZapLog.Sync()
	return nil
}

package main

import (
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
)

const (
	DefaultTimeLayout = time.RFC3339
)

func main() {
	//timeLayout := DefaultTimeLayout
	logger := zap.NewExample()
	defer logger.Sync()

	url := "http://example.org/api"
	logger.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	logger2 := logger.With(zap.Namespace("metrics"), zap.Int("counter", 1))
	logger2.Info("tracked some metrics")

	logger3 := logger.WithOptions(zap.Fields(zap.Field{Key: "test", Type: zapcore.StringType, String: "呵呵"}))
	logger3.Info("test")
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
	cfg := zap.NewProductionEncoderConfig()
	//错误级别key
	cfg.LevelKey = "l"
	//日志记录时间key
	cfg.TimeKey = "sq"
	//时间编码格式
	//cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	cfg.EncodeDuration = zapcore.MillisDurationEncoder
	// 调用栈信息
	cfg.EncodeCaller = zapcore.ShortCallerEncoder
	//cfg.TimeKey = "time"
	//cfg
	//cfg := zapcore.EncoderConfig{
	//	MessageKey: "msg",
	//	LevelKey:   "level",
	//	TimeKey:    "time",
	//	EncodeTime: zapcore.ISO8601TimeEncoder,
	//}
	f, err := os.OpenFile("/Users/ganjian1/code/go-xstep/zap/test.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	enc := zapcore.NewJSONEncoder(cfg)
	core := zapcore.NewCore(
		enc,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(f), zapcore.AddSync(os.Stdout)),
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zap.DebugLevel
		}),
	)
	log := zap.New(core, zap.AddCaller(), zap.ErrorOutput(os.Stdout))
	log.Sugar().Infof("hello Failed to fetch URL: %s", url)

}

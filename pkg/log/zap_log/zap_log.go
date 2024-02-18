package zaplog

import (
	"encoding/json"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(rawJSON []byte) *zap.Logger {
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zap.Must(cfg.Build(zap.AddCaller(), zap.AddCallerSkip(1)))
}

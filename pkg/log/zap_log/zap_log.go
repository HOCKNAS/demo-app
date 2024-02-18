package zaplog

import (
	"encoding/json"

	"go.uber.org/zap"
)

func NewLogger(rawJSON []byte) *zap.Logger {
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	return zap.Must(cfg.Build())
}

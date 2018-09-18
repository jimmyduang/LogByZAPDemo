package main

import (
	"encoding/json"
	"fmt"
	"time"

	"go.uber.org/zap"
)

func main() {
	date := time.Now().Format("2006-01-02")
	now := time.Now().Format("2006-01-02 15:04:05")
	rawJSON := []byte(fmt.Sprintf(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout", "/tmp/logs/%s"],
		"errorOutputPaths": ["stderr"],
		"initialFields": {"foo": "bar","time": "%s"},
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase"
		}
	  }`, date, now))

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Debug("logger construction succeeded")
	logger.Info("logger construction succeeded")
	logger.Warn("logger construction succeeded")
	logger.Error("logger construction succeeded")
}

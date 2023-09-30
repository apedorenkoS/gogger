package config

import (
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog"
	"strings"
)

const (
	logLevelConf = "log.level"
	logEnvConf   = "log.env"
)

var k = koanf.New(".")

func LoadConfig() error {
	return k.Load(env.Provider("", ".", func(s string) string {
		return strings.Replace(strings.ToLower(s), "_", ".", -1)
	}), nil)
	// validate required configs
}

func LogLevel() zerolog.Level {
	logLevel := k.String(logLevelConf)
	switch logLevel {
	case "trace":
		return zerolog.TraceLevel
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "panic":
		return zerolog.PanicLevel
	case "fatal":
		return zerolog.FatalLevel
	default:
		return zerolog.DebugLevel

	}
}

type LogEnvironment string

const (
	Dev  LogEnvironment = "dev"
	Prod LogEnvironment = "prod"
)

func LogEnv() LogEnvironment {
	logEnv := k.String(logEnvConf)
	switch logEnv {
	case string(Dev):
		return Dev
	case string(Prod):
		return Prod
	default:
		return Prod
	}
}

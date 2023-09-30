package main

import (
	"github.com/apedorenkoS/gogger/cmd/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func InitGlobalLogger() {
	log.Logger = log.Logger.
		Hook(AddCorrelationIdHook{}). // log x-correlation-id in each log message
		Level(config.LogLevel())

	if config.LogEnv() == config.Dev {
		log.Logger = log.Logger.Output(zerolog.ConsoleWriter{Out: os.Stdout}) // not-structured pretty colored format
	}
}

type AddCorrelationIdHook struct{}

func (cid AddCorrelationIdHook) Run(e *zerolog.Event, _ zerolog.Level, _ string) {
	ctx := e.GetCtx()
	if ctx != nil {
		correlationId := ctx.Value(CtxCorrelationIdName)
		if corrIdStr, ok := correlationId.(string); ok {
			e.Str("x-correlation-id", corrIdStr)
		}
	}
}

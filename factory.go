package fxlogger

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type LoggerFactory interface {
	Create(options ...LoggerOption) (*Logger, error)
}

type DefaultLoggerFactory struct{}

func NewDefaultLoggerFactory() LoggerFactory {
	return &DefaultLoggerFactory{}
}

func (f *DefaultLoggerFactory) Create(options ...LoggerOption) (*Logger, error) {

	zerolog.TimeFieldFormat = time.RFC3339

	appliedOpts := defaultLoggerOptions
	for _, applyOpt := range options {
		applyOpt(&appliedOpts)
	}

	log.Ctx(appliedOpts.Context)
	logger := log.
		Output(appliedOpts.OutputWriter).
		With().
		Str(Service, appliedOpts.Name).
		Logger().
		Level(appliedOpts.Level)

		//WithContext(appliedOpts.Context)

	return &Logger{&logger, false}, nil
}

package fxlogger

import (
	"context"
	"io"
	"os"

	"github.com/rs/zerolog"
)

type options struct {
	Name         string
	Level        zerolog.Level
	OutputWriter io.Writer
	Context  context.Context
}

var defaultLoggerOptions = options{
	Name:         "default",
	Level:        zerolog.InfoLevel,
	OutputWriter: os.Stdout,
	Context: context.Background(),
}

type LoggerOption func(o *options)

func WithName(n string) LoggerOption {
	return func(o *options) {
		o.Name = n
	}
}


func WithLevel(l zerolog.Level) LoggerOption {
	return func(o *options) {
		o.Level = l
	}
}

func WithOutputWriter(w io.Writer) LoggerOption {
	return func(o *options) {
		o.OutputWriter = w
	}
}

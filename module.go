package fxlogger

import (
	"io"
	"os"
	 //config "gotemplate/config"
	  "github.com/templatedop/fxconfigs/config"
	 //"github.com/templatedop/config"
	"github.com/templatedop/fxlog/fxloggertest"

	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

var FxLoggerModule = fx.Module("logger",
	fx.Provide(
		NewDefaultLoggerFactory,
		NewFxLogger,
		fx.Annotate(
			fxloggertest.GetTestLogBufferInstance,
			fx.ResultTags(`name:"test-log-buffer"`),
		),
	),
)

type FxLoggerParam struct {
	fx.In
	Factory LoggerFactory
	Config  *config.Config
	//Config  *fxconfig.EConfig
}

func NewFxLogger(p FxLoggerParam) (*Logger, error) {
	// level
	level := FetchLogLevel(p.Config.GetString("modules.log.level"))
	if p.Config.GetString("modules.log.level")=="debug" {
		level = zerolog.DebugLevel
	}	

	// output writer
	var outputWriter io.Writer
	if p.Config.GetString("app.env") == "test" {
		outputWriter = fxloggertest.GetTestLogBufferInstance()
	} else {
		outputWriter = os.Stdout
	}

	// logger
	return p.Factory.Create(
		WithName(p.Config.GetString("app.name")),		
		WithLevel(level),
		WithOutputWriter(outputWriter),
	)
}
